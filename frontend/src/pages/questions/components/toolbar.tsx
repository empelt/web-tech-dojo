import { X } from 'lucide-react'

import Filter from './filter'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { bookmarkStatus, progressStatus, tags } from '@/types/question'

type Props = {
  filterState: {
    selectedBookmarkValues: Set<string>
    selectedProgressValues: Set<string>
    selectedTagsValues: Set<string>
    setSelectedBookmarkValues: React.Dispatch<React.SetStateAction<Set<string>>>
    setSelectedProgressValues: React.Dispatch<React.SetStateAction<Set<string>>>
    setSelectedTagsValues: React.Dispatch<React.SetStateAction<Set<string>>>
    setTitle: React.Dispatch<React.SetStateAction<string>>
    title: string
  }
  onClickSearch: (title: string) => void
}

const Toolbar = ({
  filterState: {
    title,
    selectedTagsValues,
    selectedBookmarkValues,
    selectedProgressValues,
    setTitle,
    setSelectedTagsValues,
    setSelectedBookmarkValues,
    setSelectedProgressValues,
  },
  onClickSearch,
}: Props) => {
  const isFiltered =
    selectedTagsValues.size > 0 ||
    selectedBookmarkValues.size > 0 ||
    selectedProgressValues.size > 0

  return (
    <>
      <div className="flex w-full max-w-sm items-center space-x-2">
        <Input
          className="h-8 w-[150px] lg:w-[250px]"
          onChange={(event) => setTitle(event.target.value)}
          placeholder="Filter questions..."
          value={title}
        />
        <Button onClick={() => onClickSearch(title)} type="submit">
          search
        </Button>
      </div>
      <div className="flex items-center justify-between">
        <div className="flex flex-1 items-center space-x-2">
          <Filter
            options={tags}
            selectedValues={selectedTagsValues}
            setSelectedValues={setSelectedTagsValues}
            showInput
            title="Tags"
          />
          <Filter
            options={bookmarkStatus}
            selectedValues={selectedBookmarkValues}
            setSelectedValues={setSelectedBookmarkValues}
            title="Bookmark"
          />
          <Filter
            options={progressStatus}
            selectedValues={selectedProgressValues}
            setSelectedValues={setSelectedProgressValues}
            title="Progress"
          />
          {isFiltered && (
            <Button
              className="h-8 px-2 lg:px-3"
              onClick={() => {
                setSelectedTagsValues(new Set())
                setSelectedBookmarkValues(new Set())
                setSelectedProgressValues(new Set())
              }}
              variant="ghost">
              Reset
              <X />
            </Button>
          )}
        </div>
      </div>
    </>
  )
}

export default Toolbar
