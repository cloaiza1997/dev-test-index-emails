export interface HttpResponseInterface<T> {
  success: boolean
  message: string
  data: T
}
