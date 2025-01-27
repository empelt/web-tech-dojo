import { MoreHorizontal } from 'lucide-react'
import { CiHeart } from 'react-icons/ci'
import { FaHeart } from 'react-icons/fa'

import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Progress } from '@/components/ui/progress'
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
  questions: Question[]
}

const QuestionsTable = ({ questions }: Props) => {
  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Bookmark</TableHead>
            <TableHead>No.</TableHead>
            <TableHead>Title</TableHead>
            <TableHead>Tags</TableHead>
            <TableHead>Progress</TableHead>
            <TableHead />
          </TableRow>
        </TableHeader>
        <TableBody>
          {questions.map((question) => (
            <TableRow key={question.id}>
              <TableCell>
                {question.isBookmarked ? (
                  <FaHeart
                    className="ml-2 text-red-500 cursor-pointer"
                    type="button"
                  />
                ) : (
                  <CiHeart
                    className="ml-2 text-gray-500 cursor-pointer"
                    type="button"
                  />
                )}
              </TableCell>
              <TableCell>{question.id}</TableCell>
              <TableCell>{question.title}</TableCell>
              <TableCell>
                <div className="flex w-[100px] items-center">
                  {question.tags.map((tag) => (
                    <span
                      className="text-xs bg-gray-200 px-2 py-1 mr-2 rounded-full"
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
                    value={question.progress}
                  />
                  <span>{question.progress}%</span>
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
                    <DropdownMenuItem>Delete Data</DropdownMenuItem>
                    <DropdownMenuItem>問題を報告</DropdownMenuItem>
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
