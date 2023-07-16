package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.services.elasticloadbalancingv2.ApplicationLoadBalancer;
import software.amazon.awscdk.services.route53.HostedZone;
import software.amazon.awscdk.services.route53.RecordSet;
import software.amazon.awscdk.services.route53.RecordTarget;
import software.constructs.Construct;

public class Route53Stack {

  public static void build(final Construct scope, final Config config, final
  ApplicationLoadBalancer alb) {
    var hostZone = HostedZone
        .fromHostedZoneId(scope, "host-zone-id", config.hostZoneId());
    var recordTarget = RecordTarget.fromValues(alb.getLoadBalancerDnsName());
    RecordSet
        .Builder
        .create(scope, "record-set")
        .recordName(config.apiDomainName())
        .zone(hostZone)
        .target(recordTarget)
        .build();
  }
}
