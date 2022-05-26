import { Modal } from '@chakra-ui/react'
import { from } from 'linq-es2015'
import { GetServerSidePropsResult } from 'next'
import RecordList from '../../components/recordList'
import { DisplayableRecordItem, GroupedRecordItem } from '../../types/recordItem'
import { tryP } from '../../utils/try'

type P = {
  items: DisplayableRecordItem[]
  error?: Error
}

export async function getServerSideProps(): Promise<GetServerSidePropsResult<P>> {
  const [res, err0] = await tryP(() => fetch('http://localhost:3000/api/items'))
  if (err0 !== null) {
    return { props: { items: [], error: err0 as Error } }
  }

  const data = await res.json()
  return { props: data }
}

export default function Page(props: P) {
  const item = from(props.items)
    .Where((item) => item.type === 'groupedRecordItem')
    .Cast<GroupedRecordItem>()
    .OrderByDescending((item) => item.items.length)
    .First()

  return (
    <main>
      <Modal isOpen={true} onClose={() => {}} size="4xl">
        <RecordList items={item.items} />
      </Modal>
    </main>
  )
}
