export interface Props {
  url: string
}

export default class SignedUrl {
  readonly url: string

  constructor(props: Props) {
    this.url = props.url
  }
}
