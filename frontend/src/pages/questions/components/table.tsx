import { MoreHorizontal } from 'lucide-react'
import { CiHeart } from 'react-icons/ci'
import { FaHeart } from 'react-icons/fa'
import { useNavigate } from 'react-router'

import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Progress } from '@/components/ui/progress'
import { Skeleton } from '@/components/ui/skeleton'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Question } from '@/types/question'

type Props = {
  loading: boolean
  onBookmark: (id: string, isBookmarked: boolean) => void
  questions: Question[]
}

const QuestionsTable = ({ questions, onBookmark, loading }: Props) => {
  const navigate = useNavigate()

  if (loading) {
    return (
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[100px] text-center">Bookmark</TableHead>
              <TableHead>No.</TableHead>
              <TableHead>Title</TableHead>
              <TableHead>Tags</TableHead>
              <TableHead>Progress</TableHead>
              <TableHead />
            </TableRow>
          </TableHeader>
          <TableBody>
            {[...Array(10)].map((_, index) => (
              <TableRow key={index}>
                <TableCell className="text-center">
                  <Skeleton className="w-6 h-6 mx-auto" />
                </TableCell>
                <TableCell>
                  <Skeleton className="w-10 h-6" />
                </TableCell>
                <TableCell>
                  <Skeleton className="w-52 h-6" />
                </TableCell>
                <TableCell>
                  <div className="flex w-[100px] items-center">
                    <Skeleton className="w-10 h-4 mr-2 rounded-full" />
                    <Skeleton className="w-10 h-4 mr-2 rounded-full" />
                  </div>
                </TableCell>
                <TableCell>
                  <Skeleton className="w-52 h-6" />
                </TableCell>
                <TableCell>
                  <Skeleton className="w-8 h-8" />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    )
  }
  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px] text-center">Bookmark</TableHead>
            <TableHead>No.</TableHead>
            <TableHead>Title</TableHead>
            <TableHead>Tags</TableHead>
            <TableHead>Progress</TableHead>
            <TableHead />
          </TableRow>
        </TableHeader>
        <TableBody>
          {questions.map((question, index) => (
            <TableRow key={index}>
              <TableCell className="text-center">
                <button
                  className="transparent mt-1"
                  onClick={() =>
                    onBookmark(question.id, question.isBookmarked)
                  }>
                  {question.isBookmarked ? (
                    <FaHeart
                      className="text-red-500 cursor-pointer"
                      size={16}
                    />
                  ) : (
                    <CiHeart
                      className="text-gray-500 cursor-pointer"
                      size={20}
                    />
                  )}
                </button>
              </TableCell>
              <TableCell>{question.id}</TableCell>
              <TableCell>
                <Button
                  className="text-left cursor-pointer"
                  onClick={() => navigate('/questions/' + question.id)}
                  variant="ghost">
                  {question.title}
                </Button>
              </TableCell>
              <TableCell>
                <div className="flex items-center">
                  {question.tags.map((tag) => (
                    <span
                      className="text-xs bg-gray-200 px-2 py-1 mr-2 rounded-full text-nowrap"
                      key={tag}>
                      {tag}
                    </span>
                  ))}
                </div>
              </TableCell>
              <TableCell>
                <div className="flex items-center ">
                  <Progress
                    className="w-[100px] mr-1"
                    value={Math.max(question.progress, 0)}
                  />
                  <span>{Math.max(question.progress, 0)}%</span>
                </div>
              </TableCell>
              <TableCell className="text-right">
                <DropdownMenu>
                  <DropdownMenuTrigger asChild>
                    <Button
                      className="flex h-8 w-8 p-0 data-[state=open]:bg-muted"
                      variant="ghost">
                      <MoreHorizontal />
                      <span className="sr-only">Open menu</span>
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end" className="w-[160px]">
                    {/* TODO: 会話データの削除機能を実装する */}
                    {/* <DropdownMenuItem>会話データを削除する</DropdownMenuItem> */}
                    <DropdownMenuItem onClick={() => navigate('/contact')}>
                      問題を報告
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  )
}

export default QuestionsTable
