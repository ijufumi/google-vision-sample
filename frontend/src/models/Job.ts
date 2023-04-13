import InputFile, { Props as InputFileProps } from "./InputFile"
import { formatToDateString } from "../components/dates"

export enum JobStatus {
  Running = "running",
  Succeeded = "succeeded",
  Failed = "failed",
}

export interface Props {
  id: string
  name: string
  status: JobStatus
  createdAt: number
  updatedAt: number
  inputFiles: InputFileProps[] | undefined
}

export default class Job {
  readonly id: string
  readonly name: string
  readonly status: JobStatus
  readonly createdAt: number
  readonly updatedAt: number
  readonly inputFiles: InputFile[]

  constructor(props: Props) {
    this.id = props.id
    this.name = props.name
    this.status = props.status
    this.createdAt = props.createdAt
    this.updatedAt = props.updatedAt
    this.inputFiles = props.inputFiles?.map(f => new InputFile(f)) || []
  }

  get readableCreatedAt() {
    return formatToDateString(this.createdAt)
  }

  get readableUpdatedAt() {
    return formatToDateString(this.updatedAt)
  }

  get inputFile() {
    return this.inputFiles.find(() => true)
  }

  get outputFile() {
    return this.inputFiles.find(() => true)?.outputFiles.find(() => true)
  }
}
