import React, { FC, useCallback, useEffect, useMemo, useRef, useState } from "react"
import { useParams } from "react-router"
import { useNavigate } from "react-router"
import { Stage, Layer } from "react-konva"
import Konva from "konva"
import { Pane, Table, Button, Text, IconButton, CaretLeftIcon, CaretRightIcon } from "evergreen-ui"
import Job from "../models/Job"
import Rect from "../components/Rect"
import JobUseCaseImpl from "../usecases/JobUseCase"
import Image from "../components/Image"
import Loader from "../components/Loader"
import ExtractedText from "../models/ExtractedText"

export interface Props {}

const ResultPage : FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [job, setJob] = useState<Job | undefined>(undefined)
  const [resultIndex, setResultIndex] = useState<number>(-1)
  const [maxResultsNumber, setMaxResultNumber] = useState<number>(0)
  const [inputFileUrl, setInputFileUrl] = useState<string>("")
  const [extractedTexts, setExtractedTexts] = useState<ExtractedText[]>([])
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
      setMaxResultNumber(_job.inputFiles.length)
      setResultIndex(0)
    }
    setInitialized(true)
  }, [jobId, useCase])

  useEffect(() => {
    if (initialized) {
      return
    }
    initialize()
  }, [])

  useEffect(() => {
    setSelectedTextId('')
    if (!job || job.inputFiles.length < resultIndex) {
      setInputFileUrl('')
      setExtractedTexts([])
      return
    }
    const inputFile = job.inputFiles[resultIndex]
    const setResults = async () => {
      const signedUrl = await useCase.getSignedUrl(inputFile.fileKey)
      if (signedUrl) {
        setInputFileUrl(signedUrl.url)
        setExtractedTexts(inputFile.outputFiles[0].extractedTexts)
      } else {
        setInputFileUrl('')
        setExtractedTexts([])
      }
    }
    setResults()
  }, [resultIndex])

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
      <Pane marginBottom="10px">
        <Pane display="flex" justifyContent="space-between" width="59%">
          <Button onClick={handleBackToTop}>
            Return to top
          </Button>
          {maxResultsNumber > 1 &&
            <Pane display="flex" alignItems="center">
              <Text>{resultIndex+1}</Text>
              <Text marginLeft="2px" marginRight="2px">{"/"}</Text>
              <Text>{maxResultsNumber}</Text>
              <Pane display="flex" marginLeft="10px">
                <IconButton icon={CaretLeftIcon} disabled={resultIndex === 0} onClick={() => setResultIndex(resultIndex-1)} />
                <IconButton icon={CaretRightIcon} disabled={resultIndex === maxResultsNumber-1} onClick={() => setResultIndex(resultIndex+1)} />
              </Pane>
            </Pane>
          }
        </Pane>
      </Pane>
      <Pane display="flex" width="100%" height="95%">
        <Pane width="59%" marginRight={"5px"} overflow="scroll">
          { !imageLoaded && <Loader /> }
          <Stage ref={stageRef} width={window.innerWidth/2 - 10} height={window.innerHeight - 90} style={{ border: "1px solid #E6E8F0" }}>
            <Layer>
              <Image
                outerWidth={stageWidth}
                outerHeight={stageHeight}
                url={inputFileUrl}
                onLoaded={handleImageLoaded}
              />
            </Layer>
            <Layer>
              {extractedTexts.map(result => {
                return <Rect
                  key={result.id}
                  x={result.left * scale}
                  y={result.top * scale}
                  width={(result.right - result.left) * scale}
                  height={(result.bottom - result.top) * scale}
                  visible={result.id === selectedTextId}
                />
              })}
            </Layer>
          </Stage>
        </Pane>
        <Pane width="40%" height="100%">
          <Table height="calc(100% - 5px)">
            <Table.Head>
              <Table.TextHeaderCell>Texts</Table.TextHeaderCell>
            </Table.Head>
            <Table.Body overflow="scroll" height="calc(100% - 50px)">
              {extractedTexts.map((result) => {
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
