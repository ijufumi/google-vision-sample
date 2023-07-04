package jp.ijufumi.sample.vision.api.deployment.stacks;

import software.amazon.awscdk.services.ecs.Cluster;
import software.amazon.awscdk.services.ecs.Ec2Service;
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
        .build();

    Ec2Service
        .Builder
        .create(scope, "db-service")
        .assignPublicIp(false)
        .cluster(ecsCluster)
        .build();
  }
}
