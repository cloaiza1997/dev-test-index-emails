import axios, { type AxiosResponse } from 'axios'

import { environment } from '@/config/environment'
import type { EmailSearchResponseInterface } from '@/interfaces/email.interface'
import type { HttpResponseInterface } from '@/interfaces/http.interface'

export function searchEmails(term: string, page: number, limit: number) {
  return axios
    .get<HttpResponseInterface<EmailSearchResponseInterface>>(`${environment.HOST}/v1/emails`, {
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
