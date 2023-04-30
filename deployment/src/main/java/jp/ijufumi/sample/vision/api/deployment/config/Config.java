package jp.ijufumi.sample.vision.api.deployment.config;


import io.github.cdimascio.dotenv.Dotenv;
import java.io.IOException;
import java.io.UncheckedIOException;
import java.nio.file.Files;
import java.nio.file.Path;
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
      var value = this.dotenv.get(key);
      if (value == null || value.equals("")) {
        return defaultValue;
      }
      return Integer.parseInt(value);
    } catch (NumberFormatException e) {
      System.out.println(e);
      return defaultValue;
    }
  }

  public String Credentials() {
    var jsonString = this.CredentialsJson();
    if (jsonString == null || jsonString.equals("")) {
      try {
        var bytes = Files.readAllBytes(Path.of(this.CredentialsJsonFile()));
        return new String(bytes);
      } catch (IOException e) {
        throw new UncheckedIOException(e);
      }
    }
    return jsonString;
  }

  public String CredentialsJson() {
    return this.getEnv("CREDENTIALS_JSON", "");
  }

  public String CredentialsJsonFile() {
    return this.getEnv("CREDENTIALS_JSON_FILE", "");
  }

  public String ProjectId() {
    return this.getEnv("PROJECT_ID", "");
  }

  public String Region() {
    return this.getEnv("REGION", "");
  }

  public String Location() {
    return this.getEnv("LOCATION", "us-west1");
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

  public Integer BackendBucketCdnPolicyTTL() {
    return this.getEnv("BACKEND_BUCKET_CDN_POLICY_TTL", 3600);
  }

  public String BackendBucketName() {
    return this.getEnv("BACKEND_BUCKET_NAME", "backend-bucket-name");
  }

  public String CloudRunName() {
    return this.getEnv("CLOUD_RUN_NAME");
  }

  public String CloudRunContainerImage() {
    return this.getEnv("CLOUD_RUN_CONTAINER_IMAGE");
  }

  public Integer CloudRunContainerPort() {
    return this.getEnv("CLOUD_RUN_CONTAINER_PORT", 0);
  }

  public String CloudRunContainerProbePath() {
    return this.getEnv("CLOUD_RUN_CONTAINER_PROBE_PATH");
  }

  public String CloudRunV2ServiceIamMemberName() {
    return this.getEnv("CLOUD_RUN_V2_SERVICE_IAM_MEMBER_NAME", "cloud-run-v2-service-iam-member-name");
  }

  public Integer CloudRunContainerProbeSeconds() {
    return this.getEnv("CLOUD_RUN_CONTAINER_PROBE_SECOND", 100);
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
