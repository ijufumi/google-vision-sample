package jp.ijufumi.sample.vision.api.deployment;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import jp.ijufumi.sample.vision.api.deployment.stacks.CloudfrontStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.ECRImageStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.ECSStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.Route53Stack;
import jp.ijufumi.sample.vision.api.deployment.stacks.S3Stack;
import jp.ijufumi.sample.vision.api.deployment.stacks.SecretsStack;
import software.amazon.awscdk.core.Stack;
import software.amazon.awscdk.core.StackProps;
import software.constructs.Construct;

public class MainStack extends Stack {

  public MainStack(final Construct scope, final String id, final StackProps props,
      final Config config) {
    super(scope, id, props);

    var secret = SecretsStack.build(scope, config);
    var bucket = S3Stack.build(scope, config);
    var dockerImage = ECRImageStack.build(scope, config);
    CloudfrontStack.build(scope, bucket);
    var alb = ECSStack.build(scope, config, dockerImage, secret);
    Route53Stack.build(scope, config, alb);
  }
}
