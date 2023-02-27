export const readAsFile = async (url: string) => {
  const response = await fetch(url);
  const data = await response.blob();
  let metadata = {
    type: data.type
  };
  return new File([data], data.name, metadata);
}
