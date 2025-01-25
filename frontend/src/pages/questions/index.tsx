import { columns } from "./components/columns"
import { DataTable } from "./components/data-table"

const tasks = [
  { id: '1', title: 'Task 1', status: 'Description for Task 1', label: 'documentation', priority: 'medium' },
  { id: '2', title: 'Task 2', status: 'Description for Task 2', label: 'bug', priority: 'high' },
  { id: '3', title: 'Task 3', status: 'Description for Task 3', label: 'feature', priority: 'low' },
];

export default function Questions() {
  return (
    <>
      <DataTable data={tasks} columns={columns} />
    </>
  )
}