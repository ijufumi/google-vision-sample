import React, { FC } from "react"
import { Overlay } from "evergreen-ui"
import Loader from "./Loader"

export interface Props {
  isShown: boolean
}

const LoaderOverlay: FC<Props> = ({ isShown }) => {
  return (
    <Overlay
      isShown={isShown}
      shouldCloseOnClick={false}
      shouldCloseOnEscapePress={false}
      containerProps={{
        style: {
          backgroundColor: "rgba(67, 90, 111, 0.15)",
        },
      }}
    >
      <Loader />
    </Overlay>
  )
}

export default LoaderOverlay
