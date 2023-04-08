import { formatToDate } from "../components/dates"
import ExtractedText, { Props as ExtractedTextProps } from "./ExtractedText"

export interface Props {
  id: string
  fileKey: string
  fileName: string
  size: number
  contentType: string
  createdAt: number
  updatedAt: number
  extractedTexts: ExtractedTextProps[] | undefined
}

export default class OutputFile {
  readonly id: string
  readonly fileKey: string
  readonly fileName: string
  readonly size: number
  readonly contentType: string
  readonly createdAt: number
  readonly updatedAt: number
  readonly extractedTexts: ExtractedText[]

  constructor(props: Props) {
    this.id = props.id
    this.fileKey = props.fileKey
    this.fileName = props.fileName
    this.size = props.size
    this.contentType = props.contentType
    this.createdAt = props.createdAt
    this.updatedAt = props.updatedAt
    this.extractedTexts = props.extractedTexts?.map(
      (p) => new ExtractedText(p)
    ) || []
  }

  get readableCreatedAt() {
    return formatToDate(this.createdAt)
  }

  get readableUpdatedAt() {
    return formatToDate(this.updatedAt)
  }
}
