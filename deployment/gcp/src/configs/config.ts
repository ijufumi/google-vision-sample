import * as fs from 'fs';
import * as dotenv from "dotenv"

dotenv.config()

export const GCP_PROJECT = process.env.GCP_PROJECT || "" as string

export const GCP_LOCATION = process.env.GCP_LOCATION || "" as string
export const BUCKET_NAME = process.env.BACKET_NAME || "sample-bucket"

export const GCP_CREDENTIALS = fs.readFileSync(process.env.CREDENTIAL_FILE || "" as string, "utf8")
