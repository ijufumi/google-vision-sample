import InputFile, { Props as InputFileProps } from "./InputFile"
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
  inputFiles: InputFileProps[]
}

export default class Job {
  readonly id: string
  readonly status: JobStatus
  readonly createdAt: number
  readonly updatedAt: number
  readonly inputFiles: InputFile[]

  constructor(props: Props) {
    this.id = props.id
    this.status = props.status
    this.createdAt = props.createdAt
    this.updatedAt = props.updatedAt
    this.inputFiles = props.inputFiles.map(f => new InputFile(f))
  }

  get readableCreatedAt() {
    return formatToDate(this.createdAt)
  }

  get readableUpdatedAt() {
    return formatToDate(this.updatedAt)
  }

  get inputFile() {
    return this.inputFiles.find(() => true)
  }

  get outputFile() {
    return this.inputFiles.find(() => true)?.outputFiles.find(() => true)
  }
}
