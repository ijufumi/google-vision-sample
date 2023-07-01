package jp.ijufumi.sample.vision.api.deployment;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import jp.ijufumi.sample.vision.api.deployment.stacks.CloudfrontStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.S3Stack;
import software.amazon.awscdk.core.Stack;
import software.amazon.awscdk.core.StackProps;
import software.constructs.Construct;

public class MainStack extends Stack {

  public MainStack(final Construct scope, final String id, final StackProps props,
      final Config config) {
    super(scope, id, props);

    var bucket = S3Stack.build(scope, config);
    CloudfrontStack.build(scope, config, bucket);
  }
}
