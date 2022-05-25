import axios, { AxiosError } from 'axios'
import { from } from 'linq-es2015'
import _ from 'lodash'
import { NextApiRequest, NextApiResponse } from 'next'
import { DefaultDictionary } from 'typescript-collections'
import { Record } from '../../../types/record'
import {
  DisplayableRecordItemType,
  GroupedRecordItem,
  RecordItem,
  WellKnownRecordItem,
} from '../../../types/recordItem'
import { tryP } from '../../../utils/try'
import { BACKEND_SERVER_URL } from '../const'
import { RecordsResponse } from '../records'

type RecordItemsResponse = {
  items: DisplayableRecordItemType[]
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<RecordItemsResponse | Error | undefined>
) {
  const [recordsRes, err0] = await tryP(() => axios.get<RecordsResponse>('/records', { baseURL: BACKEND_SERVER_URL }))
  if (err0 !== null) {
    return res.status(500).send((err0 as AxiosError).cause)
  }

  const records = toRecordItem(recordsRes.data)
  const groupRecordItems = makeGroupRecordItems(records)

  const wellKnownRecordItems = from(records)
    .Where((rec) => rec.data.metadata.labels.recordType === 'groupedRecordItem')
    .Cast<WellKnownRecordItem>()
    .ToArray()

  res.json({ items: [...wellKnownRecordItems, ...groupRecordItems] })
}

function toRecordItem(records: Record[]): RecordItem[] {
  return records.map((rec) => ({
    type: 'RecordItem',
    data: rec,
  }))
}

function makeGroupRecordItems(records: RecordItem[]): GroupedRecordItem[] {
  const group: DefaultDictionary<string, RecordItem[]> = new DefaultDictionary(() => [])
  for (const rec of records) {
    for (const tag of _.uniq(rec.data.tags)) {
      group.getValue(tag).push(rec)
    }
  }

  return group.keys().map((k) => ({
    type: 'groupedRecordItem',
    items: group.getValue(k),
    tag: k,
  }))
}
