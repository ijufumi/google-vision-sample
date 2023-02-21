import BaseRepository from "./BaseRepository";
import ExtractionResult, { Props } from "../models/ExtractionResult";

export interface ExtractionResultRepository {
    getAll(): Promise<ExtractionResult[]>

    getById(args: { id: string }): Promise<ExtractionResult>

    create(args: { file: File }): Promise<void>

    delete(args: { id: string }): Promise<void>
}

export class ExtractionResultRepositoryImpl extends BaseRepository implements ExtractionResultRepository {
    create = async (args: { file: File }) => {
        const formData = new FormData()
        formData.set("file", args.file)
        await this._upload({path: "", body: formData})
    }

    getAll = async () => {
        const results = await this._get({path: ""})
        return results.map((p: Props) => new ExtractionResult(p))
    }

    getById = async (args: { id: string }) => {
        const result = await this._get({path: `/${args.id}`})
        return new ExtractionResult(result)
    }

    delete = async (args: { id: string }) => {
        await this._delete({path: `/${args.id}`})
    }
}
