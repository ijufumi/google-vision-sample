import React, { FC, useState, useEffect, useMemo } from "react"
import { Pane, Button, UploadIcon, TrashIcon, EyeOpenIcon, Dialog, Text, Heading, Table, IconButton, majorScale, toaster } from "evergreen-ui"
import ExtractionResult from "../models/ExtractionResult"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import FileUploadDialog from "../components/FileUploadDialog"
import Loader from "../components/Loader"

interface Props {
}

const App: FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [showLoader, setShowLoader] = useState<boolean>(false)
  const [extractionResults, setExtractionResults] = useState<ExtractionResult[]>([])
  const [showFileUploadDialog, setShowFileUploadDialog] = useState<boolean>(false)
  const [deleteTargetId, setDeleteTargetId] = useState<string>('')

  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  const loadExtractionResults = async () => {
    const _extractionResults = await useCase.getExtractionResults()
    console.info("initialize...")
    if (_extractionResults) {
      setExtractionResults(_extractionResults)
    } else {
      toaster.danger("Something wrong...")
    }
  }

  useEffect(() => {
    if (initialized) {
      return
    }
    const initialize = async () => {
      await loadExtractionResults()
      setInitialized(true)
    }
    initialize()
  }, [initialized, useCase])

  if (!initialized) {
    return null
  }

  const handleFileUpload = async (file: File) => {
    setShowFileUploadDialog(false)
    setShowLoader(true)
    const result = await useCase.startExtraction(file)
    setShowLoader(false)
    if (result) {
      toaster.success("Uploading succeeded")
    } else {
      toaster.danger("Uploading failed")
    }
    await loadExtractionResults()
  }

  const confirmDelete = (id: string) => {
    setDeleteTargetId(id)
  }

  const handleDelete = async () => {
    if (!deleteTargetId) {
      return
    }
    setShowLoader(true)
    setDeleteTargetId('')
    const result = await useCase.deleteExtractionResult(deleteTargetId)
    setShowLoader(false)
    if (result) {
      setDeleteTargetId('')
      toaster.success("Deleting succeeded")
    } else {
      toaster.danger("Deleting failed")
    }
    await loadExtractionResults()
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
                <Table.TextCell>{result.readableCreatedAt}</Table.TextCell>
                <Table.TextCell>{result.readableUpdatedAt}</Table.TextCell>
                <Table.Cell>
                  <IconButton icon={EyeOpenIcon} marginRight={majorScale(2)} />
                  <IconButton icon={TrashIcon} intent="danger" onClick={() => confirmDelete(result.id)}/>
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
    <Dialog
      isShown={!!deleteTargetId}
      title="Are you sure deleting?"
      onCloseComplete={() => setDeleteTargetId('')}
      onConfirm={handleDelete}
      confirmLabel="Delete"
    >
      <Text size={600}>Would you like to delete it?</Text>
    </Dialog>
    <Loader isShown={showLoader} />
  </Pane>
}

export default App
