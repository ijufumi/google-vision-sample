package jp.ijufumi.sample.vision.api.deployment.config;

import io.github.cdimascio.dotenv.Dotenv;

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

  public String region() {
    var region = this.getEnv("CDK_DEFAULT_REGION");
    if (region != null) {
      return region;
    }
    return this.getEnv("AWS_DEFAULT_REGION");
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
