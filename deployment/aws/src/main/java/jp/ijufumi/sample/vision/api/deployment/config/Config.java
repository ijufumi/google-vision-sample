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

  String vpcName();

  String certificationArn();

  String bucket();

  String repositoryName();

  String backendCode();

  String googleCredential();

  String hostZoneId();

  String hostZoneName();

  String apiDomainName();

  String webDomainName();

  String apiDomainFullName();

  String webDomainFullName();

  String dbName();

  String dbHost();

  String dbUser();

  String dbPassword();

  int dbPort();

  String route53Namespace();

  String googleBucketName();
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
  public String vpcName() {
    return this.getEnv("VPC_NAME");
  }

  @Override
  public String certificationArn() {
    return this.getEnv("CERTIFICATION_ARN");
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
      var filePath = Path.of(credentialFilePath);
      if (Files.exists(filePath)) {
        try {
          return Files.readString(filePath);
        } catch (IOException e) {
          e.printStackTrace();
        }
      }
      try {
        var resource = getClass().getClassLoader().getResource(credentialFilePath);
        if (resource != null) {
          return Files.readString(Path.of(resource.getPath()));
        }
      } catch (IOException e) {
        e.printStackTrace();
      }
    }

    return this.getEnv("APP_GOOGLE_CREDENTIAL");
  }

  @Override
  public String hostZoneId() {
    return this.getEnv("HOST_ZONE_ID");
  }

  @Override
  public String hostZoneName() {
    return this.getEnv("HOST_ZONE_NAME");
  }

  @Override
  public String apiDomainName() {
    return this.getEnv("API_DOMAIN_NAME");
  }

  @Override
  public String webDomainName() {
    return this.getEnv("WEB_DOMAIN_NAME");
  }

  @Override
  public String apiDomainFullName() {
    return String.format("%s.%s", apiDomainName(), hostZoneName());
  }

  @Override
  public String webDomainFullName() {
    return String.format("%s.%s", webDomainName(), hostZoneName());
  }

  @Override
  public String dbName() {
    return this.getEnv("APP_DB_NAME");
  }

  @Override
  public String dbHost() {
    return this.getEnv("APP_DB_HOST");
  }

  @Override
  public String dbUser() {
    return this.getEnv("APP_DB_USER");
  }

  @Override
  public String dbPassword() {
    return this.getEnv("APP_DB_PASSWORD");
  }

  @Override
  public int dbPort() {
    return this.getEnv("APP_DB_PORT", 5432);
  }

  @Override
  public String route53Namespace() {
    return this.getEnv("ROUTE53_NAMESPACE");
  }

  @Override
  public String googleBucketName() {
    return this.getEnv("APP_GOOGLE_STORAGE_BUCKET");
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
