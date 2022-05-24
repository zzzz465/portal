import { RecordItemType } from './recordItem'

type _Record = {
  name: string
  tags: string[]
  metadata: Metadata
}

export type Metadata = {
  dataSource: string
  labels: Record<string, string>
}

export type KnownLabels = Partial<{
  recordType: RecordItemType['type']
}>

// conflict with es5.Record
export type { _Record as Record }
