package jp.ijufumi.sample.vision.api.deployment.stacks;

import java.util.List;
import java.util.Map;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.Duration;
import software.amazon.awscdk.RemovalPolicy;
import software.amazon.awscdk.services.ec2.InstanceClass;
import software.amazon.awscdk.services.ec2.InstanceSize;
import software.amazon.awscdk.services.ec2.InstanceType;
import software.amazon.awscdk.services.ec2.Peer;
import software.amazon.awscdk.services.ec2.Port;
import software.amazon.awscdk.services.ec2.SecurityGroup;
import software.amazon.awscdk.services.ec2.SubnetSelection;
import software.amazon.awscdk.services.ec2.Vpc;
import software.amazon.awscdk.services.ecs.AddCapacityOptions;
import software.amazon.awscdk.services.ecs.AwsLogDriverProps;
import software.amazon.awscdk.services.ecs.CloudMapOptions;
import software.amazon.awscdk.services.ecs.Cluster.Builder;
import software.amazon.awscdk.services.ecs.Compatibility;
import software.amazon.awscdk.services.ecs.ContainerDefinitionProps;
import software.amazon.awscdk.services.ecs.ContainerImage;
import software.amazon.awscdk.services.ecs.Ec2Service;
import software.amazon.awscdk.services.ecs.LogDriver;
import software.amazon.awscdk.services.ecs.NetworkMode;
import software.amazon.awscdk.services.ecs.PortMapping;
import software.amazon.awscdk.services.ecs.Secret;
import software.amazon.awscdk.services.ecs.TaskDefinition;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationListener;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationLoadBalancer;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationTargetGroup;
import software.amazon.awscdk.services.elasticloadbalancingv2.HealthCheck;
import software.amazon.awscdk.services.iam.Effect;
import software.amazon.awscdk.services.iam.ManagedPolicy;
import software.amazon.awscdk.services.iam.PolicyStatement;
import software.amazon.awscdk.services.iam.Role;
import software.amazon.awscdk.services.iam.ServicePrincipal;
import software.amazon.awscdk.services.logs.LogGroup;
import software.amazon.awscdk.services.servicediscovery.DnsRecordType;
import software.amazon.awscdk.services.servicediscovery.PrivateDnsNamespace;
import software.constructs.Construct;

public class ECSStack {

