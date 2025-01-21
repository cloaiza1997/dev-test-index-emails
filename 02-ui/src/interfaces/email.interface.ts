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

export interface EmailHighlightInterface {
  email: EmailInterface
  highlight: {
    bcc?: string[]
    body?: string[]
    cc?: string[]
    from?: string[]
    subject?: string[]
    to?: string[]
    xBcc?: string[]
    xCc?: string[]
    xFrom?: string[]
    xTo?: string[]
  } | null
}

export interface EmailSearchResponseInterface {
  pagination: PaginationInterface
  items: EmailHighlightInterface[]
}

export interface PaginationInterface {
  total: number
  count: number
  pages: number
  prev: number
  next: number
}
