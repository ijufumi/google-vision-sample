import React, { FC, useMemo, useEffect, useState } from 'react'
import { Dialog } from "evergreen-ui"
import ExtractionResult from "../models/ExtractionResult"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"

export interface Props {
  extractionResultId: string
  onClose: () => {}
}

const ResultViewerDialog: FC<Props> = ({ extractionResultId, onClose }) => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [extractionResult, setExtractionResult] = useState<ExtractionResult|undefined>(undefined)
  const isShown = useMemo(() => !!extractionResultId, [extractionResultId])

  const useCase = new ExtractionUseCaseImpl()

  useEffect(() => {
    initialize()
  })

  const initialize = async () => {
    const _extractionResult = await useCase.getExtractionResult(extractionResultId)
    if (_extractionResult) {
      setExtractionResult(_extractionResult)
    }
    setInitialized(true)
  }

  if (!initialized) {
    return null
  }

  return (
    <Dialog
      isShown={isShown}
      onCloseComplete={onClose}
      hasCancel={false}
      confirmLabel="Close"
    >

    </Dialog>
  )
}

export default ResultViewerDialog
