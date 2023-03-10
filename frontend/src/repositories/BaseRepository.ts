import { UnexpectedError, ClientError, ServerError } from "../components/errors"

enum Methods {
  Get = "GET",
  Put = "PUT",
  Post = "POST",
  Delete = "DELETE",
}

abstract class BaseRepository {
  private readonly apiEndpoint: string

  constructor(apiEndpoint: string) {
    this.apiEndpoint = apiEndpoint
  }

  _get = async (args: {
    path: string
    auth?: boolean
    headers?: Record<string, string>
  }) => {
    const baseHeaders = { "Content-Type": "application/json;charset=utf-8" }
    return await this._request(
      Methods.Get,
      this.apiEndpoint,
      args.path,
      args.auth,
      undefined,
      Object.assign(baseHeaders, args.headers ? args.headers : {})
    )
  }

  _delete = async (args: {
    path: string
    auth?: boolean
    headers?: Record<string, string>
  }) => {
    const baseHeaders = { "Content-Type": "application/json;charset=utf-8" }
    return await this._request(
      Methods.Delete,
      this.apiEndpoint,
      args.path,
      args.auth,
      undefined,
      Object.assign(baseHeaders, args.headers ? args.headers : {})
    )
  }

  _post = async (args: {
    path: string
    auth?: boolean
    headers?: Record<string, string>
    body?: object
  }) => {
    const baseHeaders = { "Content-Type": "application/json;charset=utf-8" }
    return await this._request(
      Methods.Post,
      this.apiEndpoint,
      args.path,
      args.auth,
      args.body,
      Object.assign(baseHeaders, args.headers ? args.headers : {})
    )
  }

  _put = async (args: {
    path: string
    auth?: boolean
    headers?: Record<string, string>
    body?: object
  }) => {
    const baseHeaders = { "Content-Type": "application/json;charset=utf-8" }
    return await this._request(
      Methods.Put,
      this.apiEndpoint,
      args.path,
      args.auth,
      args.body,
      Object.assign(baseHeaders, args.headers ? args.headers : {})
    )
  }

  _download = async (args: {
    path: string
    auth?: boolean
    headers?: Record<string, string>
  }) => {
    return await this._request(
      Methods.Get,
      this.apiEndpoint,
      args.path,
      args.auth,
      undefined,
      args.headers,
      true
    )
  }

  _upload = async (args: {
    path: string
    auth?: boolean
    headers?: Record<string, string>
    body: FormData
  }) => {
    return await this._request(
      Methods.Post,
      this.apiEndpoint,
      args.path,
      args.auth,
      args.body,
      Object.assign({}, args.headers ? args.headers : {}),
      false,
      true
    )
  }

  _request = async (
    method: Methods,
    apiEndpoint: string,
    path: string,
    auth?: boolean,
    body?: object | FormData,
    header?: Record<string, string>,
    responseAsBlob?: boolean,
    requestAsForm?: boolean
  ) => {
    let bodyData = undefined
    if (body) {
      bodyData = (requestAsForm === true) ? (body as FormData) : JSON.stringify(body)
    }
    const response = await fetch(`${apiEndpoint}${path}`, {
      mode: "cors",
      cache: "no-cache",
      method: method.toString(),
      headers: Object.assign({}, header ? header : {}),
      body: bodyData,
    }).catch((e) => {
      throw new UnexpectedError(`No response error with ${e}`)
    })

    if (response.ok) {
      if (responseAsBlob === true) {
        return response.blob()
      }
      return response.json()
    }
    if (response.status < 500) {
      throw new ClientError(`ClientError with ${response}`)
    }
    throw new ServerError(`ServerError with ${response}`)
  }
}

export default BaseRepository
