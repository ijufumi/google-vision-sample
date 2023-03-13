import React, { FC, useCallback, useEffect, useMemo, useRef, useState } from "react"
import { useParams } from "react-router"
import { useNavigate } from "react-router"
import { Stage, Layer } from "react-konva"
import Konva from "konva"
import { Pane, Table, Button } from "evergreen-ui"
import Job from "../models/Job"
import JobUseCaseImpl from "../usecases/JobUseCase"
import Image from "../components/Image"

export interface Props {}

const ResultPage : FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [job, setJob] = useState<Job | undefined>(undefined)
  const [inputFileUrl, setInputFileUrl] = useState<string>("")

  const stageRef = useRef<Konva.Stage>(null)
  const navigate = useNavigate()
  const { jobId } = useParams()

  const stageWidth = useMemo(() => {
    if (!stageRef.current) {
      return 1
    }
    return stageRef.current.width()
  }, [stageRef])
  const stageHeight = useMemo(() => {
    if (!stageRef.current) {
      return 1
    }
    return stageRef.current.height()
  }, [stageRef])

  const useCase = useMemo(() => new JobUseCaseImpl(), [])

  const initialize = useCallback(async () => {
    if (jobId === null || jobId === undefined || jobId === "") {
      return
    }

    const _job = await useCase.getJob(jobId)
    if (_job) {
      setJob(_job)
      if (_job.inputFile) {
        const signedUrl = await useCase.getSignedUrl(_job.inputFile.fileKey)
        if (signedUrl) {
          setInputFileUrl(signedUrl.url)
        }
      }
    }
    setInitialized(true)
  }, [jobId, useCase])

  useEffect(() => {
    if (initialized) {
      return
    }
    initialize()
  }, [])

  const handleBackToTop = () => {
    navigate("/")
  }

  if (!initialized) {
    return null
  }

  return (
    <Pane width="100%" height="100%" display="flex">
      <Pane>
        <Button onClick={handleBackToTop}>
          Return to top
        </Button>
      </Pane>
      <Pane width="47%" marginRight={"5px"}>
        <Stage ref={stageRef}>
          <Layer>
            <Image outerWidth={stageWidth} outerHeight={stageHeight} url={inputFileUrl} />
          </Layer>
        </Stage>
      </Pane>
      <Pane width="47%" height="100%" overflow="scroll">
        <Table>
          <Table.Head>
            <Table.TextHeaderCell>Texts</Table.TextHeaderCell>
          </Table.Head>
          <Table.Body>
            {job?.extractedTexts.map((result) => {
              return (
                <Table.Row key={result.id}>
                  <Table.TextCell>{result.text}</Table.TextCell>
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table>
      </Pane>
    </Pane>
  )
}

export default ResultPage
