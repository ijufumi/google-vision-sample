import {Construct} from "constructs";
import * as google from "@cdktf/provider-google"
import {GCP_LOCATION, BUCKET_NAME, GCP_PROJECT} from "../configs/config";
export const createBucket = (scope: Construct) => {
  const cors = [
    {
      method: ["*"],
      origin: ["*"],
      maxAgeSeconds: 3600,
    } as google.storageBucket.StorageBucketCors
  ]
  new google.storageBucket.StorageBucket(scope, 'StorageBucket', {
    project: GCP_PROJECT,
    location: GCP_LOCATION,
    name: BUCKET_NAME,
    forceDestroy: true,
    cors,
  });
}
