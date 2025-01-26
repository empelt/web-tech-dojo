import DropdownAvatar from '../Header/DropdownAvatar'

import { Button } from '@/components/ui/button'

export const Header = () => {
  // TODO: Implement user authentication
  const loggedIn = false

  return (
    <div className="flex justify-between px-8 w-screen h-16 items-center drop-shadow-2xl shadow-sm">
      <h1 className="font-syuku font-bold text-3xl">WebTech 道場</h1>
      {loggedIn ? (
        // TODO: Implement user authentication
        <DropdownAvatar
          fallback="name"
          imgSrc="https://github.com/shadcn.png"
        />
      ) : (
        <Button>ログイン</Button>
      )}
    </div>
  )
}
