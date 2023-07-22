package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.core.RemovalPolicy;
import software.amazon.awscdk.services.s3.Bucket;
import software.amazon.awscdk.services.s3.IBucket;
import software.constructs.Construct;

public class S3Stack {

  public static IBucket build(final Construct scope, final Config config) {
    return Bucket.Builder
        .create(scope, "bucket")
        .bucketName(config.bucket())
        .versioned(true)
        .autoDeleteObjects(true)
        .removalPolicy(RemovalPolicy.DESTROY)
        .build();
  }
}
