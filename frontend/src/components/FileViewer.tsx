import React, { FC, useMemo, useEffect, useState } from "react"
import { Pane, Dialog } from "evergreen-ui"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"
import { readAsFile } from "./files"
import Loader from "./Loader"


export interface Props {
  key: string
  isShown: boolean
}

const FileViewer: FC<Props> = ({ key, isShown }) => {
  const [loaded, setLoaded] = useState<boolean>(false)
  const [file, setFile] = useState<File|undefined>(undefined)
  const useCase = useMemo(() => new ExtractionUseCaseImpl(), [])

  useEffect(() => {
    const loadFile = async () => {
      const signedUrl = await useCase.getSignedUrl(key)
      if (signedUrl) {
        const _file = await readAsFile(signedUrl.url)
        setFile(_file)
      }
      setLoaded(true)
    }
    loadFile()
  }, [key, useCase])

  const renderFile = () => {
    if (!file) {
      return null
    }
    if (file.type.startsWith("image/")) {
      const image = React.createElement("img", {
        src: URL.createObjectURL(file)
      })
      return image
    }
  }

  if (!loaded) {
    return <Loader isShown={!loaded} />
  }

  return <Pane>
    <Dialog isShown={isShown}>
      { renderFile() }
    </Dialog>
  </Pane>
}
