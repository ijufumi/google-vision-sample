package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucket;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketCdnPolicy;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketConfig;
import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucket;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class CloudCDNStack {

  public static void create(final Construct scope, final Config config,
      final StorageBucket bucket) {
    var backendBucketCdnPolicy = ComputeBackendBucketCdnPolicy
        .builder()
        .cacheMode("CACHE_ALL_STATIC")
        .maxTtl(config.BackendBucketCdnPolicyTTL())
        .clientTtl(config.BackendBucketCdnPolicyTTL())
        .defaultTtl(config.BackendBucketCdnPolicyTTL())
        .build();
    var backendBucketConfig = ComputeBackendBucketConfig
        .builder()
        .bucketName(bucket.getName())
        .enableCdn(true)
        .cdnPolicy(backendBucketCdnPolicy)
        .compressionMode("AUTOMATIC")
        .name(config.BackendBucketName())
        .build();
    new ComputeBackendBucket(scope, "backend-bucket", backendBucketConfig);
  }
}
