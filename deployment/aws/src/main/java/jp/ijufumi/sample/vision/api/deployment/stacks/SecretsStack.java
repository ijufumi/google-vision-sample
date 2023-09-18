package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.RemovalPolicy;
import software.amazon.awscdk.SecretValue;
import software.amazon.awscdk.services.secretsmanager.Secret;
import software.constructs.Construct;

public class SecretsStack {

  public static Secret build(final Construct scope, final Config config) {
    var secretValue = SecretValue.Builder.create(config.googleCredential()).build();
    return Secret
        .Builder
        .create(scope, "google-credential-secret")
        .removalPolicy(RemovalPolicy.DESTROY)
        .secretName("GOOGLE_CREDENTIAL")
        .secretStringValue(secretValue)
        .build();
  }
}
