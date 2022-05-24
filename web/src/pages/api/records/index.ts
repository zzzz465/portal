import axios, { AxiosError } from 'axios'
import { NextApiRequest, NextApiResponse } from 'next'
import { Record } from '../../../types/record'
import { tryP } from '../../../utils/try'
import { BACKEND_SERVER_URL } from '../const'

export type RecordsResponse = Record[]

export default async function handler(req: NextApiRequest, res: NextApiResponse<RecordsResponse | any>) {
  const [recordsResponse, err0] = await tryP(() =>
    axios.get<RecordsResponse>('/records', { baseURL: BACKEND_SERVER_URL })
  )
  if (err0 != null) {
    return res.status(500).send((err0 as AxiosError).cause)
  }

  const records = recordsResponse.data

  return res.status(200).json({ records })
}
