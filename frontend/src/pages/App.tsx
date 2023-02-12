import React, { FC, Fragment, useState, useEffect } from 'react'
import ExtractionResult from "../models/ExtractionResult";
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase";

interface Props {}

const App: FC<Props> = () => {
    const [initialized, setInitialized] = useState<boolean>(false)
    const [extractionResults, setExtractionResults] = useState<ExtractionResult[]>([])

    const useCase = new ExtractionUseCaseImpl()

    useEffect(() => {
        const initialize = async () => {
            const _extractionResults = await useCase.getExtractionResults()
            if (_extractionResults) {
                setExtractionResults(_extractionResults)
            } else {
                console.error("something wrong...")
            }
            setInitialized(true)
        }
        initialize()
    }, [])

    if (!initialized) {
        return null
    }

    return <Fragment>

    </Fragment>
}

export default App;
