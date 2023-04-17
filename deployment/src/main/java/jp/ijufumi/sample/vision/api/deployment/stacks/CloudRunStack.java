package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2Service;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceConfig;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class CloudRunStack {
    public static CloudRunV2Service create(final Construct scope, final Config config) {
        var cloudRunConfig = CloudRunV2ServiceConfig.builder().build();
        return new CloudRunV2Service(scope, "cloud-run", cloudRunConfig);
    }
}
