import { useMemo, useState } from 'react'

import QuestionsTable from './components/table'
import Toolbar from './components/toolbar'

import { Question } from '@/types/question'

const questions: Question[] = [
  {
    id: '1',
    title: 'How to use React Query?',
    tags: ['react', 'query'],
    isBookmarked: false,
    progress: 0,
  },
  {
    id: '2',
    title: 'How to use React Table?',
    tags: ['react', 'table'],
    isBookmarked: false,
    progress: 20,
  },
  {
    id: '3',
    title: 'How to use React Hook Form?',
    tags: ['react', 'hook', 'form'],
    isBookmarked: false,
    progress: 80,
  },
]

const Questions = () => {
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

  return (
    <div className="container mx-auto">
      <h1 className="text-2xl font-bold my-4">Questions</h1>
      <div className="flex flex-col gap-4">
        <Toolbar filterState={filterState} />
        <QuestionsTable questions={questions} />
      </div>
    </div>
  )
}

export default Questions
