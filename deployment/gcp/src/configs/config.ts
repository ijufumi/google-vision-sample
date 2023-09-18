import * as fs from 'fs';
import * as dotenv from "dotenv"

dotenv.config()

export const PROJECT = process.env.PROJECT || "" as string

export const LOCATION = process.env.LOCATION || "" as string
export const BUCKET_NAME = process.env.BUCKET_NAME || "sample-bucket"

export const CREDENTIALS = fs.readFileSync(process.env.CREDENTIAL_FILE || "" as string, "utf8")
