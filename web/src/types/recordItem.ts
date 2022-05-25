import { Record } from './record'

export interface Displayable {
  name: string
}

export type DisplayableRecordItem = WellKnownRecordItem | GroupedRecordItem

/**
 * RecordItem is a basic element that holds Record object.
 */
export interface RecordItem {
  data: Record
}

/**
 * WellKnownRecordItem is an element that can be displayed alone.
 */
export interface WellKnownRecordItem extends RecordItem, Displayable {
  type: 'wellKnownRecordItem'
}

/**
 * GroupedRecordItem is an element that holds multiple RecordItem
 */
export interface GroupedRecordItem extends Displayable {
  type: 'groupedRecordItem'
  items: RecordItem[]
}
