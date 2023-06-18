package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucket;
import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucketConfig;
import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucketCors;
import com.hashicorp.cdktf.providers.google.storage_bucket_iam_member.StorageBucketIamMember;
import com.hashicorp.cdktf.providers.google.storage_bucket_iam_member.StorageBucketIamMemberConfig;
import java.util.List;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class BucketStack {

  public static StorageBucket create(final Construct scope, final Config config) {
    var bucketCors = StorageBucketCors
        .builder()
        .method(config.BucketCorsMethods())
        .origin(config.BucketCorsOrigins())
        .maxAgeSeconds(config.BucketCorsMaxAge())
        .build();
    var bucketConfig = StorageBucketConfig
        .builder()
        .project(config.ProjectId())
        .location(config.Region())
        .name(config.BucketName())
        .cors(List.of(bucketCors))
        .forceDestroy(true)
        .build();
    var bucket = new StorageBucket(scope, "storage-bucket", bucketConfig);

    var bucketIamMemberConfig = StorageBucketIamMemberConfig
        .builder()
        .bucket(bucket.getName())
        .member("allUsers")
        .role("roles/storage.legacyObjectReader")
        .build();
    new StorageBucketIamMember(scope, "storage-bucket-iam", bucketIamMemberConfig);

    return bucket;
  }
}
