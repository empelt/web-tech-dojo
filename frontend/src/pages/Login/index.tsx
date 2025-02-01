import { useEffect } from 'react'

import {
  useAuthState,
  useSignInWithEmailAndPassword,
  useSignInWithGoogle,
} from 'react-firebase-hooks/auth'
import { useNavigate } from 'react-router'

import { auth } from '@/libs/firebase'
import { LoginForm } from '@/pages/login/components/login-form'

const LoginPage = () => {
  const [user, loading, error] = useAuthState(auth)
  const [signInWithGoogle] = useSignInWithGoogle(auth)
  const [signInWithEmailAndPassword] = useSignInWithEmailAndPassword(auth)
  const navigate = useNavigate()

  useEffect(() => {
    if (user) {
      navigate('/questions')
    }
  }, [user, navigate])

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error: {error?.message}</div>
  }

  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
      <div className="w-full max-w-sm">
        <LoginForm
          signInWithEmailAndPassword={(email, password) =>
            signInWithEmailAndPassword(email, password)
          }
          signInWithGoogle={() => signInWithGoogle()}
        />
      </div>
    </div>
  )
}

export default LoginPage
