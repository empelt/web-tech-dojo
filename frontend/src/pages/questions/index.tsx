import { useEffect, useMemo, useState } from 'react'

import axios, { AxiosResponse } from 'axios'

import QuestionsTable from './components/table'
import Toolbar from './components/toolbar'

import { toast } from '@/hooks/use-toast'
import { Question } from '@/types/question'

const QuestionsPage = () => {
  const [loading, setLoading] = useState<boolean>(false)
  const [questions, setQuestions] = useState<Question[]>([])
  const [filteredQuestions, setFilteredQuestions] = useState<Question[]>([])
  const [title, setTitle] = useState<string>('')
  const [selectedTagsValues, setSelectedTagsValues] = useState<Set<string>>(
    new Set(),
  )
  const [selectedBookmarkValues, setSelectedBookmarkValues] = useState<
    Set<string>
  >(new Set())
  const [selectedProgressValues, setSelectedProgressValues] = useState<
    Set<string>
  >(new Set())

  const filterState = useMemo(
    () => ({
      title,
      selectedTagsValues,
      selectedBookmarkValues,
      selectedProgressValues,
      setTitle,
      setSelectedTagsValues,
      setSelectedBookmarkValues,
      setSelectedProgressValues,
    }),
    [title, selectedTagsValues, selectedBookmarkValues, selectedProgressValues],
  )

  useEffect(() => {
    setLoading(true)
    // 2回実行されるのを防ぐ
    // see https://react.dev/learn/synchronizing-with-effects#fetching-data
    let ignore = false
    axios
      .get('/api/question')
      .then((res: AxiosResponse<Question[]>) => {
        const { data, status } = res
        if (!ignore && status === 200) {
          setFilteredQuestions(data)
          setQuestions(data)
        }
      })
      .catch((error) => {
        console.error(error)
      })
      .finally(() => {
        setLoading(false)
      })
    return () => {
      ignore = true
    }
  }, [])

  const onBookmark = async (id: string, isBookmarked: boolean) => {
    try {
      if (isBookmarked) {
        const { status }: AxiosResponse<{ message: string }> =
          await axios.delete(`/api/bookmark/question/${id}`)
        if (status === 200) {
          setQuestions((prevQuestions) =>
            prevQuestions.map((question) =>
              question.id === id
                ? { ...question, isBookmarked: !isBookmarked }
                : question,
            ),
          )
        }
      } else {
        const { status }: AxiosResponse<{ message: string }> = await axios.post(
          `/api/bookmark/question/${id}`,
        )
        if (status === 200) {
          setQuestions((prevQuestions) =>
            prevQuestions.map((question) =>
              question.id === id
                ? { ...question, isBookmarked: !isBookmarked }
                : question,
            ),
          )
        }
      }
    } catch {
      toast({
        variant: 'destructive',
        title: 'Uh oh! Something went wrong.',
        description: 'There was a problem with your request.',
      })
    }
  }

  useEffect(() => {
    setFilteredQuestions(
      questions.filter((question) => {
        if (title && !question.title.includes(title)) {
          return false
        }
        if (
          selectedTagsValues.size > 0 &&
          !question.tags.some((tag) => selectedTagsValues.has(tag))
        ) {
          return false
        }
        if (selectedBookmarkValues.size > 0) {
          if (
            question.isBookmarked &&
            !selectedBookmarkValues.has('isBookmarked')
          ) {
            return false
          }
          if (
            !question.isBookmarked &&
            !selectedBookmarkValues.has('isNotBookmarked')
          ) {
            return false
          }
        }
        if (selectedProgressValues.size > 0) {
          if (question.progress <= 0 && !selectedProgressValues.has('todo')) {
            return false
          }
          if (
            question.progress === 100 &&
            !selectedProgressValues.has('completed')
          ) {
            return false
          }
          if (
            question.progress > 0 &&
            question.progress < 100 &&
            !selectedProgressValues.has('inProgress')
          ) {
            return false
          }
        }
        return true
      }),
    )
  }, [
    title,
    selectedTagsValues,
    selectedBookmarkValues,
    selectedProgressValues,
    questions,
  ])

  return (
    <div className="container mx-auto">
      <h1 className="text-2xl font-bold mt-24 mb-4">問題一覧</h1>
      <div className="flex flex-col gap-4">
        <Toolbar filterState={filterState} />
        <QuestionsTable
          loading={loading}
          onBookmark={onBookmark}
          questions={filteredQuestions}
        />
      </div>
    </div>
  )
}

export default QuestionsPage
