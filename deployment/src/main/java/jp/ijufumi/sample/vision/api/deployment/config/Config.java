package jp.ijufumi.sample.vision.api.deployment.config;

import io.github.cdimascio.dotenv.Dotenv;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public interface Config {

  static Config build() {
    var dotenv = Dotenv
        .configure()
        .systemProperties()
        .ignoreIfMalformed()
        .ignoreIfMissing()
        .load();

    return new ConfigObj(dotenv);
  }

  String accountId();

  String region();

  String bucket();

  String repositoryName();

  String backendCode();

  String googleCredential();

  String hostZoneId();

  String apiDomainName();
}

class ConfigObj implements Config {

  @Override
  public String accountId() {
    var accountId = this.getEnv("CDK_DEFAULT_ACCOUNT");
    if (accountId != null) {
      return accountId;
    }
    return this.getEnv("AWS_ACCOUNT_ID");
  }

  @Override
  public String region() {
    var region = this.getEnv("CDK_DEFAULT_REGION");
    if (region != null) {
      return region;
    }
    return this.getEnv("AWS_DEFAULT_REGION");
  }

  @Override
  public String bucket() {
    return this.getEnv("S3_BUCKET_NAME");
  }

  @Override
  public String repositoryName() {
    return this.getEnv("ECR_REPOSITORY_NAME");
  }

  @Override
  public String backendCode() {
    return this.getEnv("BACKEND_CODE_DIRECTORY");
  }

  @Override
  public String googleCredential() {
    var credentialFilePath = this.getEnv("APP_GOOGLE_CREDENTIAL_FILE");
    if (credentialFilePath != null && !credentialFilePath.equals("")) {
      try {
        return Files.readString(Path.of(credentialFilePath));
      } catch (IOException e) {
        e.printStackTrace();
      }
    }

    return this.getEnv("APP_GOOGLE_CREDENTIAL");
  }

  @Override
  public String hostZoneId() {
    return this.getEnv("APP_HOST_ZONE_ID");
  }

  @Override
  public String apiDomainName() {
    return this.getEnv("APP_API_DOMAIN_NAME");
  }

  private final Dotenv dotenv;

  ConfigObj(Dotenv dotenv) {
    this.dotenv = dotenv;
  }

  String getEnv(String key, String defaultValue) {
    return this.dotenv.get(key, defaultValue);
  }

  String getEnv(String key) {
    return this.getEnv(key, (String) null);
  }

  Integer getEnv(String key, Integer defaultValue) {
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
}
