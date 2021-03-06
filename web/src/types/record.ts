type _Record = {
  name: string
  tags: string[]
  metadata: Metadata
}

export type Metadata = {
  dataSource: string
  labels: Record<string, string> & KnownLabels
}

export type KnownLabels = Partial<{
  recordType: 'well-known'
}>

// conflict with es5.Record
export type { _Record as Record }
