package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.sql_database.SqlDatabase;
import com.hashicorp.cdktf.providers.google.sql_database.SqlDatabaseConfig;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstance;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceConfig;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceSettings;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceSettingsIpConfiguration;
import com.hashicorp.cdktf.providers.google.sql_user.SqlUser;
import com.hashicorp.cdktf.providers.google.sql_user.SqlUserConfig;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class DatabaseStack {

  public static SqlDatabaseInstance create(final Construct scope, final Config config) {
    var ipConfiguration = SqlDatabaseInstanceSettingsIpConfiguration
        .builder()
        .ipv4Enabled(true)
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

    new SqlDatabase(scope, "sql-database", databaseConfig);

    return databaseInstance;
  }
}
