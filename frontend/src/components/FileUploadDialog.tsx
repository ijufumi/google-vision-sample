import React, { FC, useState, useCallback } from 'react'
import { Dialog, Pane, FileUploader, FileCard, FileRejection, Button } from "evergreen-ui"

interface Props {
  isShown: boolean
  onClose: () => void
  onUpload:  (file: File) => void
}

const FileUploadDialog: FC<Props> = ({ isShown, onClose, onUpload }) => {
  const [files, setFiles] = useState<File[]>([])
  const [fileRejections, setFileRejections] = useState<FileRejection[]>([])
  const handleChange = useCallback((files: File[]) => setFiles([files[0]]), [])
  const handleRejected = useCallback((fileRejections: FileRejection[]) => setFileRejections([fileRejections[0]]), [])
  const handleRemove = useCallback(() => {
    setFiles([])
    setFileRejections([])
  }, [])

  const handleUpload = () => {
    if (files.length > 0) {
      onUpload(files[0])
    }
  }

  const handleClose = () => {
    setFiles([])
    onClose()
  }

  return <Pane>
    <Dialog
      isShown={isShown}
      title="File uploading"
      onCloseComplete={handleClose}
      confirmLabel="Upload"
      cancelLabel="Close"
      onConfirm={handleUpload}
      isConfirmDisabled={files.length === 0}
    >
      <FileUploader
        label="Upload File"
        description="You can upload 1 file. File can be up to 50 MB."
        maxSizeInBytes={50 * 1024 ** 2}
        maxFiles={1}
        onChange={handleChange}
        onRejected={handleRejected}
        renderFile={(file) => {
          const { name, size, type } = file
          const fileRejection = fileRejections.find((fileRejection) => fileRejection.file === file)
          const { message } = fileRejection || {}
          return (
            <FileCard
              key={name}
              isInvalid={fileRejection != null}
              name={name}
              onRemove={handleRemove}
              sizeInBytes={size}
              type={type}
              validationMessage={message}
            />
          )
        }}
        values={files}
      />
    </Dialog>
  </Pane>
}

export default FileUploadDialog
