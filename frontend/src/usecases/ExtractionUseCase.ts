import {ExtractionResultRepository, ExtractionResultRepositoryImpl} from "../repositories/ExtractionResultRepository";
import ExtractionResult from "../models/ExtractionResult";
import {ENDPOINT_URL} from '../configs/config'

export interface ExtractionUseCase {
    startExtraction(file: File): Promise<boolean>
    getExtractionResults(): Promise<ExtractionResult[]|undefined>
    getExtractionResult(id: string): Promise<ExtractionResult|undefined>
    deleteExtractionResult(id: string): Promise<boolean>
}

export default class ExtractionUseCaseImpl implements ExtractionUseCase{
    private extractionRepository: ExtractionResultRepository

    constructor() {
        this.extractionRepository = new ExtractionResultRepositoryImpl(`${ENDPOINT_URL}/detect_texts`)
    }

    getExtractionResult = async (id: string) => {
        try {
            return await this.extractionRepository.getById({id})
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
            const result = await this.extractionRepository.create({file})
            return result.status
        } catch (e) {
            console.error(e)
        }
        return false
    }

    deleteExtractionResult = async (id: string) => {
        try {
            const result = await this.extractionRepository.delete({id})
            return result.status
        } catch (e) {
            console.error(e)
        }
        return false
    }
}
