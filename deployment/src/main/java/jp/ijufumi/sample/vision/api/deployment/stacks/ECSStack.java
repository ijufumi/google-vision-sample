package jp.ijufumi.sample.vision.api.deployment.stacks;

import software.amazon.awscdk.services.ecs.Cluster;
import software.amazon.awscdk.services.ecs.Compatibility;
import software.amazon.awscdk.services.ecs.ContainerDefinitionProps;
import software.amazon.awscdk.services.ecs.ContainerImage;
import software.amazon.awscdk.services.ecs.Ec2Service;
import software.amazon.awscdk.services.ecs.TaskDefinition;
import software.constructs.Construct;

public class ECSStack {

  public static void build(final Construct scope) {

    var ecsCluster = Cluster
        .Builder
        .create(scope, "ecs-cluster")
        .clusterName("ecs-cluster")
        .build();

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
