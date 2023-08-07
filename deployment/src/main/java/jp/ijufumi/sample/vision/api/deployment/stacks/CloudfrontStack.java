package jp.ijufumi.sample.vision.api.deployment.stacks;

import java.util.List;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.core.Duration;
import software.amazon.awscdk.services.cloudfront.AllowedMethods;
import software.amazon.awscdk.services.cloudfront.BehaviorOptions;
import software.amazon.awscdk.services.cloudfront.CachePolicy;
import software.amazon.awscdk.services.cloudfront.CacheQueryStringBehavior;
import software.amazon.awscdk.services.cloudfront.CachedMethods;
import software.amazon.awscdk.services.cloudfront.Distribution;
import software.amazon.awscdk.services.cloudfront.ErrorResponse;
import software.amazon.awscdk.services.cloudfront.GeoRestriction;
import software.amazon.awscdk.services.cloudfront.OriginAccessIdentity;
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

  public static void build(final Construct scope, final IBucket bucket) {

    var originAccessIdentity = OriginAccessIdentity
        .Builder
        .create(scope, "origin-access-identity")
        .build();

    var s3Origin = S3Origin
        .Builder
        .create(bucket)
        .originAccessIdentity(originAccessIdentity)
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
        .allowedMethods(AllowedMethods.ALLOW_GET_HEAD_OPTIONS)
        .viewerProtocolPolicy(ViewerProtocolPolicy.ALLOW_ALL)
        .cachePolicy(cachePolicy)
        .build();

    var errorResponse401 = ErrorResponse
        .builder()
        .httpStatus(401)
        .responseHttpStatus(200)
        .responsePagePath("index.html")
        .ttl(Duration.millis(0))
        .build();

    var errorResponse403 = ErrorResponse
        .builder()
        .httpStatus(403)
        .responseHttpStatus(200)
        .responsePagePath("index.html")
        .ttl(Duration.millis(0))
        .build();

    var errorResponse404 = ErrorResponse
        .builder()
        .httpStatus(404)
        .responseHttpStatus(200)
        .responsePagePath("index.html")
        .ttl(Duration.millis(0))
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
        .geoRestriction(GeoRestriction.allowlist("AQ", "CV"))
        .errorResponses(List.of(errorResponse401, errorResponse403, errorResponse404))
        .build();
  }
}
