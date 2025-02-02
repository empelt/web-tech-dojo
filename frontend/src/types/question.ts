import { CheckCircle, Circle, CircleOff, HelpCircle, Timer } from 'lucide-react'

export type Question = {
  content: string
  id: string
  isBookmarked: boolean
  progress: number
  tags: string[]
  title: string
}

export const tags = [
  {
    value: 'network',
    label: 'Network',
    icon: HelpCircle,
  },
  {
    value: 'security',
    label: 'Security',
    icon: Circle,
  },
  {
    value: 'database',
    label: 'Database',
    icon: Timer,
  },
  {
    value: 'frontend',
    label: 'Frontend',
    icon: CheckCircle,
  },
  {
    value: 'backend',
    label: 'Backend',
    icon: CircleOff,
  },
  {
    value: 'react',
    label: 'React',
    icon: CircleOff,
  },
  {
    value: 'hook',
    label: 'Hook',
    icon: CircleOff,
  },
]

export const bookmarkStatus = [
  {
    value: 'isBookmarked',
    label: 'Is Bookmarked',
    icon: HelpCircle,
  },
  {
    value: 'isNotBookmarked',
    label: 'Is Not Bookmarked',
    icon: Circle,
  },
]

export const progressStatus = [
  {
    value: 'todo',
    label: 'To Do',
    icon: HelpCircle,
  },
  {
    value: 'inProgress',
    label: 'In Progress',
    icon: Circle,
  },
  {
    value: 'completed',
    label: 'Completed',
    icon: CheckCircle,
  },
]
