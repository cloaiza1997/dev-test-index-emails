import { format } from '@formkit/tempo'

export function getDate(date: string): string {
  return format(new Date(date), 'YYYY-MM-DD')
}
