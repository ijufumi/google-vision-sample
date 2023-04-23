package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucket;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketCdnPolicy;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketCdnPolicyCacheKeyPolicy;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketConfig;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class CloudCDNStack {

  public static void create(final Construct scope, final Config config) {
    var backendBucketCdnPolicyCacheKeyPolicy = ComputeBackendBucketCdnPolicyCacheKeyPolicy.builder()
        .build();
    var backendBucketCdnPolicy = ComputeBackendBucketCdnPolicy
        .builder()
        .cacheKeyPolicy(backendBucketCdnPolicyCacheKeyPolicy)
        .cacheMode("CACHE_ALL_STATIC")
        .build();
    var backendBucketConfig = ComputeBackendBucketConfig
        .builder()
        .bucketName(config.BucketName())
        .enableCdn(true)
        .cdnPolicy(backendBucketCdnPolicy)
        .compressionMode("AUTOMATIC")
        .build();
    new ComputeBackendBucket(scope, "backend-bucket", backendBucketConfig);
  }
}
