package jp.ijufumi.sample.vision.api.deployment;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.core.Stack;
import software.amazon.awscdk.core.StackProps;
import software.constructs.Construct;
// import software.amazon.awscdk.Duration;
// import software.amazon.awscdk.services.sqs.Queue;

public class MainStack extends Stack {

  public MainStack(final Construct scope, final String id, final StackProps props,
      final Config config) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    // final Queue queue = Queue.Builder.create(this, "DeploymentQueue")
    //         .visibilityTimeout(Duration.seconds(300))
    //         .build();
  }
}
