package jp.ijufumi.sample.vision.api.deployment;

import com.hashicorp.cdktf.TerraformStack;
import com.hashicorp.cdktf.providers.google.provider.GoogleProvider;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import jp.ijufumi.sample.vision.api.deployment.stacks.ArtifactRegistryStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.BucketStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.CloudCDNStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.CloudRunStack;
import software.constructs.Construct;

public class MainStack extends TerraformStack {

  public MainStack(final Construct scope, final String id, final Config config) {
    super(scope, id);

    GoogleProvider
        .Builder
        .create(this, "gcp-provider")
        .region(config.Region())
        .project(config.ProjectId())
        .credentials(config.Credentials())
        .build();

    CloudRunStack.create(this, config);
    var bucket = BucketStack.create(this, config);
    ArtifactRegistryStack.create(this, config);
    CloudCDNStack.create(this, config, bucket);
  }
}
