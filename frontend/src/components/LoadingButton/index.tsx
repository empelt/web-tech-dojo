import { VariantProps } from 'class-variance-authority'

import { Button, buttonVariants } from '@/components/ui/button'

type Props = React.ButtonHTMLAttributes<HTMLButtonElement> &
  VariantProps<typeof buttonVariants>

const LoadingButton = ({ className, variant }: Props) => {
  return (
    <Button className={`p-0 h-9 w-20 ${className}`} disabled variant={variant}>
      <div className="animate-spin h-4 w-4 border-2 border-t-transparent border-blue-500 border-solid rounded-full" />{' '}
    </Button>
  )
}

export default LoadingButton
