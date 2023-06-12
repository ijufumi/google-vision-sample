package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucket;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketCdnPolicy;
import com.hashicorp.cdktf.providers.google.compute_backend_bucket.ComputeBackendBucketConfig;
import com.hashicorp.cdktf.providers.google.compute_global_address.ComputeGlobalAddress;
import com.hashicorp.cdktf.providers.google.compute_global_address.ComputeGlobalAddressConfig;
import com.hashicorp.cdktf.providers.google.compute_global_forwarding_rule.ComputeGlobalForwardingRule;
import com.hashicorp.cdktf.providers.google.compute_global_forwarding_rule.ComputeGlobalForwardingRuleConfig;
import com.hashicorp.cdktf.providers.google.compute_target_http_proxy.ComputeTargetHttpProxy;
import com.hashicorp.cdktf.providers.google.compute_target_http_proxy.ComputeTargetHttpProxyConfig;
import com.hashicorp.cdktf.providers.google.compute_url_map.ComputeUrlMap;
import com.hashicorp.cdktf.providers.google.compute_url_map.ComputeUrlMapConfig;
import com.hashicorp.cdktf.providers.google.storage_bucket.StorageBucket;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class CloudCDNStack {

  public static void create(final Construct scope, final Config config,
      final StorageBucket bucket) {
    var backendBucketCdnPolicy = ComputeBackendBucketCdnPolicy
        .builder()
        .cacheMode("CACHE_ALL_STATIC")
        .maxTtl(config.BackendBucketCdnPolicyTTL())
        .clientTtl(config.BackendBucketCdnPolicyTTL())
        .defaultTtl(config.BackendBucketCdnPolicyTTL())
        .build();
    var backendBucketConfig = ComputeBackendBucketConfig
        .builder()
        .bucketName(bucket.getName())
        .enableCdn(true)
        .cdnPolicy(backendBucketCdnPolicy)
        .compressionMode("AUTOMATIC")
        .name(config.BackendBucketName())
        .build();
    var backendBucket = new ComputeBackendBucket(scope, "backend-bucket", backendBucketConfig);

    var urlMapConfig = ComputeUrlMapConfig
        .builder()
        .defaultService(backendBucket.getId())
        .name("url-loadbalancer")
        .build();
    var urlMap = new ComputeUrlMap(scope, "compute-url-map", urlMapConfig);

    var globalAddressConfig = ComputeGlobalAddressConfig
        .builder()
        .project(config.ProjectId())
        .name("ip-loadbalancer")
        .build();
    var globalAddress = new ComputeGlobalAddress(scope, "cdn-compute-global-address",
        globalAddressConfig);

    var httpProxyConfig = ComputeTargetHttpProxyConfig
        .builder()
        .name("proxy-loadbalancer")
        .urlMap(urlMap.getId())
        .build();
    var httpProxy = new ComputeTargetHttpProxy(scope, "compute-target-http-proxy", httpProxyConfig);

    var globalForwardingRuleConfig = ComputeGlobalForwardingRuleConfig
        .builder()
        .name("forwarding-rule-loadbalancer")
        .ipProtocol("TCP")
        .loadBalancingScheme("EXTERNAL")
        .ipAddress(globalAddress.getId())
        .target(httpProxy.getId())
        .portRange("443")
        .build();
    new ComputeGlobalForwardingRule(scope, "compute-global-forwarding-rule",
        globalForwardingRuleConfig);

  }
}
