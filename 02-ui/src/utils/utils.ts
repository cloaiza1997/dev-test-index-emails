import { format } from '@formkit/tempo'

export function getDate(date: string): string {
  return format(new Date(date), 'YYYY-MM-DD')
}

export function replaceHighlight(
  text: string | null | undefined,
  highlights: string[] | null | undefined,
) {
  let result = text?.replace(/\n/gim, '</br>') ?? ''

  if (!highlights) {
    return result
  }

  highlights.forEach((highlight) => {
    const _highlight = highlight.replace(/…/gim, '').replace(/\n/gim, '</br>')
    const searchValue = _highlight.replace(/<mark>|<\/mark>/gim, '').trim()

    result = result.replace(searchValue, _highlight)
  })

  return result
}
