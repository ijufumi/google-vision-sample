import React, { FC, useState, useEffect, useMemo, useCallback } from "react"
import { Pane, Button, Dialog, Text, Heading, Table, IconButton, Badge, majorScale, toaster, UploadIcon, TrashIcon, EyeOpenIcon, DocumentOpenIcon } from "evergreen-ui"
import ExtractionResult, { ExtractionResultStatus } from "../models/ExtractionResult"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import FileUploadDialog from "../components/FileUploadDialog"
import Loader from "../components/Loader"
import FileViewer from "../components/FileViewer"

interface Props {
}

const App: FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [showLoader, setShowLoader] = useState<boolean>(false)
  const [extractionResults, setExtractionResults] = useState<ExtractionResult[]>([])
  const [showFileUploadDialog, setShowFileUploadDialog] = useState<boolean>(false)
  const [deleteTargetId, setDeleteTargetId] = useState<string>('')
  const [fileKey, setFileKey] = useState<string>('')

  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  const loadExtractionResults = useCallback(async () => {
    const _extractionResults = await useCase.getExtractionResults()
    console.info("initialize...")
    if (_extractionResults) {
      setExtractionResults(_extractionResults)
    } else {
      toaster.danger("Something wrong...")
    }
  }, [useCase])

  useEffect(() => {
    if (initialized) {
      return
    }
    const initialize = async () => {
      console.info("initialize2...")
      await loadExtractionResults()
      setInitialized(true)
    }
    initialize()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

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

  const renderStatus = (status: ExtractionResultStatus) => {
    let color: "red" | "blue" | "green" = "red"
    if (status === ExtractionResultStatus.Running) {
      color = "blue"
    }
    if (status === ExtractionResultStatus.Succeeded) {
      color = "green"
    }
    return <Badge color={color} style={{
      borderRadius: "10px",
      height: "30px",
      width: "100px",
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
      margin: "5px",
    }}>
      {status}
    </Badge>
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
                <Table.TextCell>{renderStatus(result.status)}</Table.TextCell>
                <Table.TextCell>
                  {result.imageKey && <IconButton icon={DocumentOpenIcon} appearance="minimal" onClick={() => setFileKey(result.imageKey)} />}
                </Table.TextCell>
                <Table.TextCell>
                  {result.outputKey && <IconButton icon={DocumentOpenIcon} appearance="minimal" onClick={() => setFileKey(result.outputKey)} />}
                </Table.TextCell>
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
    <FileViewer fileKey={fileKey} isShown={!!fileKey} onClose={() => setFileKey('')} />
  </Pane>
}

export default App
