export type Message = {
  createdAt: string
  params?: Params
  sentByUser: boolean
  text: string
}

type Params = {
  score?: number
  suggestedQuestion?: number
}
