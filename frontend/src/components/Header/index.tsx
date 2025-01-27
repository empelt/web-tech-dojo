import DropdownAvatar from '../Header/DropdownAvatar'

import { Button } from '@/components/ui/button'

export const Header = () => {
  // TODO: Implement user authentication
  const loggedIn = false

  return (
    <div className="flex justify-between px-8 w-screen h-16 items-center drop-shadow-2xl shadow-sm">
      <button
        className="font-syuku font-bold text-3xl cursor-pointer"
        onClick={() => (window.location.href = '/')}>
        WebTech 道場
      </button>
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
