import { Record } from './record'

export type RecordItemType = RecordItem | GroupedRecordItem | WellKnownRecordItem
export type DisplayableRecordItemType = WellKnownRecordItem | GroupedRecordItem

/**
 * RecordItem is an element that is grouped by tags.
 */
export interface RecordItem {
  type: 'RecordItem'
  data: Record
}

/**
 * WellKnownRecordItem is an element that can be displayed alone.
 */
export type WellKnownRecordItem = Omit<RecordItem, 'type'> & {
  type: 'wellKnownRecordItem'
}

/**
 * GroupedRecordItem is an element that holds multiple RecordItem
 */
export interface GroupedRecordItem {
  type: 'groupedRecordItem'
  tag: string
  items: RecordItem[]
}
