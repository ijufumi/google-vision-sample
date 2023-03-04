import React, { FC, useEffect, useMemo, useState } from "react"
import { Dialog, Pane, toaster } from "evergreen-ui"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import { readAsBlob, readAsText, isTextType } from "./files"
import Loader from "./Loader"


export interface Props {
  fileKey: string
  isShown: boolean
  onClose: () => void
}

const FileViewer: FC<Props> = ({ fileKey, isShown, onClose }) => {
  const [loaded, setLoaded] = useState<boolean>(false)
  const [blobData, setBlobData] = useState<Blob|undefined>(undefined)
  const [textData, setTextData] = useState<string>('')
  const [contentType, setContentType] = useState<string>('')
  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  useEffect(() => {
    if (loaded) {
      return
    }
    const loadFile = async () => {
      if (!fileKey) {
        return
      }
      const signedUrl = await useCase.getSignedUrl(fileKey)
      if (signedUrl) {
        try {
          const fileData = await readAsBlob(signedUrl.url)
          setContentType(fileData.type)
          if (isTextType(fileData.type)) {
            setTextData(await readAsText(fileData))
          } else {
            setBlobData(fileData)
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
  }, [fileKey, loaded, onClose, useCase])

  const showUnsupportedFileAlert = () => {
    toaster.warning("Unsupported file...")
    onClose()
  }

  const renderFile = () => {
    if (!blobData && !textData) {
      showUnsupportedFileAlert()
      return <div />
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
    showUnsupportedFileAlert()
    return <div />
  }

  if (isShown && !loaded) {
    return <Loader isShown={!loaded} />
  }

  return <Pane>
    <Dialog isShown={isShown} onConfirm={onClose} hasCancel={false} confirmLabel={"OK"}>
      {renderFile()}
    </Dialog>
  </Pane>
}

export default FileViewer
