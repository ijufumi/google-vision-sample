import SignedUrl from "../models/SignedUrl"
import BaseRepository from "./BaseRepository"

export interface SignedUrlRepository {
  get(args: { key: string }): Promise<SignedUrl>
}

export default class SignedUrlRepositoryImpl
  extends BaseRepository
  implements SignedUrlRepository
{
  get = async (args: { key: string }) => {
    const path = `?key=${args.key}`
    const result = await this._get({ path })
    return new SignedUrl(result)
  }
}