  public static ApplicationLoadBalancer build(final Construct scope, final Config config,
      final Vpc vpc,
      final ContainerImage appImage,
      final software.amazon.awscdk.services.secretsmanager.Secret googleCredential) {

    var statement = PolicyStatement
        .Builder
        .create()
        .actions(List.of("s3:*", "logs:*", "ecr:*"))
        .resources(List.of("*"))
        .effect(Effect.ALLOW)
        .build();

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
        .instanceType(InstanceType.of(InstanceClass.T3A, InstanceSize.XLARGE))
        .allowAllOutbound(true)
        .build();

    var privateDnsNamespace = PrivateDnsNamespace
        .Builder.create(scope, "private-dns-namespace")
        .vpc(vpc)
        .name(config.route53Namespace())
        .build();

    var ecsCluster = Builder
        .create(scope, "ecs-cluster")
        .clusterName("ecs-cluster")
        .capacity(capacityOptions)
        .vpc(vpc)
        .build();

    var appTaskDefinition = TaskDefinition
        .Builder
        .create(scope, "app-task-definition")
        .family("app-task-definition")
        .compatibility(Compatibility.EC2)
        .taskRole(ecsRole)
        .executionRole(ecsRole)
        .networkMode(NetworkMode.AWS_VPC)
        .build();

    var googleCredentialSecret = Secret.fromSecretsManager(googleCredential);

    var appPortMapping = PortMapping
        .builder()
        .containerPort(8080)
        .hostPort(8080)
        .build();
    var appEnvironment = Map.of(
        "DB_HOST", String.format("%s.%s", config.dbHost(), config.route53Namespace()),
        "DB_NAME", config.dbName(),
        "DB_USER", config.dbUser(),
        "DB_PASSWORD", config.dbPassword(),
        "DB_PORT", Integer.toString(config.dbPort()),
        "GOOGLE_STORAGE_BUCKET", "ijufumi-sample",
        "GIN_MOD", "release",
        "MIGRATION_PATH", "/app/migration"
    );
    var appLogGroup = LogGroup.Builder
        .create(scope, "app-container-log-group")
        .logGroupName("app-container")
        .removalPolicy(RemovalPolicy.DESTROY)
        .build();
    var appLogProps = AwsLogDriverProps
        .builder()
        .logGroup(appLogGroup)
        .streamPrefix("app-container")
        .build();
    var appLogConfig = LogDriver.awsLogs(appLogProps);
    var appContainer = ContainerDefinitionProps
        .builder()
        .containerName("app")
        .image(appImage)
        .portMappings(List.of(appPortMapping))
        .secrets(Map.of("GOOGLE_CREDENTIAL", googleCredentialSecret))
        .environment(appEnvironment)
        .taskDefinition(appTaskDefinition)
        .cpu(100)
        .memoryLimitMiB(256)
        .privileged(true)
        .logging(appLogConfig)
        .build();
    appTaskDefinition.addContainer("app-container", appContainer);

    var appCloudMapOption = CloudMapOptions
        .builder()
        .dnsRecordType(DnsRecordType.A)
        .dnsTtl(Duration.seconds(300))
        .failureThreshold(1)
        .cloudMapNamespace(privateDnsNamespace)
        .name("app")
        .build();
    var app = Ec2Service
        .Builder
        .create(scope, "app-service")
        .assignPublicIp(false)
        .cluster(ecsCluster)
        .serviceName("app")
        .taskDefinition(appTaskDefinition)
        .cloudMapOptions(appCloudMapOption)
        .vpcSubnets(SubnetSelection.builder().subnets(vpc.getPrivateSubnets()).build())
        .build();

    var albSecurityGroup = SecurityGroup
        .Builder
        .create(scope, "alb-security-group")
        .vpc(vpc)
        .securityGroupName("alb-security-group")
        .build();
    albSecurityGroup.addIngressRule(Peer.prefixList("pl-58a04531"), Port.tcp(80));

    var alb = ApplicationLoadBalancer
        .Builder
        .create(scope, "ecs-alb")
        .vpc(vpc)
        .securityGroup(albSecurityGroup)
        .internetFacing(true)
        .vpcSubnets(SubnetSelection.builder().subnets(vpc.getPublicSubnets()).build())
        .build();

    var albTargetGroup = ApplicationTargetGroup
        .Builder
        .create(scope, "alb-target-group")
        .targets(List.of(app))
        .port(8080)
        .vpc(vpc)
        .healthCheck(HealthCheck.builder().path("/api/health").build())
        .build();

    ApplicationListener
        .Builder
        .create(scope, "ecs-alb-listener")
        .loadBalancer(alb)
        .defaultTargetGroups(List.of(albTargetGroup))
        .port(80)
        .build();

    var dbTaskDefinition = TaskDefinition
        .Builder
        .create(scope, "db-task-definition")
        .family("db-task-definition")
        .compatibility(Compatibility.EC2)
        .taskRole(ecsRole)
        .networkMode(NetworkMode.AWS_VPC)
        .build();

    var dbImage = ContainerImage
        .fromRegistry("postgres:latest");

    var dbLogGroup = LogGroup.Builder
        .create(scope, "db-container-log-group")
        .logGroupName("db-container")
        .removalPolicy(RemovalPolicy.DESTROY)
        .build();
    var dbLogProps = AwsLogDriverProps
        .builder()
        .logGroup(dbLogGroup)
        .streamPrefix("db-container")
        .build();
    var dbLogConfig = LogDriver.awsLogs(dbLogProps);
    var dbEnvironment = Map.of(
        "POSTGRES_DB", config.dbName(),
        "POSTGRES_USER", config.dbUser(),
        "POSTGRES_PASSWORD", config.dbPassword()
    );

    var dbPortMapping = PortMapping
        .builder()
        .containerPort(5432)
        .hostPort(5432)
        .build();
    var dbContainer = ContainerDefinitionProps
        .builder()
        .containerName(config.dbHost())
        .image(dbImage)
        .portMappings(List.of(dbPortMapping))
        .taskDefinition(dbTaskDefinition)
        .environment(dbEnvironment)
        .logging(dbLogConfig)
        .cpu(100)
        .memoryLimitMiB(256)
        .privileged(true)
        .build();
    dbTaskDefinition.addContainer("db-container", dbContainer);

    var dbCloudMapOption = CloudMapOptions
        .builder()
        .dnsRecordType(DnsRecordType.A)
        .dnsTtl(Duration.seconds(300))
        .failureThreshold(1)
        .name("db")
        .cloudMapNamespace(privateDnsNamespace)
        .build();
    var dbSecurityGroup = SecurityGroup.Builder.create(scope, "db-security-group").vpc(vpc)
        .allowAllOutbound(true).build();
    dbSecurityGroup.addIngressRule(Peer.ipv4(vpc.getVpcCidrBlock()), Port.tcp(5432));
    Ec2Service
        .Builder
        .create(scope, "db-service")
        .assignPublicIp(false)
        .cluster(ecsCluster)
        .serviceName("db")
        .taskDefinition(dbTaskDefinition)
        .cloudMapOptions(dbCloudMapOption)
        .vpcSubnets(SubnetSelection.builder().subnets(vpc.getPrivateSubnets()).build())
        .securityGroups(List.of(dbSecurityGroup))
        .build();

    return alb;
  }
}
