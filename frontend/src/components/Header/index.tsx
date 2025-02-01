import { useNavigate } from 'react-router'

import DropdownAvatar from '../Header/DropdownAvatar'

import { Button } from '@/components/ui/button'
import { useAuth } from '@/hooks/useAuth'

export const Header = () => {
  const { user } = useAuth()
  const navigate = useNavigate()

  const switchToLoginPage = () => {
    navigate('/login')
  }
  return (
    <div className="flex justify-between px-8 w-screen h-16 items-center shadow-sm fixed top-0 left-0 bg-white z-50">
      <button className="cursor-pointer w-60" onClick={() => navigate('/')}>
        <img alt="logo" className="w-full h-auto" src="/logo.svg" />
      </button>
      {user ? (
        <DropdownAvatar
          fallback="name"
          imgSrc="https://github.com/shadcn.png"
        />
      ) : (
        <Button onClick={switchToLoginPage}>ログイン</Button>
      )}
    </div>
  )
}
