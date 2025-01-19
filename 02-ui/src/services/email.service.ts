import axios, { type AxiosResponse } from 'axios'

import type { EmailSearchResponseInterface } from '@/interfaces/email.interface'
import type { HttpResponseInterface } from '@/interfaces/http.interface'

const HOST = 'http://localhost:8000'

export function searchEmails(term: string, page: number, limit: number) {
  return axios
    .get<HttpResponseInterface<EmailSearchResponseInterface>>(`${HOST}/v1/emails`, {
      params: {
        term,
        page,
        limit,
      },
    })
    .catch((error) => {
      console.error(error)

      return error as AxiosResponse<HttpResponseInterface<EmailSearchResponseInterface>>
    })
}
