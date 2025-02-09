import { useEffect, useState } from 'react'

import axios, { AxiosResponse } from 'axios'
import { useNavigate } from 'react-router'

import LoadingButton from '@/components/LoadingButton'
import { Button } from '@/components/ui/button'
import { Question } from '@/types/question'

type Props = {
  questionId: number
}
const SuggestMessage = ({ questionId }: Props) => {
  const navigate = useNavigate()
  const [loading, setLoading] = useState(false)
  const [title, setTitle] = useState('')

  useEffect(() => {
    if (!questionId) {
      return
    }
    let ignore = false
    ;(async () => {
      setLoading(true)
      // 2回実行されるのを防ぐ
      // see https://react.dev/learn/synchronizing-with-effects#fetching-data

      try {
        // 問題のタイトルを取得
        const {
          data: questionResponse,
          status: questionResponseStatus,
        }: AxiosResponse<Question> = await axios.get(
          `/api/question/${questionId}`,
        )
        if (!ignore && questionResponseStatus === 200) {
          setTitle(questionResponse.title)
        }
      } catch {
        setTitle('問題が見つかりませんでした')
      } finally {
        setLoading(false)
      }
    })()
    return () => {
      ignore = true
    }
  }, [questionId])

  if (loading) {
    return (
      <div className="mx-auto mb-2 w-full text-left max-w-[800px]">
        <div className="ml-12 bg-gray-300 p-4 rounded-lg inline-block">
          <p className="font-bold">師範からの提案</p>
          <p>
            まだこの問題に挑むときではないかもしれない。こんな問題はどうじゃ？
          </p>
          <LoadingButton className="w-52 mt-2" variant="secondary" />
        </div>
      </div>
    )
  }

  return (
    <div className="mx-auto mb-2 w-full text-left max-w-[800px]">
      <div className="ml-12 bg-gray-300 p-4 rounded-lg inline-block">
        <p className="font-bold">師範からの提案</p>
        <p>
          まだこの問題に挑むときではないかもしれない。こんな問題はどうじゃ？
        </p>
        <Button
          className="mt-2"
          onClick={() => navigate(`/questions/${questionId}`)}
          variant="secondary">
          <div className="flex items-center">
            <div className="py-1 px-2 bg-black text-white rounded-lg mr-2">
              Q{questionId}
            </div>
            <h1>{title}</h1>
          </div>
        </Button>
      </div>
    </div>
  )
}

export default SuggestMessage
