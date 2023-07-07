package jp.ijufumi.sample.vision.api.deployment.stacks;

import java.util.List;
import software.amazon.awscdk.services.ecs.Cluster;
import software.amazon.awscdk.services.ecs.Compatibility;
import software.amazon.awscdk.services.ecs.ContainerDefinitionProps;
import software.amazon.awscdk.services.ecs.ContainerImage;
import software.amazon.awscdk.services.ecs.Ec2Service;
import software.amazon.awscdk.services.ecs.TaskDefinition;
import software.amazon.awscdk.services.iam.ManagedPolicy;
import software.amazon.awscdk.services.iam.PolicyStatement;
import software.amazon.awscdk.services.iam.Role;
import software.amazon.awscdk.services.iam.ServicePrincipal;
import software.constructs.Construct;

public class ECSStack {

  public static void build(final Construct scope) {
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

    var ecsCluster = Cluster
        .Builder
        .create(scope, "ecs-cluster")
        .clusterName("ecs-cluster")
        .build();

    var appTaskDefinition = TaskDefinition
        .Builder
        .create(scope, "app-task-definition")
        .compatibility(Compatibility.EC2)
        .taskRole(ecsRole)
        .build();

    var appImage = ContainerImage
        .fromRegistry("postgres:latest");

    var appContainer = ContainerDefinitionProps
        .builder()
        .containerName("db")
        .taskDefinition(appTaskDefinition)
        .image(appImage)
        .build();
    appTaskDefinition.addContainer("app-container", appContainer);

    Ec2Service
        .Builder
        .create(scope, "app-service")
        .assignPublicIp(true)
        .cluster(ecsCluster)
        .serviceName("app")
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
  }
}
