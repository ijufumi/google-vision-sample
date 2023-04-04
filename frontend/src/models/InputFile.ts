import { formatToDate } from "../components/dates"
import OutputFile, { Props as OutputFileProp } from "./OutputFile"

export interface Props {
  id: string
  fileKey: string
  fileName: string
  size: number
  contentType: string
  createdAt: number
  updatedAt: number
  outputFiles: OutputFileProp[]
}

export default class InputFile {
  readonly id: string
  readonly fileKey: string
  readonly fileName: string
  readonly size: number
  readonly contentType: string
  readonly createdAt: number
  readonly updatedAt: number
  readonly outputFiles: OutputFile[]

  constructor(props: Props) {
    this.id = props.id
    this.fileKey = props.fileKey
    this.fileName = props.fileName
    this.size = props.size
    this.contentType = props.contentType
    this.createdAt = props.createdAt
    this.updatedAt = props.updatedAt
    this.outputFiles = props.outputFiles.map(file => new OutputFile(file))
  }

  get readableCreatedAt() {
    return formatToDate(this.createdAt)
  }

  get readableUpdatedAt() {
    return formatToDate(this.updatedAt)
  }

  get isJSON() {
    return "application/json" === this.contentType
  }

  get isImage() {
    return this.contentType.startsWith("image/")
  }
}
