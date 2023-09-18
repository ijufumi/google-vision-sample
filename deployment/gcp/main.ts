import { Construct } from "constructs";
import { App, TerraformStack } from "cdktf";
import {createBucket} from "./src/stacks/bucket";
import * as google from "@cdktf/provider-google"
import {CREDENTIALS, PROJECT} from "./src/configs/config";

class MyStack extends TerraformStack {
  constructor(scope: Construct, id: string) {
    super(scope, id);
    new google.provider.GoogleProvider(this, "google", {
      project: PROJECT,
      credentials: CREDENTIALS,
    })
    createBucket(this)
    // define resources here
  }
}

const app = new App();
new MyStack(app, "gcp");
app.synth();
