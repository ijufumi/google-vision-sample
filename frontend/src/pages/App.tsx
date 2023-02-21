import React, { FC, useState, useEffect } from "react"
import { Pane, Button, UploadIcon, TrashIcon, EyeOpenIcon, Text, Heading, Table, IconButton, majorScale, toaster } from "evergreen-ui"
import ExtractionResult from "../models/ExtractionResult"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import FileUploadDialog from "../components/FileUploadDialog"

interface Props {
}

const App: FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [extractionResults, setExtractionResults] = useState<ExtractionResult[]>([])
  const [showFileUploadDialog, setShowFileUploadDialog] = useState<boolean>(false)

  const useCase = new ExtractionUseCaseImpl()

  useEffect(() => {
    if (initialized) {
      return
    }
    const initialize = async () => {
      const _extractionResults = await useCase.getExtractionResults()
      console.info("initialize...")
      if (_extractionResults) {
        if (_extractionResults.length < 20) {
          const extra = Array.from(Array(20 - _extractionResults.length).keys()).map((v) => {
            return {
              id: String(v),
            } as ExtractionResult
          })
          _extractionResults.push(...extra)
        }
        setExtractionResults(_extractionResults)
      } else {
        toaster.danger("Something wrong...")
      }
      setInitialized(true)
    }
    initialize()
  }, [])

  if (!initialized) {
    return null
  }

  const handleFileUpload = async (file: File) => {
    const result = await useCase.startExtraction(file)
    if (result) {
      toaster.success("Uploading succeeded")
      setShowFileUploadDialog(false)
    } else {
      toaster.danger("Uploading failed")
    }
  }

  const handleDelete = async (id: string) => {
    // TODO: Add confirmation
    const result = await useCase.deleteExtractionResult(id)
    if (result) {
      toaster.success("Deleting succeeded")
    } else {
      toaster.danger("Deleting failed")
    }
  }

  const renderResults = () => {
    return (
      <Pane>
        <Table>
          <Table.Head>
            <Table.TextHeaderCell>ID</Table.TextHeaderCell>
            <Table.TextHeaderCell>Status</Table.TextHeaderCell>
            <Table.TextHeaderCell>Input</Table.TextHeaderCell>
            <Table.TextHeaderCell>Output</Table.TextHeaderCell>
            <Table.TextHeaderCell>CreatedAt</Table.TextHeaderCell>
            <Table.TextHeaderCell>UpdatedAt</Table.TextHeaderCell>
            <Table.TextHeaderCell>Operations</Table.TextHeaderCell>
          </Table.Head>
          <Table.Body>
            {extractionResults.map(result => (
              <Table.Row key={result.id}>
                <Table.TextCell>{result.id}</Table.TextCell>
                <Table.TextCell>{result.status}</Table.TextCell>
                <Table.TextCell>{result.imageUri}</Table.TextCell>
                <Table.TextCell>{result.outputUri}</Table.TextCell>
                <Table.TextCell>{result.createdAt}</Table.TextCell>
                <Table.TextCell>{result.updatedAt}</Table.TextCell>
                <Table.Cell>
                  <IconButton icon={EyeOpenIcon} marginRight={majorScale(2)} />
                  <IconButton icon={TrashIcon} intent="danger" onClick={() => handleDelete(result.id)}/>
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Pane>
    )
  }

  return <Pane display="flex" flexDirection="column" backgroundColor="#FFFFFF" margin="20px" height="calc(100vh - 40px)">
    <Pane padding="20px">
      <Pane display="flex" justifyContent="space-between" width="100%" paddingBottom="40px">
        <Heading size={900}>
          Google Vision API Client
        </Heading>
        <Button appearance="primary" iconAfter={UploadIcon} onClick={() => setShowFileUploadDialog(true)}>
          Upload
        </Button>
      </Pane>
      {renderResults()}
    </Pane>
    <FileUploadDialog
      isShown={showFileUploadDialog}
      onClose={() => setShowFileUploadDialog(false)}
      onUpload={handleFileUpload}
    />
  </Pane>
}

export default App
