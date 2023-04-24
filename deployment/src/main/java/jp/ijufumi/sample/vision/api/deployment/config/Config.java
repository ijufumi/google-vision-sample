package jp.ijufumi.sample.vision.api.deployment.config;


import io.github.cdimascio.dotenv.Dotenv;
import java.util.List;

public class Config {

  private final Dotenv dotenv;

  private Config(Dotenv dotenv) {
    this.dotenv = dotenv;
  }

  public String getEnv(String key, String defaultValue) {
    return this.dotenv.get(key, defaultValue);
  }

  public String getEnv(String key) {
    return this.getEnv(key, (String) null);
  }

  public Integer getEnv(String key, Integer defaultValue) {
    try {
      return Integer.parseInt(this.dotenv.get(key));
    } catch (NumberFormatException e) {
      System.out.println(e);
      return defaultValue;
    }
  }
  
  public String ProjectId() {
    return this.getEnv("PROJECT_ID", "");
  }

  public String Region() {
    return this.getEnv("REGION", "");
  }

  public String Location() {
    return this.getEnv("LOCATION", "");
  }

  public String BucketName() {
    return this.getEnv("BUCKET_NAME", "");
  }

  public List<String> BucketCorsMethods() {
    return List.of(this.getEnv("BUCKET_CORS_METHODS", "*").split(","));
  }

  public List<String> BucketCorsOrigins() {
    return List.of(this.getEnv("BUCKET_CORS_ORIGINS", "*").split(","));
  }

  public Integer BucketCorsMaxAge() {
    return this.getEnv("BUCKET_CORS_MAX_AGE", 3600);
  }

  public String CloudRunContainerImage() {
    return this.getEnv("CLOUD_RUN_CONTAINER_IMAGE");
  }

  public Integer CloudRunContainerPort() {
    return this.getEnv("CLOUD_RUN_CONTAINER_PORT", 0);
  }

  public String RepositoryId() {
    return this.getEnv("REPOSITORY_ID");
  }

  public static Config read() {
    var dotenv = Dotenv
        .configure()
        .systemProperties()
        .ignoreIfMalformed()
        .ignoreIfMissing()
        .load();

    return new Config(dotenv);
  }
}
