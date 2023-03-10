import React, { FC, useMemo, useEffect, useState, useCallback } from "react"
import { Pane, Dialog, Table } from "evergreen-ui"
import Job from "../models/Job"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import { readAsBlob } from "./files"

export interface Props {
  extractionResultId: string
  onClose: () => void
}

const ResultViewerDialog: FC<Props> = ({ extractionResultId, onClose }) => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [extractionResult, setExtractionResult] = useState<
    Job | undefined
  >(undefined)
  const [blobData, setBlobData] = useState<Blob | undefined>(undefined)

  const isShown = useMemo(() => !!extractionResultId, [extractionResultId])

  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  const initialize = useCallback(async () => {
    const _extractionResult = await useCase.getExtractionResult(
      extractionResultId
    )
    if (_extractionResult) {
      setExtractionResult(_extractionResult)
      if (_extractionResult.outputFile) {
        const signedUrl = await useCase.getSignedUrl(_extractionResult.outputFile.fileKey)
        if (signedUrl) {
          const fileData = await readAsBlob(signedUrl.url)
          setBlobData(fileData)
        }
      }
    }
    setInitialized(true)
  }, [extractionResultId, useCase])

  useEffect(() => {
    if (initialized) {
      return
    }
    initialize()
  }, [])

  if (!initialized) {
    return null
  }

  const renderImageView = () => {
    if (blobData) {
      return React.createElement("img", {
        src: URL.createObjectURL(blobData),
        height: "100%",
        width: "100%",
        alt: "image",
        style: {
          objectFit: "contain",
        },
      })
    }
    return <div />
  }

  return (
    <Pane>
      <Dialog
        isShown={isShown}
        onCloseComplete={onClose}
        hasCancel={false}
        confirmLabel="Close"
        shouldCloseOnEscapePress={false}
        shouldCloseOnOverlayClick={false}
        width={"1000px"}
      >
        <Pane width="100%" height="700px" display="flex">
          <Pane width="47%" marginRight={"5px"}>
            {renderImageView()}
          </Pane>
          <Pane width="47%" height="100%" overflow="scroll">
            <Table>
              <Table.Head>
                <Table.TextHeaderCell>Texts</Table.TextHeaderCell>
              </Table.Head>
              <Table.Body>
                {extractionResult?.extractedTexts.map((result) => {
                  return (
                    <Table.Row>
                      <Table.TextCell>{result.text}</Table.TextCell>
                    </Table.Row>
                  )
                })}
              </Table.Body>
            </Table>
          </Pane>
        </Pane>
      </Dialog>
    </Pane>
  )
}

export default ResultViewerDialog
