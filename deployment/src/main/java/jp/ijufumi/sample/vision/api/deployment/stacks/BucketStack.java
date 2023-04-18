package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;
import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucket;
import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucketConfig;
import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucketCors;

import java.util.List;

public class BucketStack {
    public static void create(Construct scope, Config config) {
        var bucketCors = StorageBucketCors
                .builder()
                .method(config.BucketCorsMethods())
                .origin(config.BucketCorsOrigins())
                .maxAgeSeconds(config.BucketCorsMaxAge())
                .build();
        var bucketConfig = StorageBucketConfig
                .builder()
                .project(config.ProjectId())
                .location(config.Location())
                .name(config.BucketName())
                .cors(List.of(bucketCors))
                .build();
        new StorageBucket(scope, "storage-bucket", bucketConfig);
    }
}
