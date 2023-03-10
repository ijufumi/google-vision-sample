import React, { FC, useEffect, useState } from "react"
import { Image as KonvaImage } from "react-konva"
import useImage from "use-image"

export interface Props {
  outerWidth: number
  outerHeight: number
  url: string
}

const Image: FC<Props> = ({ url , outerWidth, outerHeight}) => {
  const [image, loaded] = useImage(url)
  const [scale, setScale] = useState<number>(1)

  useEffect(() => {
    if (!loaded || !image) {
      return
    }
    const imageWidth = image.width
    const imageHeight = image.height
    const scaleWidth = imageWidth / outerWidth
    const scaleHeight = imageHeight /outerHeight
    const _scale = Math.min(scaleHeight, scaleWidth)
    setScale(_scale)
  }, [image, loaded])

  if (!loaded || !image) {
    return null
  }
  return <KonvaImage image={image} scaleY={scale} scaleX={scale} />
}

export default Image
