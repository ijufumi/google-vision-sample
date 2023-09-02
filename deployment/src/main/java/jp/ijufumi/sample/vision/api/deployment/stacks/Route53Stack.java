package jp.ijufumi.sample.vision.api.deployment.stacks;

import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.amazon.awscdk.services.cloudfront.Distribution;
import software.amazon.awscdk.services.route53.HostedZone;
import software.amazon.awscdk.services.route53.HostedZoneAttributes;
import software.amazon.awscdk.services.route53.RecordSet;
import software.amazon.awscdk.services.route53.RecordTarget;
import software.amazon.awscdk.services.route53.RecordType;
import software.constructs.Construct;

public class Route53Stack {

  public static void build(final Construct scope, final Config config, final
  Distribution apiCloudFront) {
    var hostZoneAttribute = HostedZoneAttributes
        .builder()
        .hostedZoneId(config.hostZoneId())
        .zoneName(config.hostZoneName())
        .build();
    var hostZone = HostedZone
        .fromHostedZoneAttributes(scope, "host-zone", hostZoneAttribute);
    var recordTarget = RecordTarget.fromValues(apiCloudFront.getDistributionDomainName());
    RecordSet
        .Builder
        .create(scope, "record-set")
        .recordName(config.apiDomainName())
        .zone(hostZone)
        .target(recordTarget)
        .recordType(RecordType.CNAME)
        .build();
  }
}
