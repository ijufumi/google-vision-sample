package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2Service;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceConfig;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceTemplate;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceTemplateContainers;
import java.util.List;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class CloudRunStack {

  public static void create(final Construct scope, final Config config) {
    var container = CloudRunV2ServiceTemplateContainers
        .builder()
        .image(config.CloudRunContainerImage())
        .build();
    var template = CloudRunV2ServiceTemplate
        .builder()
        .containers(List.of(container))
        .build();
    var cloudRunConfig = CloudRunV2ServiceConfig
        .builder()
        .template(template)
        .build();
    new CloudRunV2Service(scope, "cloud-run", cloudRunConfig);
  }
}
