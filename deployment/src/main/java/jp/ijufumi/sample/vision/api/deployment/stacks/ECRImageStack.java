package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.core.RemovalPolicy;
import software.amazon.awscdk.services.ecr.Repository;
import software.amazon.awscdk.services.ecr.TagMutability;
import software.amazon.awscdk.services.ecr.assets.DockerImageAsset;
import software.amazon.awscdk.services.ecr.assets.Platform;
import software.amazon.awscdk.services.ecs.ContainerImage;
import software.constructs.Construct;

public class ECRImageStack {

  public static ContainerImage build(final Construct scope, final Config config) {
    var ecrRepository = Repository
        .Builder
        .create(scope, "ecr-repository")
        .repositoryName(config.repositoryName())
        .imageScanOnPush(true)
        .imageTagMutability(TagMutability.MUTABLE)
        .removalPolicy(RemovalPolicy.DESTROY)
        .build();
    var dockerImageAsset = DockerImageAsset
        .Builder
        .create(scope, "docker-image-asset")
        .directory(config.backendCode())
        .platform(Platform.LINUX_ARM64)
        .build();
    dockerImageAsset.setRepository(ecrRepository);
    return ContainerImage.fromDockerImageAsset(dockerImageAsset);
  }
}
