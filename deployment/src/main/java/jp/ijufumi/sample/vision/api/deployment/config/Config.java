package jp.ijufumi.sample.vision.api.deployment.config;


import io.github.cdimascio.dotenv.Dotenv;

public class Config {

    private Dotenv dotenv;
    private Config(Dotenv dotenv) {
        this.dotenv = dotenv;
    }

    public static Config read() {
        Dotenv dotenv = Dotenv.load();

        return new Config(dotenv);
    }
}
