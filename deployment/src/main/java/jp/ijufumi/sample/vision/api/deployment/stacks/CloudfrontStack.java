package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.services.cloudfront.BehaviorOptions;
import software.amazon.awscdk.services.cloudfront.CachePolicy;
import software.amazon.awscdk.services.cloudfront.CacheQueryStringBehavior;
import software.amazon.awscdk.services.cloudfront.CachedMethods;
import software.amazon.awscdk.services.cloudfront.Distribution;
import software.amazon.awscdk.services.cloudfront.PriceClass;
import software.amazon.awscdk.services.cloudfront.ViewerProtocolPolicy;
import software.amazon.awscdk.services.cloudfront.origins.S3Origin;
import software.amazon.awscdk.services.s3.IBucket;
import software.constructs.Construct;

/**
 * @reference
 * https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/quickref-cloudfront.html
 */
public class CloudfrontStack {

  public static void build(final Construct scope, final Config config, final IBucket bucket) {
    var s3Origin = S3Origin
        .Builder
        .create(bucket)
        .build();

    var cachePolicy = CachePolicy
        .Builder
        .create(scope, "cache-policy")
        .queryStringBehavior(CacheQueryStringBehavior.all())
        .build();

    var behaviorOption = BehaviorOptions
        .builder()
        .origin(s3Origin)
        .cachedMethods(CachedMethods.CACHE_GET_HEAD_OPTIONS)
        .viewerProtocolPolicy(ViewerProtocolPolicy.ALLOW_ALL)
        .cachePolicy(cachePolicy)
        .build();

    Distribution
        .Builder
        .create(scope, "id-cloudfront")
        .defaultBehavior(behaviorOption)
        .enabled(true)
        .defaultRootObject("index.html")
        .logBucket(bucket)
        .logFilePrefix("logs")
        .logIncludesCookies(true)
        .priceClass(PriceClass.PRICE_CLASS_200)
        .build();
  }
}
