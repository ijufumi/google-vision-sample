import React, { FC, useEffect, useMemo, useState } from "react"
import { Dialog, Pane, toaster } from "evergreen-ui"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import { readAsBlob, readAsText, isTextType } from "./files"
import Loader from "./Loader"
import JobFile from "../models/JobFile"

export interface Props {
  jobFile: JobFile
  isShown: boolean
  onClose: () => void
}

const FileViewer: FC<Props> = ({ jobFile, isShown, onClose }) => {
  const [loaded, setLoaded] = useState<boolean>(false)
  const [blobData, setBlobData] = useState<Blob | undefined>(undefined)
  const [textData, setTextData] = useState<string>("")
  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  useEffect(() => {
    if (loaded) {
      return
    }
    const loadFile = async () => {
      const signedUrl = await useCase.getSignedUrl(jobFile.fileKey)
      if (signedUrl) {
        try {
          const fileData = await readAsBlob(signedUrl.url)
          if (isTextType(jobFile.contentType)) {
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
  }, [jobFile, loaded, onClose, useCase])

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
      if (jobFile.isJSON) {
        return React.createElement(
          "div",
          {
            height: "100%",
            width: "100%",
            style: {
              border: "1px solid #000000",
              borderRadius: "10px",
              whiteSpace: "break-spaces",
            },
          },
          JSON.stringify(JSON.parse(textData), undefined, 2)
        )
      }
    }
    if (blobData) {
      if (jobFile.isImage) {
        return React.createElement("img", {
          src: URL.createObjectURL(blobData),
          height: "100%",
          width: "100%",
          style: {
            objectFit: "contain",
          },
        })
      }
    }
    showUnsupportedFileAlert()
    return <div />
  }

  if (isShown && !loaded) {
    return <Loader isShown={!loaded} />
  }

  return (
    <Pane>
      <Dialog
        isShown={isShown}
        onConfirm={onClose}
        hasCancel={false}
        confirmLabel={"OK"}
        shouldCloseOnOverlayClick={false}
        shouldCloseOnEscapePress={false}
        width={"1000px"}
      >
        <Pane height="700px" width="100%">
          {renderFile()}
        </Pane>
      </Dialog>
    </Pane>
  )
}

export default FileViewer
