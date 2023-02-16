import React, { FC, useMemo, useEffect, useState } from 'react'
import { Dialog } from "evergreen-ui"

export interface Props {
  extractionResultId: string
  onClose: () => {}
}

const ResultViewer: FC<Props> = ({ extractionResultId, onClose }) => {
  const isShown = useMemo(() => !!extractionResultId, [extractionResultId])


  return (
    <Dialog
      isShown={isShown}
      onCloseComplete={onClose}
      hasCancel={false}
      confirmLabel="Close"
    >

    </Dialog>
  )
}

export default ResultViewer
