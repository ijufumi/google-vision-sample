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
    const initialize = async () => {
      const _extractionResults = await useCase.getExtractionResults()
      if (_extractionResults) {
        if (_extractionResults.length < 20) {
          Array.from(Array(20 - _extractionResults.length).keys()).forEach((v) => {
            _extractionResults.push({
              id: String(v),
            } as ExtractionResult)
          })
        }
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

  const handleFileUpload = async (file: File) => {
    const result = await useCase.startExtraction(file)
    if (result) {
      toaster.success("Uploading succeeded")
      setShowFileUploadDialog(false)
    } else {
      toaster.danger("Uploading failed")
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
          <Table.VirtualBody>
            {renderData()}
          </Table.VirtualBody>
        </Table>
      </Pane>
    )
  }

  const renderData = () => {
    if (extractionResults.length === 0) {
      return (
        <Table.Row>
          <Table.Cell>
            <Heading size={700}>There was no results</Heading>
          </Table.Cell>
        </Table.Row>
      )
    }

    return <React.Fragment>
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
            <IconButton icon={TrashIcon} intent="danger" />
          </Table.Cell>
        </Table.Row>
      ))}
    </React.Fragment>
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
