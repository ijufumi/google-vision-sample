import BaseRepository from "./BaseRepository"
import ExtractionResult, { Props } from "../models/ExtractionResult"
import Status from "../models/Status"

export interface ExtractionResultRepository {
  getAll(): Promise<ExtractionResult[]>

  getById(args: { id: string }): Promise<ExtractionResult>

  create(args: { file: File }): Promise<Status>

  delete(args: { id: string }): Promise<Status>
}

export class ExtractionResultRepositoryImpl
  extends BaseRepository
  implements ExtractionResultRepository
{
  create = async (args: { file: File }) => {
    const formData = new FormData()
    formData.set("file", args.file)
    const result = await this._upload({ path: "", body: formData })
    return new Status(result)
  }

  getAll = async () => {
    const results = await this._get({ path: "" })
    return results.map((p: Props) => new ExtractionResult(p))
  }

  getById = async (args: { id: string }) => {
    const result = await this._get({ path: `/${args.id}` })
    return new ExtractionResult(result)
  }

  delete = async (args: { id: string }) => {
    const result = await this._delete({ path: `/${args.id}` })
    return new Status(result)
  }
}
