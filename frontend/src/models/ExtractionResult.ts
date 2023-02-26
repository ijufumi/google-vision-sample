import ExtractedText, {Props as ExtractedTextProps} from "./ExtractedText";
import { formatToDate } from "../components/dates"

export enum ExtractionResultStatus {
    Running = "running",
    Succeeded = "succeeded",
    Failed = "failed"
}

export interface Props {
    id: string
    status: ExtractionResultStatus
    imageKey: string
    outputKey: string
    createdAt: number
    updatedAt: number
    extractedTexts: ExtractedTextProps[]
}

export default class ExtractionResult {
    readonly id: string
    readonly status: ExtractionResultStatus
    readonly imageKey: string
    readonly outputKey: string
    readonly createdAt: number
    readonly updatedAt: number
    readonly extractedTexts: ExtractedText[]

    constructor(props: Props) {
        this.id = props.id
        this.status = props.status
        this.imageKey = props.imageKey
        this.outputKey = props.outputKey
        this.createdAt = props.createdAt
        this.updatedAt = props.updatedAt
        if (props.extractedTexts) {
            this.extractedTexts = props.extractedTexts.map((p => new ExtractedText(p)))
        } else {
            this.extractedTexts = []
        }
    }

    get readableCreatedAt(){
        return formatToDate(this.createdAt)
    }

    get readableUpdatedAt(){
        return formatToDate(this.updatedAt)
    }
}
