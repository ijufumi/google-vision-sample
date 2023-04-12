import React, { FC, useRef, useEffect, useState } from "react"
import { Rect as KonvaRect, Transformer } from "react-konva"
import Konva from "konva"

export interface Props extends Konva.RectConfig {

}

const Rect: FC<Props> = (props) => {
  const [selected, setSelected] = useState(false)
  const tfRef = useRef<Konva.Transformer>(null)
  const rectRef = useRef<Konva.Rect>(null)

  useEffect(() => {
    if (tfRef.current && rectRef.current) {
      if (selected && (props.visible === true)) {
        tfRef.current.nodes([rectRef.current])
        tfRef.current.update()
      } else {
        setSelected(false)
        tfRef.current.nodes([])
        tfRef.current.update()
      }
    }
  }, [selected, props])

  return <React.Fragment>
    <KonvaRect
      draggable={true}
      ref={rectRef}
      stroke={"#D14343"}
      fill={"#F9DADA"}
      opacity={0.3}
      strokeWidth={2}
      {...props}
      onClick={() => setSelected(!selected)}
    />
    <Transformer ref={tfRef} />
  </React.Fragment>
}

export default Rect
