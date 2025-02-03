import { signOut } from 'firebase/auth'
import { useAuthState } from 'react-firebase-hooks/auth'
import { useNavigate } from 'react-router'

import DropdownAvatar from '../Header/DropdownAvatar'

import { Button } from '@/components/ui/button'
import { auth } from '@/lib/firebase'

export const Header = () => {
  const [user] = useAuthState(auth)
  const navigate = useNavigate()

  const logout = () => {
    signOut(auth)
  }
  return (
    <div className="flex justify-between px-8 w-screen h-16 items-center shadow-sm top-0 left-0 bg-white">
      <button className="cursor-pointer w-60" onClick={() => navigate('/')}>
        <img alt="logo" className="w-full h-auto" src="/logo.svg" />
      </button>
      {user ? (
        <DropdownAvatar
          fallback={user.displayName?.charAt(0) || ''}
          imgSrc={user.photoURL || ''}
          logout={logout}
        />
      ) : (
        <Button onClick={() => navigate('/login')}>ログイン</Button>
      )}
    </div>
  )
}
