package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.artifact_registry_repository.ArtifactRegistryRepository;
import com.hashicorp.cdktf.providers.google.artifact_registry_repository.ArtifactRegistryRepositoryConfig;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class ArtifactRegistryStack {
    public static void create(final Construct scope, final Config config) {
        var registryConfig = ArtifactRegistryRepositoryConfig
                .builder()
                .build();
        new ArtifactRegistryRepository(scope, "container-registry", registryConfig);
    }
}
