import React, { FC, useState, useEffect } from "react"
import { Pane, Button, UploadIcon, Text, Table } from "evergreen-ui"
import ExtractionResult from "../models/ExtractionResult"
import ExtractionUseCaseImpl from "../usecases/ExtractionUseCase"

interface Props {
}

const App: FC<Props> = () => {
  const [initialized, setInitialized] = useState<boolean>(false)
  const [extractionResults, setExtractionResults] = useState<ExtractionResult[]>([])

  const useCase = new ExtractionUseCaseImpl()

  useEffect(() => {
    const initialize = async () => {
      const _extractionResults = await useCase.getExtractionResults()
      if (_extractionResults) {
        setExtractionResults(_extractionResults)
      } else {
        console.error("something wrong...")
      }
      setInitialized(true)
    }
    initialize()
  }, [])

  if (!initialized) {
    return null
  }

  const renderResults = () => {
    if (extractionResults.length === 0) {
      return (
        <Pane>
          <Text>There was no results</Text>
        </Pane>
      )
    }
    return (
      <Pane>
        <Table>
          <Table.Head>
            <Table.TextHeaderCell>ID</Table.TextHeaderCell>
            <Table.TextHeaderCell>Status</Table.TextHeaderCell>
            <Table.TextHeaderCell>Input</Table.TextHeaderCell>
            <Table.TextHeaderCell>Output</Table.TextHeaderCell>
            <Table.TextHeaderCell>CreatedAt</Table.TextHeaderCell>
            <Table.TextHeaderCell>UpdatedAt</Table.TextHeaderCell>
            <Table.TextHeaderCell>Operations</Table.TextHeaderCell>
          </Table.Head>
          <Table.VirtualBody>
            {extractionResults.map(result => (
              <Table.Row key={result.id}>
                <Table.TextCell>{result.id}</Table.TextCell>
                <Table.TextCell>{result.status}</Table.TextCell>
                <Table.TextCell>{result.imageUri}</Table.TextCell>
                <Table.TextCell>{result.outputUri}</Table.TextCell>
                <Table.TextCell>{result.createdAt}</Table.TextCell>
                <Table.TextCell>{result.updatedAt}</Table.TextCell>
                <Table.Cell>

                </Table.Cell>
              </Table.Row>
            ))}
          </Table.VirtualBody>
        </Table>
      </Pane>
    )
  }

  return <Pane display="flex" flexDirection="column">
    <Pane display="flex" alignContent="flex-end">
      <Button appearance="primary" iconAfter={UploadIcon}>
        Upload
      </Button>
    </Pane>
    {renderResults()}
  </Pane>
}

export default App
