import React, { FC, useCallback, useEffect, useMemo, useRef, useState } from "react"
import { useParams } from "react-router"
import { useNavigate } from "react-router"
import { Stage, Layer } from "react-konva"
import Konva from "konva"
import { Pane, Table, Button } from "evergreen-ui"
import Job from "../models/Job"
import JobUseCaseImpl from "../usecases/JobUseCase"
import Image from "../components/Image"
import Loader from "../components/Loader"

export interface Props {}

const ResultPage : FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [job, setJob] = useState<Job | undefined>(undefined)
  const [inputFileUrl, setInputFileUrl] = useState<string>("")
  const [stageWidth, setStageWidth] = useState<number>(1)
  const [stageHeight, setStageHeight] = useState<number>(1)
  const [imageLoaded, setImageLoaded] = useState<boolean>(false)

  const stageRef = useRef<Konva.Stage>(null)
  const navigate = useNavigate()
  const { jobId } = useParams()

  useEffect(() => {
    if (stageRef.current) {
      setStageWidth(stageRef.current.width())
      setStageHeight(stageRef.current.height())
    }
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

  const handleImageLoaded = () => {
    setImageLoaded(true)
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
      <Pane>
        <Button onClick={handleBackToTop}>
          Return to top
        </Button>
      </Pane>
      <Pane display="flex" width="100%" height="100%">
        <Pane width="55%" marginRight={"5px"} height="100%" overflow="scroll">
          { !imageLoaded && <Loader /> }
          <Stage ref={stageRef} width={window.innerWidth/2 - 10} height={window.innerHeight - 100}>
            <Layer>
              <Image
                outerWidth={stageWidth}
                outerHeight={stageHeight}
                url={inputFileUrl}
                onLoaded={handleImageLoaded}
              />
            </Layer>
          </Stage>
        </Pane>
        <Pane width="40%" height="100%">
          <Table height="100%">
            <Table.Head>
              <Table.TextHeaderCell>Texts</Table.TextHeaderCell>
            </Table.Head>
            <Table.Body overflow="scroll" height="100%">
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
    </Pane>
  )
}

export default ResultPage
