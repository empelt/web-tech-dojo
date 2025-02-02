import { useEffect } from 'react'

import { useAuthState } from 'react-firebase-hooks/auth'
import { useNavigate } from 'react-router'

import { auth } from '@/libs/firebase'
import LoginForm from '@/pages/login/components/form'

const LoginPage = () => {
  const [user] = useAuthState(auth)
  const navigate = useNavigate()

  useEffect(() => {
    if (user) {
      navigate('/questions')
    }
  }, [user, navigate])

  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
      <div className="w-full max-w-sm">
        <LoginForm />
      </div>
    </div>
  )
}

export default LoginPage
