'use client'

import { Table } from '@tanstack/react-table'
import { X } from 'lucide-react'

import { DataTableFacetedFilter } from './data-table-faceted-filter'
import { priorities, statuses } from '../data/data'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'

interface DataTableToolbarProps<TData> {
  table: Table<TData>
}

export const DataTableToolbar = <TData,>({
  table,
}: DataTableToolbarProps<TData>) => {
  const isFiltered = table.getState().columnFilters.length > 0

  return (
    <div className="flex items-center justify-between">
      <div className="flex flex-1 items-center space-x-2">
        <Input
          className="h-8 w-[150px] lg:w-[250px]"
          onChange={(event) =>
            table.getColumn('title')?.setFilterValue(event.target.value)
          }
          placeholder="Filter tasks..."
          value={(table.getColumn('title')?.getFilterValue() as string) ?? ''}
        />
        {table.getColumn('status') && (
          <DataTableFacetedFilter
            column={table.getColumn('status')}
            options={statuses}
            title="Status"
          />
        )}
        {table.getColumn('priority') && (
          <DataTableFacetedFilter
            column={table.getColumn('priority')}
            options={priorities}
            title="Priority"
          />
        )}
        {isFiltered && (
          <Button
            className="h-8 px-2 lg:px-3"
            onClick={() => table.resetColumnFilters()}
            variant="ghost">
            Reset
            <X />
          </Button>
        )}
      </div>
    </div>
  )
}
