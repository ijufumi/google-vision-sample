import React, { FC, useCallback, useEffect, useMemo, useState } from "react"
import { useNavigate } from "react-router"
import {
  Badge,
  Button,
  Dialog,
  EyeOpenIcon,
  Heading,
  IconButton,
  majorScale,
  Pane,
  Table,
  Text,
  toaster,
  TrashIcon,
  UploadIcon,
} from "evergreen-ui"
import Job, { JobStatus } from "../models/Job"
import JobUseCaseImpl from "../usecases/JobUseCase"
import FileUploadDialog from "../components/FileUploadDialog"
import LoaderOverlay from "../components/LoaderOverlay"

interface Props {}

const App: FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [showLoader, setShowLoader] = useState<boolean>(false)
  const [jobs, setJobs] = useState<Job[]>([])
  const [showFileUploadDialog, setShowFileUploadDialog] =
    useState<boolean>(false)
  const [deleteTargetId, setDeleteTargetId] = useState<string>("")

  const navigate = useNavigate()
  const useCase = useMemo(() => new JobUseCaseImpl(), [])

  const loadJobs = useCallback(async () => {
    const _jobs = await useCase.getJobs()
    if (_jobs) {
      setJobs(_jobs)
    } else {
      toaster.danger("Something wrong...")
    }
  }, [useCase])

  useEffect(() => {
    if (initialized) {
      return
    }
    const initialize = async () => {
      await loadJobs()
      setInitialized(true)
    }
    initialize()
  }, [initialized, loadJobs])

  useEffect(() => {
    const runningJob = jobs.find(j => j.status === JobStatus.Running)
    if (!runningJob) {
      return
    }
    const intervalId = window.setInterval(loadJobs, 1000)
    return function() {
      window.clearInterval(intervalId)
    }
  }, [jobs])

  const handleFileUpload = async (file: File) => {
    setShowFileUploadDialog(false)
    setShowLoader(true)
    const result = await useCase.startJob(file)
    setShowLoader(false)
    if (result) {
      toaster.success("Uploading succeeded")
    } else {
      toaster.danger("Uploading failed")
    }
    await loadJobs()
  }

  const confirmDelete = (id: string) => {
    setDeleteTargetId(id)
  }

  const handleDelete = async () => {
    if (!deleteTargetId) {
      return
    }
    setShowLoader(true)
    setDeleteTargetId("")
    const result = await useCase.deleteJob(deleteTargetId)
    setShowLoader(false)
    if (result) {
      setDeleteTargetId("")
      toaster.success("Deleting succeeded")
    } else {
      toaster.danger("Deleting failed")
    }
    await loadJobs()
  }

  const showResultPage = (id: string) => {
    navigate(`/${id}`)
  }

  const renderStatus = (status: JobStatus) => {
    let color: "red" | "blue" | "green" = "red"
    if (status === JobStatus.Running) {
      color = "blue"
    }
    if (status === JobStatus.Succeeded) {
      color = "green"
    }
    return (
      <Badge
        color={color}
        style={{
          borderRadius: "10px",
          height: "30px",
          width: "100px",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          margin: "5px",
        }}
      >
        {status}
      </Badge>
    )
  }

  const renderResults = () => {
    return (
      <Pane>
        <Table>
          <Table.Head>
            <Table.TextHeaderCell>ID</Table.TextHeaderCell>
            <Table.TextHeaderCell>Name</Table.TextHeaderCell>
            <Table.TextHeaderCell>Status</Table.TextHeaderCell>
            <Table.TextHeaderCell>Created At</Table.TextHeaderCell>
            <Table.TextHeaderCell>Updated At</Table.TextHeaderCell>
            <Table.TextHeaderCell>Operations</Table.TextHeaderCell>
          </Table.Head>
          <Table.Body>
            {jobs.map((result) => (
              <Table.Row key={result.id}>
                <Table.TextCell>{result.id}</Table.TextCell>
                <Table.TextCell>
                  { result.name }
                </Table.TextCell>
                <Table.TextCell>{renderStatus(result.status)}</Table.TextCell>
                <Table.TextCell>{result.readableCreatedAt}</Table.TextCell>
                <Table.TextCell>{result.readableUpdatedAt}</Table.TextCell>
                <Table.Cell>
                  <IconButton
                    icon={EyeOpenIcon}
                    marginRight={majorScale(2)}
                    onClick={() => showResultPage(result.id)}
                  />
                  <IconButton
                    icon={TrashIcon}
                    intent="danger"
                    onClick={() => confirmDelete(result.id)}
                  />
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Pane>
    )
  }

  return (
    <Pane
      display="flex"
      flexDirection="column"
      backgroundColor="#FFFFFF"
      margin="20px"
      position="relative"
      height="calc(100vh - 40px)"
    >
      <Pane padding="20px">
        <Pane
          display="flex"
          justifyContent="space-between"
          width="100%"
          paddingBottom="40px"
        >
          <Heading size={900}>Google Vision API Client</Heading>
          <Button
            appearance="primary"
            iconAfter={UploadIcon}
            onClick={() => setShowFileUploadDialog(true)}
          >
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
      {!!deleteTargetId && (
        <Dialog
          key={deleteTargetId}
          isShown={!!deleteTargetId}
          title="Are you sure deleting?"
          onCloseComplete={() => setDeleteTargetId("")}
          onConfirm={handleDelete}
          confirmLabel="Delete"
        >
          <Text size={600}>Would you like to delete it?</Text>
        </Dialog>
      )}
      <LoaderOverlay isShown={showLoader} />
    </Pane>
  )
}

export default App
