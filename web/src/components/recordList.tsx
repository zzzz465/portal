import {
  Input,
  ModalBody,
  ModalContent,
  ModalHeader,
  Table,
  TableContainer,
  Tbody,
  Text,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react'
import { useState } from 'react'
import { RecordItem } from '../types/recordItem'

type P = {
  items: RecordItem[]
}

export default function RecordList({ items }: P) {
  const [search, setSearch] = useState('')

  const filterItem = (item: RecordItem) => item.data.name.includes(search)
  const toName = (item: RecordItem) => (
    <a className="block" href={`http://${item.data.name}`}>
      <Text>{item.data.name}</Text>
    </a>
  )
  const toTag = (item: RecordItem) => <></>
  const toAction = (item: RecordItem) => <></>

  items = items.filter(filterItem)

  return (
    <ModalContent>
      <ModalHeader>
        <Input placeholder="Search" value={search} onChange={(e) => setSearch(e.target.value)} />
      </ModalHeader>
      <ModalBody>
        <TableContainer>
          <Table>
            <Thead>
              <Tr>
                <Th>Name</Th>
                <Th>Tags</Th>
                <Th>Actions</Th>
              </Tr>
            </Thead>
            <Tbody>
              <Tr id="Name">{items.map(toName)}</Tr>

              <Tr id="Tags">{items.map(toTag)}</Tr>

              <Tr id="Actions">{items.map(toAction)}</Tr>
            </Tbody>
          </Table>
        </TableContainer>
      </ModalBody>
    </ModalContent>
  )
}
