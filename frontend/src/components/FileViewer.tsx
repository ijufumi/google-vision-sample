import React, { FC } from 'react'
import { Pane, Dialog } from "evergreen-ui"


export interface Props {
  key: string
  isShown: boolean
}

const FileViewer: FC<Props> = ({ key, isShown }) => {

  return <Pane>
    <Dialog isShown={isShown}>

    </Dialog>
  </Pane>
}
