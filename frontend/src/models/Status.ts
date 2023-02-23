export interface Props {
  status: boolean
}

export default class Status {
  readonly status: boolean

  constructor(props: Props) {
    this.status = props.status
  }
}
