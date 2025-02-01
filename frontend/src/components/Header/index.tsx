import { useNavigate } from 'react-router'

import DropdownAvatar from '../Header/DropdownAvatar'

import { Button } from '@/components/ui/button'

export const Header = () => {
  const navigate = useNavigate()
  // TODO: Implement user authentication
  const loggedIn = false

  return (
    <div className="flex justify-between px-8 w-screen h-16 items-center shadow-sm fixed top-0 left-0 bg-white z-50">
      <button className="cursor-pointer w-60" onClick={() => navigate('/')}>
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
