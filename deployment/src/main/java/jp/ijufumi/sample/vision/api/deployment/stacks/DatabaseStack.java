package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.compute_global_address.ComputeGlobalAddress;
import com.hashicorp.cdktf.providers.google.compute_global_address.ComputeGlobalAddressConfig;
import com.hashicorp.cdktf.providers.google.compute_network.ComputeNetwork;
import com.hashicorp.cdktf.providers.google.compute_network.ComputeNetworkConfig;
import com.hashicorp.cdktf.providers.google.service_networking_connection.ServiceNetworkingConnection;
import com.hashicorp.cdktf.providers.google.service_networking_connection.ServiceNetworkingConnectionConfig;
import com.hashicorp.cdktf.providers.google.sql_database.SqlDatabase;
import com.hashicorp.cdktf.providers.google.sql_database.SqlDatabaseConfig;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstance;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceConfig;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceSettings;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceSettingsIpConfiguration;
import com.hashicorp.cdktf.providers.google.sql_user.SqlUser;
import com.hashicorp.cdktf.providers.google.sql_user.SqlUserConfig;
import java.util.List;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class DatabaseStack {

  public static SqlDatabase create(final Construct scope, final Config config) {
    var networkConfig = ComputeNetworkConfig
        .builder()
        .project(config.ProjectId())
        .name("private-network")
        .build();

    var network = new ComputeNetwork(scope, "private-network", networkConfig);

    var computeGlobalAddressConfig = ComputeGlobalAddressConfig
        .builder()
        .name("global-access")
        .purpose("VPC_PEERING")
        .addressType("INTERNAL")
        .prefixLength(16)
        .network(network.getId())
        .build();

    var computeGlobalAddress = new ComputeGlobalAddress(scope, "db-compute-global-address",
        computeGlobalAddressConfig);

    var serviceNetworkingConnectionConfig = ServiceNetworkingConnectionConfig
        .builder()
        .network(network.getId())
        .service("servicenetworking.googleapis.com")
        .reservedPeeringRanges(List.of(computeGlobalAddress.getName()))
        .build();

    var serviceNetworkingConnection = new ServiceNetworkingConnection(scope,
        "service-networking-connection", serviceNetworkingConnectionConfig);

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

    var databaseInstanceConfig = SqlDatabaseInstanceConfig
        .builder()
        .project(config.ProjectId())
        .region(config.Region())
        .databaseVersion("POSTGRES_14")
        .settings(databaseSetting)
        .deletionProtection(false)
        .dependsOn(List.of(serviceNetworkingConnection))
        .build();

    var databaseInstance = new SqlDatabaseInstance(scope, "sql-database-instance",
        databaseInstanceConfig);

    var userConfig = SqlUserConfig
        .builder()
        .name(config.AppDbUser())
        .password(config.AppDbPassword())
        .instance(databaseInstance.getId())
        .build();
    new SqlUser(scope, "sql-user", userConfig);

    var databaseConfig = SqlDatabaseConfig
        .builder()
        .name(config.AppDbName())
        .instance(databaseInstance.getId())
        .build();

    return new SqlDatabase(scope, "sql-database", databaseConfig);
  }
}
