import { formatToDate } from "../components/dates"

export interface Props {
  id: string
  jobID: string
  text: string
  top: number
  bottom: number
  left: number
  right: number
  createdAt: number
  updatedAt: number
}

export default class ExtractedText {
  readonly id: string
  readonly jobID: string
  readonly text: string
  readonly top: number
  readonly bottom: number
  readonly left: number
  readonly right: number
  readonly createdAt: number
  readonly updatedAt: number

  constructor(props: Props) {
    this.id = props.id
    this.jobID = props.jobID
    this.text = props.text
    this.top = props.top
    this.bottom = props.bottom
    this.left = props.left
    this.right = props.right
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
