package jp.ijufumi.sample.vision.api.deployment.stacks;

import java.util.List;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.services.certificatemanager.Certificate;
import software.amazon.awscdk.services.cloudfront.AllowedMethods;
import software.amazon.awscdk.services.cloudfront.BehaviorOptions;
import software.amazon.awscdk.services.cloudfront.CacheHeaderBehavior;
import software.amazon.awscdk.services.cloudfront.CachePolicy;
import software.amazon.awscdk.services.cloudfront.CacheQueryStringBehavior;
import software.amazon.awscdk.services.cloudfront.CachedMethods;
import software.amazon.awscdk.services.cloudfront.Distribution;
import software.amazon.awscdk.services.cloudfront.OriginProtocolPolicy;
import software.amazon.awscdk.services.cloudfront.PriceClass;
import software.amazon.awscdk.services.cloudfront.ViewerProtocolPolicy;
import software.amazon.awscdk.services.cloudfront.origins.HttpOrigin;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationLoadBalancer;
import software.amazon.awscdk.services.s3.IBucket;
import software.constructs.Construct;

/**
 * @reference
 * https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/quickref-cloudfront.html
 */
public class APICloudfrontStack {

  public static Distribution build(final Construct scope, final Config config, final IBucket bucket,
      final
      ApplicationLoadBalancer alb) {
    var httpOrigin = HttpOrigin
        .Builder
        .create(alb.getLoadBalancerDnsName())
        .httpPort(80)
        .protocolPolicy(OriginProtocolPolicy.HTTP_ONLY)
        .build();

    var cachePolicy = CachePolicy
        .Builder
        .create(scope, "cache-policy-for-api")
        .queryStringBehavior(CacheQueryStringBehavior.all())
        .headerBehavior(CacheHeaderBehavior.allowList("Access-Control-Request-Headers",
            "Access-Control-Request-Method", "Origin"))
        .build();

    var behaviorOption = BehaviorOptions
        .builder()
        .origin(httpOrigin)
        .cachedMethods(CachedMethods.CACHE_GET_HEAD_OPTIONS)
        .allowedMethods(AllowedMethods.ALLOW_ALL)
        .viewerProtocolPolicy(ViewerProtocolPolicy.ALLOW_ALL)
        .cachePolicy(cachePolicy)
        .compress(true)
        .build();

    var certificate = Certificate.fromCertificateArn(scope, "api-cloudfront-certificate",
        config.certificationArn());
    return Distribution
        .Builder
        .create(scope, "cloudfront-for-api")
        .defaultBehavior(behaviorOption)
        .enabled(true)
        .enableLogging(true)
        .logBucket(bucket)
        .logFilePrefix("logs/api")
        .logIncludesCookies(true)
        .priceClass(PriceClass.PRICE_CLASS_200)
        .domainNames(List.of(config.apiDomainFullName()))
        .certificate(certificate)
        .defaultRootObject("/")
        .build();
  }
}
