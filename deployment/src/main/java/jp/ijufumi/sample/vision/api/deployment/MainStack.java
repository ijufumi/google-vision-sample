package jp.ijufumi.sample.vision.api.deployment;

import com.hashicorp.cdktf.TerraformStack;
import com.hashicorp.cdktf.providers.aws.provider.AwsProvider;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class MainStack extends TerraformStack {

  public MainStack(final Construct scope, final String id, final Config config) {
    super(scope, id);

    AwsProvider
        .Builder
        .create(this, "gcp-provider")
        .build();
  }
}
