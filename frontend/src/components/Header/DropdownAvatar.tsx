import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar'
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from '@/components/ui/dropdown-menu'

type Props = {
  fallback: string
  imgSrc: string
  logout: () => void
}

const DropdownAvatar = ({ imgSrc, fallback, logout }: Props) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Avatar className="cursor-pointer">
          <AvatarImage
            alt="avatar image"
            src={imgSrc ?? 'https://github.com/shadcn.png'}
          />
          <AvatarFallback>{fallback}</AvatarFallback>
          <span className="sr-only">Toggle user menu</span>
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        {/* <DropdownMenuItem>DashBoard</DropdownMenuItem> */}
        {/* <DropdownMenuItem>Settings</DropdownMenuItem> */}
        {/* <DropdownMenuSeparator /> */}
        <DropdownMenuItem onClick={logout}>Logout</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

export default DropdownAvatar
