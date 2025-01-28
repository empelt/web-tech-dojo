export type Message = {
  createdAt: string
  params?: Params
  sentByAI: boolean
  text: string
}

type Params = {
  progress?: number
  sugested_questions?: string[]
}
