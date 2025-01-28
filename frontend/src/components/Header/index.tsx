import DropdownAvatar from '../Header/DropdownAvatar'

import { Button } from '@/components/ui/button'

export const Header = () => {
  // TODO: Implement user authentication
  const loggedIn = false

  return (
    <div className="flex justify-between px-8 w-screen h-16 items-center drop-shadow-2xl shadow-sm">
      <button
        className="cursor-pointer w-60"
        onClick={() => (window.location.href = '/')}>
        <img alt="logo" className="w-full h-auto" src="/logo.svg" />
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
