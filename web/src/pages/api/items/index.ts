import axios, { AxiosError } from 'axios'
import _ from 'lodash'
import { NextApiRequest, NextApiResponse } from 'next'
import { DefaultDictionary } from 'typescript-collections'
import { Record } from '../../../types/record'
import { DisplayableRecordItem, GroupedRecordItem, RecordItem, WellKnownRecordItem } from '../../../types/recordItem'
import { tryP } from '../../../utils/try'
import { BACKEND_SERVER_URL } from '../const'

type RecordItemsResponse = {
  items: DisplayableRecordItem[]
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<RecordItemsResponse | Error | undefined>
) {
  const [recordsRes, err0] = await tryP(() => axios.get<Record[]>('/records', { baseURL: BACKEND_SERVER_URL }))
  if (err0 !== null) {
    return res.status(500).send((err0 as AxiosError).cause)
  }

  const records = recordsRes.data
  const recordItems = toRecordItems(records)
  const groupRecordItems = makeGroupRecordItems(recordItems)
  const wellKnownRecordItems = makeWellKnownRecordItems(recordItems)

  res.json({ items: [...wellKnownRecordItems, ...groupRecordItems] })
}

function toRecordItems(items: Record[]): RecordItem[] {
  return items.map((rec) => ({
    data: rec,
  }))
}

function makeWellKnownRecordItems(records: RecordItem[]): WellKnownRecordItem[] {
  // Type Difference is not implemented. see: https://github.com/microsoft/TypeScript/issues/4183
  return records
    .filter((rec) => rec.data.metadata.labels.recordType === 'well-known')
    .map((rec) =>
      Object.assign<RecordItem, WellKnownRecordItem>(rec, { type: 'wellKnownRecordItem', name: rec.data.name } as any)
    )
}

function makeGroupRecordItems(records: RecordItem[]): GroupedRecordItem[] {
  const group: DefaultDictionary<string, RecordItem[]> = new DefaultDictionary(() => [])
  for (const item of records) {
    for (const tag of _.uniq(item.data.tags)) {
      group.getValue(tag).push(item)
    }
  }

  return group.keys().map((k) => ({
    type: 'groupedRecordItem',
    items: group.getValue(k),
    name: k,
  }))
}
