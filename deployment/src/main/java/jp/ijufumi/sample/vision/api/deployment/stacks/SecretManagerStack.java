package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.secret_manager_secret.SecretManagerSecret;
import com.hashicorp.cdktf.providers.google.secret_manager_secret.SecretManagerSecretConfig;
import com.hashicorp.cdktf.providers.google.secret_manager_secret_version.SecretManagerSecretVersion;
import com.hashicorp.cdktf.providers.google.secret_manager_secret_version.SecretManagerSecretVersionConfig;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class SecretManagerStack {

  public static SecretManagerSecretVersion create(final Construct scope, final Config config) {
    var secretConfig = SecretManagerSecretConfig
        .builder()
        .project(config.ProjectId())
        .secretId("google-credential")
        .build();
    var secret = new SecretManagerSecret(scope, "secret-credential", secretConfig);
    var secretVersionConfig = SecretManagerSecretVersionConfig
        .builder()
        .secret(secret.getSecretId())
        .build();
    return new SecretManagerSecretVersion(scope, "secret-version-credential", secretVersionConfig);
  }
}
