import React, { FC, useEffect, useMemo, useState, useRef } from "react"
import { Dialog, Pane, toaster } from "evergreen-ui"
import { Stage, Layer, Text } from "react-konva"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import { readAsBlob, readAsText, isTextType } from "./files"
import Loader from "./Loader"
import Image from "./Image"
import JobFile from "../models/JobFile"
import Konva from "konva"

export interface Props {
  jobFile: JobFile | undefined
  onClose: () => void
}

const FileViewer: FC<Props> = ({ jobFile, onClose }) => {
  const [loaded, setLoaded] = useState<boolean>(false)
  const [fileUrl, setFileUrl] = useState<string>('')
  const [textData, setTextData] = useState<string>("")
  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  const stageRef = useRef<Konva.Stage>(null)

  useEffect(() => {
    if (loaded || !jobFile) {
      return
    }
    const loadFile = async () => {
      const signedUrl = await useCase.getSignedUrl(jobFile.fileKey)
      if (signedUrl) {
        setFileUrl(signedUrl.url)
        try {
          if (isTextType(jobFile.contentType)) {
            const fileData = await readAsBlob(signedUrl.url)
            setTextData(await readAsText(fileData))
          }
        } catch (e) {
          console.error(e)
          toaster.danger("Failed to load file")
          onClose()
        }
      } else {
        toaster.danger("Failed to load file")
        onClose()
      }
      setLoaded(true)
    }
    loadFile()
  }, [jobFile, loaded, onClose, useCase])

  useEffect(() => {
    console.info(stageRef.current)
  }, [loaded])

  const showUnsupportedFileAlert = () => {
    toaster.warning("Unsupported file...")
    onClose()
  }

  const renderFile = () => {
    if (!stageRef.current || !jobFile) {
      return null
    }
    if (textData.length) {
      if (jobFile.isJSON) {
        const text = JSON.stringify(JSON.parse(textData), undefined, 2)
        return <Text text={text} />
      }
    }
    if (jobFile.isImage) {
      return <Image url={fileUrl} outerHeight={1} outerWidth={1}/>
    }
    showUnsupportedFileAlert()
    return null
  }

  if (!jobFile) {
    return null
  }

  if (!loaded || !!stageRef.current) {
    return <Loader isShown={!loaded} />
  }

  return (
    <Pane height={"100%"}>
      <Stage width={window.innerWidth} height={window.innerHeight} ref={stageRef}>
        <Layer>
          {renderFile()}
        </Layer>
      </Stage>
    </Pane>
  )
}

export default FileViewer
