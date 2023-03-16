import React, { FC } from "react"
import { Watch } from "react-loader-spinner"
import { Pane } from "evergreen-ui"

export interface Props{}

const Loader: FC<Props> = () => {
  return (
    <Pane
      display="flex"
      alignItems="center"
      justifyContent="center"
      width="100%"
      height="100%"
    >
      <Watch
        height="100"
        width="100"
        radius="48"
        ariaLabel="watch-loading"
        wrapperStyle={{
          zIndex: "100",
        }}
        visible={true}
      />
    </Pane>
  )
}

export default Loader
