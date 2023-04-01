import React, { FC, useCallback, useEffect, useMemo, useRef, useState } from "react"
import { useParams } from "react-router"
import { useNavigate } from "react-router"
import { Stage, Layer, Rect } from "react-konva"
import Konva from "konva"
import { Pane, Table, Button } from "evergreen-ui"
import Job from "../models/Job"
import JobUseCaseImpl from "../usecases/JobUseCase"
import Image from "../components/Image"
import Loader from "../components/Loader"
import ExtractedText from "../models/ExtractedText"

export interface Props {}

const ResultPage : FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [job, setJob] = useState<Job | undefined>(undefined)
  const [inputFileUrl, setInputFileUrl] = useState<string>("")
  const [stageWidth, setStageWidth] = useState<number>(1)
  const [stageHeight, setStageHeight] = useState<number>(1)
  const [imageLoaded, setImageLoaded] = useState<boolean>(false)
  const [selectedTextId, setSelectedTextId] = useState<string>('')
  const [scale, setScale] = useState<number>(1)

  const stageRef = useRef<Konva.Stage>(null)
  const navigate = useNavigate()
  const { jobId } = useParams()

  const basicStyle = {
    cursor: "pointer",
  }

  const selectedStyle = {
    border: "2px solid #52BD95",
    cursor: "pointer",
  }

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

  const handleSelectText = (text: ExtractedText) => {
    if (selectedTextId === text.id) {
      setSelectedTextId('')
    } else {
      setSelectedTextId(text.id)
    }
  }

  const handleImageLoaded = (scale: number) => {
    setImageLoaded(true)
    setScale(scale)
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
      <Pane display="flex" width="100%" height="99%">
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
            <Layer>
              {job?.outputFile?.extractedTexts.map(result => {
                return <Rect
                  key={result.id}
                  x={result.left * scale}
                  y={result.top * scale}
                  width={(result.right - result.left) * scale}
                  height={(result.bottom - result.top) * scale}
                  stroke={"#D14343"}
                  fill={"#F9DADA"}
                  opacity={0.3}
                  strokeWidth={2}
                  visible={result.id === selectedTextId}
                />
              })}
            </Layer>
          </Stage>
        </Pane>
        <Pane width="40%" height="calc(100% - 30px)">
          <Table height="calc(100% - 30px)">
            <Table.Head>
              <Table.TextHeaderCell>Texts</Table.TextHeaderCell>
            </Table.Head>
            <Table.Body overflow="scroll" height="calc(100% - 40px)">
              {job?.outputFile?.extractedTexts.map((result) => {
                return (
                  <Table.Row key={result.id} style={selectedTextId === result.id ? selectedStyle : basicStyle}>
                    <Table.TextCell onClick={() => handleSelectText(result)}>{result.text}</Table.TextCell>
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
