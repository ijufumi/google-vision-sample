package jp.ijufumi.sample.vision.api.deployment;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import jp.ijufumi.sample.vision.api.deployment.stacks.APICloudfrontStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.ECRImageStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.ECSStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.Route53Stack;
import jp.ijufumi.sample.vision.api.deployment.stacks.S3Stack;
import jp.ijufumi.sample.vision.api.deployment.stacks.SecretsStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.VpcStack;
import jp.ijufumi.sample.vision.api.deployment.stacks.WebCloudfrontStack;
import software.amazon.awscdk.Stack;
import software.amazon.awscdk.StackProps;
import software.constructs.Construct;

public class MainStack extends Stack {

  public MainStack(final Construct scope, final String id, final StackProps props,
      final Config config) {
    super(scope, id, props);

    var vpc = VpcStack.build(this, config);
    var secret = SecretsStack.build(this, config);
    var bucket = S3Stack.build(this, config);
    var dockerImage = ECRImageStack.build(this, config);
    WebCloudfrontStack.build(this, bucket);
    var alb = ECSStack.build(this, config, vpc, dockerImage, secret);
    var cloudfront = APICloudfrontStack.build(this, bucket, alb);
    Route53Stack.build(this, config, cloudfront);
  }
}
