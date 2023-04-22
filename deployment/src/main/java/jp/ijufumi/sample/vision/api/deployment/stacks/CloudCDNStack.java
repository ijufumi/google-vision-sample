package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucket;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketConfig;

public class CloudCDNStack {
    public static void create(final Construct scope, final Config config) {
        var backendBucketConfig = ComputeBackendBucketConfig
            .builder()
            .bucketName(config.BucketName())
            .build();
        var backendBucket = new ComputeBackendBucket(scope, "backend-bucket", backendBucketConfig);
    }
}
