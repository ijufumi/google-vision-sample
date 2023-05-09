package jp.ijufumi.sample.vision.api.deployment.stacks;

import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstance;
import com.hashicorp.cdktf.providers.google.sql_database_instance.SqlDatabaseInstanceConfig;
import jp.ijufumi.sample.vision.api.deployment.config.Config;
import software.constructs.Construct;

public class DatabaseStack {

  public static SqlDatabaseInstance create(final Construct scope, final Config config) {
    var databaseConfig = SqlDatabaseInstanceConfig
        .builder()
        .project(config.ProjectId())
        .region(config.Region())
        .databaseVersion("POSTGRES_14")
        .build();

    return new SqlDatabaseInstance(scope, "sql-database-instance", databaseConfig);
  }
}
