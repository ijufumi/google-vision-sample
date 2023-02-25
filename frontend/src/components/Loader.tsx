import React, { FC } from 'react'
import { Overlay, Spinner, Pane } from "evergreen-ui"
import { Watch} from "react-loader-spinner"

export interface Props{
  isShown: boolean
}

const Loader: FC<Props> = ({ isShown }) => {
  return (
    <Overlay
    isShown={isShown}
    shouldCloseOnClick={false}
    shouldCloseOnEscapePress={false}
    containerProps={
      {
        style: {
          "backgroundColor": "rgba(67, 90, 111, 0.15)"
        }
      }
    }
  >
      <Pane display="flex" alignItems="center" justifyContent="center" width="100%" height="100%">
        <Watch
          height="100"
          width="100"
          radius="48"
          ariaLabel="watch-loading"
          wrapperStyle={{
            "zIndex": "100"
          }}
          visible={true}
        />
      </Pane>
  </Overlay>
  )
}

export default Loader
