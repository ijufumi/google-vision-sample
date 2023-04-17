package jp.ijufumi.sample.vision.api.deployment.config;


import io.github.cdimascio.dotenv.Dotenv;

public class Config {

    private final Dotenv dotenv;
    private Config(Dotenv dotenv) {
        this.dotenv = dotenv;
    }

    public String getEnv(String key, String defaultValue) {
        return this.dotenv.get(key, defaultValue);
    }

    public Integer getEnv(String key, Integer defaultValue) {
        try {
            return Integer.parseInt(this.dotenv.get(key));
        } catch (NumberFormatException e) {
            System.out.println(e);
            return defaultValue;
        }
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
