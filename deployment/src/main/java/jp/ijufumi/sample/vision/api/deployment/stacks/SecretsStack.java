package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.core.RemovalPolicy;
import software.amazon.awscdk.core.SecretValue;
import software.amazon.awscdk.services.secretsmanager.Secret;
import software.amazon.awscdk.services.secretsmanager.SecretStringGenerator;
import software.constructs.Construct;

public class SecretsStack {

  public static void build(final Construct scope, final Config config) {
    var secretValue = SecretValue.Builder.create(config.googleCredential()).build();
    Secret
        .Builder
        .create(scope, "google-credential-secret")
        .removalPolicy(RemovalPolicy.DESTROY)
        .secretName("GOOGLE_CREDENTIAL")
        .secretStringValue(secretValue)
        .generateSecretString(SecretStringGenerator.builder().build())
        .build();
  }
}
