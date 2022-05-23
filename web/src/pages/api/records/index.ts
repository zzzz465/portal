import { NextApiRequest, NextApiResponse } from 'next'
import { Record } from '../../../types/record'

export type RecordsResponse = {
  records: Record[]
}

export default async function handler(req: NextApiRequest, res: NextApiResponse<RecordsResponse>) {
  return res.status(200).json({
    records: [],
  })
}
