package jp.ijufumi.sample.vision.api.deployment;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import jp.ijufumi.sample.vision.api.deployment.stacks.CloudRunStack;
import software.constructs.Construct;
import com.hashicorp.cdktf.providers.google.provider.GoogleProvider;
import com.hashicorp.cdktf.TerraformStack;

public class MainStack extends TerraformStack
{
    public MainStack(final Construct scope, final String id, final Config config) {
        super(scope, id);

        GoogleProvider
                .Builder
                .create(this, "gcp-provider")
                .region(config.Region())
                .project(config.ProjectId())
                .build();
        CloudRunStack.create(this, config);
        // define resources here
    }
}
