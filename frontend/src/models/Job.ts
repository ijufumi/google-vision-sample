import ExtractedText, { Props as ExtractedTextProps } from "./ExtractedText"
import File, { Props as FilProps } from "./File"
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
  files: FilProps[]
}

export default class Job {
  readonly id: string
  readonly status: JobStatus
  readonly createdAt: number
  readonly updatedAt: number
  readonly extractedTexts: ExtractedText[]
  readonly files: File[]

  constructor(props: Props) {
    this.id = props.id
    this.status = props.status
    this.createdAt = props.createdAt
    this.updatedAt = props.updatedAt
    this.extractedTexts = props.extractedTexts.map(
      (p) => new ExtractedText(p)
    )
    this.files = props.files.map(f => new File(f))
  }

  get readableCreatedAt() {
    return formatToDate(this.createdAt)
  }

  get readableUpdatedAt() {
    return formatToDate(this.updatedAt)
  }
}
