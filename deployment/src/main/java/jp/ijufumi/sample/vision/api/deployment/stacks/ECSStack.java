package jp.ijufumi.sample.vision.api.deployment.stacks;

import java.util.List;
import java.util.Map;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.services.ec2.InstanceClass;
import software.amazon.awscdk.services.ec2.InstanceSize;
import software.amazon.awscdk.services.ec2.InstanceType;
import software.amazon.awscdk.services.ecs.AddCapacityOptions;
import software.amazon.awscdk.services.ecs.Cluster.Builder;
import software.amazon.awscdk.services.ecs.Compatibility;
import software.amazon.awscdk.services.ecs.ContainerDefinitionProps;
import software.amazon.awscdk.services.ecs.ContainerImage;
import software.amazon.awscdk.services.ecs.Ec2Service;
import software.amazon.awscdk.services.ecs.PortMapping;
import software.amazon.awscdk.services.ecs.Secret;
import software.amazon.awscdk.services.ecs.TaskDefinition;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationListener;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationLoadBalancer;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationTargetGroup;
import software.amazon.awscdk.services.iam.ManagedPolicy;
import software.amazon.awscdk.services.iam.PolicyStatement;
import software.amazon.awscdk.services.iam.Role;
import software.amazon.awscdk.services.iam.ServicePrincipal;
import software.constructs.Construct;

public class ECSStack {

  public static ApplicationLoadBalancer build(final Construct scope, final Config config,
      final ContainerImage appImage,
      final software.amazon.awscdk.services.secretsmanager.Secret googleCredential) {
    var statement = PolicyStatement
        .Builder
        .create()
        .build();
    statement.addActions("s3:*");
    statement.addAllResources();

    var ecsTaskRolePolicy = ManagedPolicy
        .Builder
        .create(scope, "ecs-role-policy")
        .statements(List.of(statement))
        .build();

    var ecsTaskServicePrincipal = ServicePrincipal
        .Builder
        .create("ecs-tasks.amazonaws.com")
        .build();

    var ecsRole = Role
        .Builder
        .create(scope, "ecs-role")
        .roleName("ecs-role")
        .assumedBy(ecsTaskServicePrincipal)
        .managedPolicies(List.of(ecsTaskRolePolicy))
        .build();

    var capacityOptions = AddCapacityOptions
        .builder()
        .instanceType(InstanceType.of(InstanceClass.ARM1, InstanceSize.MICRO))
        .allowAllOutbound(true)
        .build();

    var ecsCluster = Builder
        .create(scope, "ecs-cluster")
        .clusterName("ecs-cluster")
        .capacity(capacityOptions)
        .build();

    var appTaskDefinition = TaskDefinition
        .Builder
        .create(scope, "app-task-definition")
        .compatibility(Compatibility.EC2)
        .taskRole(ecsRole)
        .build();

    var googleCredentialSecret = Secret.fromSecretsManager(googleCredential);

    var appPortMapping = PortMapping
        .builder()
        .containerPort(8080)
        .hostPort(80)
        .build();
    var appEnvironment = Map.of(
        "DB_HOST", config.dbHost(),
        "DB_NAME", config.dbName(),
        "DB_USER", config.dbUser(),
        "DB_PASSWORD", config.dbPassword(),
        "DB_PORT", Integer.toString(config.dbPort())
    );
    var appContainer = ContainerDefinitionProps
        .builder()
        .containerName("db")
        .taskDefinition(appTaskDefinition)
        .image(appImage)
        .portMappings(List.of(appPortMapping))
        .secrets(Map.of("GOOGLE_CREDENTIAL", googleCredentialSecret))
        .environment(appEnvironment)
        .build();
    appTaskDefinition.addContainer("app-container", appContainer);

    var app = Ec2Service
        .Builder
        .create(scope, "app-service")
        .assignPublicIp(false)
        .cluster(ecsCluster)
        .serviceName("app")
        .build();

    var alb = ApplicationLoadBalancer
        .Builder
        .create(scope, "ecs-alb")
        .loadBalancerName(config.apiDomainName())
        .build();

    var albTargetGroup = ApplicationTargetGroup
        .Builder
        .create(scope, "alb-target-group")
        .targets(List.of(app))
        .build();
    ApplicationListener
        .Builder
        .create(scope, "ecs-alb-listener")
        .loadBalancer(alb)
        .defaultTargetGroups(List.of(albTargetGroup))
        .build();

    var dbTaskDefinition = TaskDefinition
        .Builder
        .create(scope, "db-task-definition")
        .compatibility(Compatibility.EC2)
        .taskRole(ecsRole)
        .build();

    var dbImage = ContainerImage
        .fromRegistry("postgres:latest");

    var dbContainer = ContainerDefinitionProps
        .builder()
        .containerName("db")
        .taskDefinition(dbTaskDefinition)
        .image(dbImage)
        .build();
    dbTaskDefinition.addContainer("db-container", dbContainer);

    Ec2Service
        .Builder
        .create(scope, "db-service")
        .assignPublicIp(false)
        .cluster(ecsCluster)
        .serviceName("db")
        .build();

    return alb;
  }
}
