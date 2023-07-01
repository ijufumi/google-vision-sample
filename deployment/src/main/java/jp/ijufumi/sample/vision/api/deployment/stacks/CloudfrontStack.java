package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.services.cloudfront.BehaviorOptions;
import software.amazon.awscdk.services.cloudfront.Distribution;
import software.amazon.awscdk.services.cloudfront.origins.S3Origin;
import software.amazon.awscdk.services.s3.IBucket;
import software.constructs.Construct;

public class CloudfrontStack {

  public static void build(final Construct scope, final Config config, final IBucket bucket) {
    var s3Origin = S3Origin
        .Builder
        .create(bucket)
        .build();
    var behaviorOption = BehaviorOptions
        .builder()
        .origin(s3Origin)
        .build();
    Distribution
        .Builder
        .create(scope, "id-cloudfront")
        .defaultBehavior(behaviorOption)
        .build();
  }
}
