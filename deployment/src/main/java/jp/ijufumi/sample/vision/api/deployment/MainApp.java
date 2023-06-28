package jp.ijufumi.sample.vision.api.deployment;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.App;
import software.amazon.awscdk.Environment;
import software.amazon.awscdk.StackProps;

public class MainApp {

  public static void main(final String[] args) {
    var app = new App();

    var config = Config.build();
    new MainStack(app, "DeploymentStack", StackProps.builder()
        .env(Environment.builder()
            .account(config.accountId())
            .region(config.region())
            .build())
        .build(), config);

    app.synth();
  }
}

