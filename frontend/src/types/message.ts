export type Message = {
  createdAt: string
  params?: Params
  sentByUser: boolean
  text: string
}

export type MessageResponse = {
  message: string
  score: number
  suggestedQuestionId: number
}

type Params = {
  score?: number
  suggestedQuestion?: number
}
