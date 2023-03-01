import React, { FC, useEffect, useMemo, useState } from "react"
import { Dialog, Pane } from "evergreen-ui"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import { readAsBlob, readAsText, isTextType } from "./files"
import Loader from "./Loader"


export interface Props {
  key: string
  isShown: boolean
  onClose: () => void
}

const FileViewer: FC<Props> = ({ key, isShown, onClose }) => {
  const [loaded, setLoaded] = useState<boolean>(false)
  const [blobData, setBlobData] = useState<Blob|undefined>(undefined)
  const [textData, setTextData] = useState<string>('')
  const [contentType, setContentType] = useState<string>('')
  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  useEffect(() => {
    const loadFile = async () => {
      const signedUrl = await useCase.getSignedUrl(key)
      if (signedUrl) {
        const fileData = await readAsBlob(signedUrl.url)
        setContentType(fileData.type)
        if (isTextType(fileData.type)) {
          setTextData(await readAsText(fileData))
        } else {
          setBlobData(fileData)
        }
      }
      setLoaded(true)
    }
    loadFile()
  }, [key, useCase])

  const renderFile = () => {
    if (!blobData && !textData) {
      return null
    }
    if (textData.length) {
      if (contentType === "application/json") {
        return React.createElement("pre", JSON.stringify(JSON.parse(textData)))
      }
    }
    if (blobData) {
      if (contentType.startsWith("image/")) {
        return  React.createElement("img", {
          src: URL.createObjectURL(blobData)
        })
      }
    }
    return null
  }

  if (!loaded) {
    return <Loader isShown={!loaded} />
  }

  return <Pane>
    <Dialog isShown={isShown} onConfirm={onClose} hasCancel={false}>
      {renderFile()}
    </Dialog>
  </Pane>
}

export default FileViewer
