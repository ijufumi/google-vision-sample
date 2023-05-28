package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.compute_network.ComputeNetwork;
import com.hashicorp.cdktf.providers.google.compute_network.ComputeNetworkConfig;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstance;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceConfig;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceSettings;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceSettingsIpConfiguration;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class DatabaseStack {

  public static SqlDatabaseInstance create(final Construct scope, final Config config) {
    var networkConfig = ComputeNetworkConfig
        .builder()
        .project(config.ProjectId())
        .name("private-network")
        .build();

    var network = new ComputeNetwork(scope, "private-network", networkConfig);

    var ipConfiguration = SqlDatabaseInstanceSettingsIpConfiguration
        .builder()
        .ipv4Enabled(false)
        .enablePrivatePathForGoogleCloudServices(true)
        .privateNetwork(network.getId())
        .build();

    var databaseSetting = SqlDatabaseInstanceSettings
        .builder()
        .tier("db-f1-micro")
        .ipConfiguration(ipConfiguration)
        .build();

    var databaseConfig = SqlDatabaseInstanceConfig
        .builder()
        .project(config.ProjectId())
        .region(config.Region())
        .databaseVersion("POSTGRES_14")
        .settings(databaseSetting)
        .deletionProtection(false)
        .build();

    return new SqlDatabaseInstance(scope, "sql-database-instance", databaseConfig);
  }
}
