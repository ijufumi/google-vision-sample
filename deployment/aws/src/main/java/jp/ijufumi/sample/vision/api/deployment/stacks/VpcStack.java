package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.services.ec2.Vpc;
import software.constructs.Construct;

public class VpcStack {

  public static Vpc build(final Construct scope, final Config config) {
    return Vpc
        .Builder
        .create(scope, "vpc")
        .vpcName(config.vpcName())
        .maxAzs(2)
        .build();
  }
}
