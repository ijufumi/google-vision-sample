package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.project_service.ProjectService;
import com.hashicorp.cdktf.providers.google.project_service.ProjectServiceConfig;
import com.hashicorp.cdktf.providers.google.secret_manager_secret.SecretManagerSecret;
import com.hashicorp.cdktf.providers.google.secret_manager_secret.SecretManagerSecretConfig;
import com.hashicorp.cdktf.providers.google.secret_manager_secret.SecretManagerSecretReplication;
import com.hashicorp.cdktf.providers.google.secret_manager_secret_version.SecretManagerSecretVersion;
import com.hashicorp.cdktf.providers.google.secret_manager_secret_version.SecretManagerSecretVersionConfig;
import java.util.List;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class SecretManagerStack {

  public static SecretManagerSecretVersion create(final Construct scope, final Config config) {
    var serviceConfig = ProjectServiceConfig
        .builder()
        .project(config.ProjectId())
        .service("secretmanager.googleapis.com")
        .build();
    var service = new ProjectService(scope, "project-service", serviceConfig);
    var replication = SecretManagerSecretReplication.builder().automatic(true).build();
    var secretConfig = SecretManagerSecretConfig
        .builder()
        .project(config.ProjectId())
        .secretId("google-credential")
        .dependsOn(List.of(service))
        .replication(replication)
        .build();
    var secret = new SecretManagerSecret(scope, "secret-credential", secretConfig);
    var secretVersionConfig = SecretManagerSecretVersionConfig
        .builder()
        .secret(secret.getSecretId())
        .secretData(config.Credentials())
        .enabled(true)
        .build();
    return new SecretManagerSecretVersion(scope, "secret-version-credential", secretVersionConfig);
  }
}
