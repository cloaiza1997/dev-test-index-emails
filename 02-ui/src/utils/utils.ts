import { format } from '@formkit/tempo'

export function getDate(date: string | undefined | null): string {
  if (!date) {
    return ''
  }

  return format(new Date(date), 'YYYY-MM-DD')
}

export function replaceHighlight(
  text: string | null | undefined,
  highlights: string[] | null | undefined,
) {
  let result = text?.replace(/(\r)?\n/gim, '</br>') ?? ''

  if (!highlights) {
    return result
  }

  result = decodeHtmlEntities(result)

  highlights.forEach((highlight) => {
    const _highlight = decodeHtmlEntities(
      highlight.replace(/…/gim, '').replace(/(\r)?\n/gim, '</br>'),
    ).trim()

    const searchValue = _highlight.replace(/<mark>|<\/mark>/gim, '')

    result = result.replace(searchValue, _highlight)
  })

  return result
    .replace(/<mark>/gim, '+[+mark+]+')
    .replace(/<\/mark>/gim, '+[+/mark+]+')
    .replace(/<strong>/gim, '+[+strong+]+')
    .replace(/<\/strong>/gim, '+[+/strong+]+')
    .replace(/<\/br>/gim, '+[+/br+]+')
    .replace(/</gim, '&lt;')
    .replace(/>/gim, '&gt;')
    .replace(/\+\[\+/gim, '<')
    .replace(/\+\]\+/gim, '>')
}

function decodeHtmlEntities(text: string) {
  const textArea = document.createElement('textarea')
  textArea.innerHTML = text

  return textArea.value
}
