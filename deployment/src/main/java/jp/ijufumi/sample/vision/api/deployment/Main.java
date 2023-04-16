package jp.ijufumi.sample.vision.api.deployment;

import com.hashicorp.cdktf.App;


public class Main
{
    public static void main(String[] args) {
        final App app = new App();
        new MainStack(app, "deployment");
        app.synth();
    }
}
