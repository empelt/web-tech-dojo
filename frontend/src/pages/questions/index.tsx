import { useEffect, useMemo, useState } from 'react'

import axios, { AxiosResponse } from 'axios'

import QuestionsTable from './components/table'
import Toolbar from './components/toolbar'

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

  const onBookmark = (id: string) => {
    setQuestions((prevQuestions) =>
      prevQuestions.map((question) =>
        question.id === id
          ? { ...question, isBookmarked: !question.isBookmarked }
          : question,
      ),
    )
  }

  useEffect(() => {
    setFilteredQuestions(
      questions.filter((question) => {
        if (title) {
          return question.title.includes(title)
        }
        if (selectedTagsValues.size > 0) {
          return question.tags.some((tag) => selectedTagsValues.has(tag))
        }
        if (selectedBookmarkValues.size > 0) {
          switch (question.isBookmarked) {
            case true:
              return selectedBookmarkValues.has('isBookmarked')
            default:
              return selectedBookmarkValues.has('isNotBookmarked')
          }
        }
        if (selectedProgressValues.size > 0) {
          switch (question.progress) {
            case 0:
              return selectedProgressValues.has('todo')
            case 100:
              return selectedProgressValues.has('completed')
            default:
              return selectedProgressValues.has('inProgress')
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
      <h1 className="text-2xl font-bold mt-24 mb-4">Questions</h1>
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
