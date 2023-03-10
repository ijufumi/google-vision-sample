import BaseRepository from "./BaseRepository"
import Job, { Props } from "../models/Job"
import Status from "../models/Status"

export interface JobRepository {
  getAll(): Promise<Job[]>

  getById(args: { id: string }): Promise<Job>

  create(args: { file: File }): Promise<Status>

  delete(args: { id: string }): Promise<Status>
}

export class JobRepositoryImpl
  extends BaseRepository
  implements JobRepository
{
  create = async (args: { file: File }) => {
    const formData = new FormData()
    formData.set("file", args.file)
    const result = await this._upload({ path: "", body: formData })
    return new Status(result)
  }

  getAll = async () => {
    const results = await this._get({ path: "" })
    return results.map((p: Props) => new Job(p))
  }

  getById = async (args: { id: string }) => {
    const result = await this._get({ path: `/${args.id}` })
    return new Job(result)
  }

  delete = async (args: { id: string }) => {
    const result = await this._delete({ path: `/${args.id}` })
    return new Status(result)
  }
}
