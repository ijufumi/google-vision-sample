import React, { FC, useEffect, useState } from "react"
import { Image as KonvaImage } from "react-konva"
import useImage from "use-image"

export interface Props {
  outerWidth: number
  outerHeight: number
  url: string
  onLoaded: () => void
}

const Image: FC<Props> = ({ url , outerWidth, outerHeight, onLoaded}) => {
  const [image, status] = useImage(url)
  const [scale, setScale] = useState<number>(1)

  useEffect(() => {
    if (status !== 'loaded' || !image) {
      return
    }
    const imageWidth = image.naturalWidth
    const imageHeight = image.naturalHeight
    const scaleWidth = outerWidth / imageWidth
    const scaleHeight = outerHeight / imageHeight
    const _scale = Math.min(scaleHeight, scaleWidth)

    setScale(_scale)
    onLoaded()
  }, [image, status])

  return <KonvaImage image={image} scaleY={scale} scaleX={scale} />
}

export default Image
