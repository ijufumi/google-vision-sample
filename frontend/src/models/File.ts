import { formatToDate } from "../components/dates"

export interface Props {
  id: string
  isOutput: boolean
  fileKey: string
  size: number
  contentType: string
  createdAt: number
  updatedAt: number
}

export default class File {
  readonly id: string
  readonly isOutput: boolean
  readonly fileKey: string
  readonly size: number
  readonly contentType: string
  readonly createdAt: number
  readonly updatedAt: number

  constructor(props: Props) {
    this.id = props.id
    this.isOutput = props.isOutput
    this.fileKey = props.fileKey
    this.size = props.size
    this.contentType = props.contentType
    this.createdAt = props.createdAt
    this.updatedAt = props.updatedAt
  }

  get readableCreatedAt() {
    return formatToDate(this.createdAt)
  }

  get readableUpdatedAt() {
    return formatToDate(this.updatedAt)
  }
}
