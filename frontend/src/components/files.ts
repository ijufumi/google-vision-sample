export const readAsFile = async (url: string) => {
  const response = await fetch(url)
  const data = await response.blob()
  let metadata = {
    type: data.type,
  }
  return new File([data], data.name, metadata)
}

export const readAsBlob = async (url: string) => {
  const response = await fetch(url)
  return await response.blob()
}

export const readAsText = async (blob: Blob) => {
  return await blob.text()
}

export const isTextType = (contentType: string) => {
  return ["application/json"].includes(contentType)
}
