package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.artifact_registry_repository.ArtifactRegistryRepository;
import com.hashicorp.cdktf.providers.google.artifact_registry_repository.ArtifactRegistryRepositoryConfig;
import com.hashicorp.cdktf.providers.google.artifact_registry_repository.ArtifactRegistryRepositoryDockerConfig;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class ArtifactRegistryStack {
    public static void create(final Construct scope, final Config config) {
        var dockerConfig = ArtifactRegistryRepositoryDockerConfig
                .builder()
                .immutableTags(true)
                .build();
        var registryConfig = ArtifactRegistryRepositoryConfig
                .builder()
                .location(config.Location())
                .repositoryId(config.RepositoryId())
                .dockerConfig(dockerConfig)
                .format("DOCKER")
                .build();
        new ArtifactRegistryRepository(scope, "container-registry", registryConfig);
    }
}
