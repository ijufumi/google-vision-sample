import ExtractedText, { Props as ExtractedTextProps } from "./ExtractedText"
import JobFile, { Props as JobFileProps } from "./JobFile"
import { formatToDate } from "../components/dates"

export enum JobStatus {
  Running = "running",
  Succeeded = "succeeded",
  Failed = "failed",
}

export interface Props {
  id: string
  status: JobStatus
  createdAt: number
  updatedAt: number
  extractedTexts: ExtractedTextProps[]
  jobFiles: JobFileProps[]
}

export default class Job {
  readonly id: string
  readonly status: JobStatus
  readonly createdAt: number
  readonly updatedAt: number
  readonly extractedTexts: ExtractedText[]
  readonly jobFiles: JobFile[]

  constructor(props: Props) {
    this.id = props.id
    this.status = props.status
    this.createdAt = props.createdAt
    this.updatedAt = props.updatedAt
    this.extractedTexts = props.extractedTexts.map(
      (p) => new ExtractedText(p)
    )
    this.jobFiles = props.jobFiles.map(f => new JobFile(f))
  }

  get readableCreatedAt() {
    return formatToDate(this.createdAt)
  }

  get readableUpdatedAt() {
    return formatToDate(this.updatedAt)
  }

  get inputFile() {
    return this.jobFiles.find(f => !f.isOutput)
  }

  get outputFile() {
    return this.jobFiles.find(f => f.isOutput)
  }
}
