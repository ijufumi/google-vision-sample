package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2Service;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceConfig;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceTemplate;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceTemplateContainers;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceTemplateContainersPorts;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceTemplateContainersStartupProbe;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service.CloudRunV2ServiceTemplateContainersStartupProbeHttpGet;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service_iam_member.CloudRunV2ServiceIamMember;
import com.hashicorp.cdktf.providers.google.cloud_run_v2_service_iam_member.CloudRunV2ServiceIamMemberConfig;
import java.util.List;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class CloudRunStack {

  public static void create(final Construct scope, final Config config) {
    var containerPort = CloudRunV2ServiceTemplateContainersPorts
        .builder()
        .containerPort(config.CloudRunContainerPort())
        .build();
    var startupProbeGet = CloudRunV2ServiceTemplateContainersStartupProbeHttpGet
        .builder()
        .path(config.CloudRunContainerProbePath())
        .build();
    var startupProbe = CloudRunV2ServiceTemplateContainersStartupProbe
        .builder()
        .httpGet(startupProbeGet)
        .periodSeconds(config.CloudRunContainerProbeSeconds())
        .timeoutSeconds(config.CloudRunContainerProbeSeconds())
        .build();
    var container = CloudRunV2ServiceTemplateContainers
        .builder()
        .image(config.CloudRunContainerImage())
        .ports(List.of(containerPort))
        .startupProbe(startupProbe)
        .build();
    var template = CloudRunV2ServiceTemplate
        .builder()
        .containers(List.of(container))
        .build();
    var cloudRunConfig = CloudRunV2ServiceConfig
        .builder()
        .template(template)
        .name(config.CloudRunName())
        .location(config.Location())
        .build();
    var cloudRun = new CloudRunV2Service(scope, "cloud-run", cloudRunConfig);

    var memberConfig = CloudRunV2ServiceIamMemberConfig
        .builder()
        .project(cloudRun.getProject())
        .location(config.Location())
        .name("")
        .build();
    new CloudRunV2ServiceIamMember(scope, "cloud-run-v2-service-iam-member", memberConfig);
  }
}
