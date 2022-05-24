import axios, { AxiosError } from 'axios'
import { NextApiRequest, NextApiResponse } from 'next'
import { tryP } from '../../../utils/try'
import { BACKEND_SERVER_URL } from '../const'

export type DatasourcesResponse = {
  datasources: string[]
}

export default async function handler(req: NextApiRequest, res: NextApiResponse<DatasourcesResponse | string>) {
  const [dsRes, err0] = await tryP(() => axios.get('/datasources', { baseURL: BACKEND_SERVER_URL }))
  if (err0 !== null) {
    return res.status(500).send((err0 as AxiosError).message)
  }

  const datasources: string[] = dsRes.data.dataSources

  return res.status(200).json({ datasources })
}
