import {ExtractionResultRepository, ExtractionResultRepositoryImpl} from "../repositories/ExtractionResultRepository";
import ExtractionResult from "../models/ExtractionResult";

export interface ExtractionUseCase {
    startExtraction(file: File): Promise<boolean>
    getExtractionResults(): Promise<ExtractionResult[]|undefined>
    getExtractionResult(id: string): Promise<ExtractionResult|undefined>
}

export default class ExtractionUseCaseImpl implements ExtractionUseCase{
    private extractionRepository: ExtractionResultRepository

    constructor() {
        this.extractionRepository = new ExtractionResultRepositoryImpl("/detect_texts")
    }

    getExtractionResult = async (id: string) => {
        try {
            return  await this.extractionRepository.getById({id})
        } catch (e) {
            console.error(e)
        }
        return undefined
    }

    getExtractionResults = async () => {
        try {
            return await this.extractionRepository.getAll()
        } catch (e) {
            console.error(e)
        }
        return undefined
    }

    startExtraction = async (file: File) => {
        try {
            await this.extractionRepository.create({file})
            return true
        } catch (e) {
            console.error(e)
        }
        return false
    }

}
