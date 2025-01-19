export interface EmailInterface {
  messageId: string
  date: string
  from: string
  to: string
  cc: string
  bcc: string
  subject: string
  xFrom: string
  xTo: string
  xCc: string
  xBcc: string
  xFolder: string
  xOrigin: string
  xFileName: string
  body: string
  path: string
  mainFolder: string
}

export interface EmailSearchResponseInterface {
  pagination: PaginationInterface
  items: EmailInterface[]
}

export interface PaginationInterface {
  total: number
  count: number
  pages: number
  prev: number
  next: number
}
