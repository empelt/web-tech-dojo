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
    <div className="flex justify-between px-8 w-screen h-16 items-center drop-shadow-2xl shadow-sm">
      <h1 className="font-syuku font-bold text-3xl">WebTech 道場</h1>
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
