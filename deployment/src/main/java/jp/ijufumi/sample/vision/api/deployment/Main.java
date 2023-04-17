package jp.ijufumi.sample.vision.api.deployment;

import com.hashicorp.cdktf.App;
import jp.ijufumi.sample.vision.api.deployment.config.Config;


public class Main
{
    public static void main(String[] args) {
        final var app = new App();
        final var config = Config.read();
        new MainStack(app, "deployment", config);
        app.synth();
    }
}
