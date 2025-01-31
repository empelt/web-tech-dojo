import { useEffect } from 'react'

import {
  useSignInWithEmailAndPassword,
  useSignInWithGoogle,
} from 'react-firebase-hooks/auth'
import { useNavigate } from 'react-router'

import { useAuth } from '@/hooks/useAuth'
import auth from '@/libs/firebase'
import { LoginForm } from '@/pages/Login/components/login-form'

const LoginPage = () => {
  const { user } = useAuth()
  const [signInWithGoogle, , googleLoading, googleError] =
    useSignInWithGoogle(auth)
  const [signInWithEmailAndPassword, , emailLoading, emailError] =
    useSignInWithEmailAndPassword(auth)
  const navigate = useNavigate()
  useEffect(() => {
    if (user) {
      console.log(user)
      navigate('/')
    }
  }, [user, navigate])
  const switchToSignup = () => {
    console.log('switchToSignup')
    navigate('/signup')
  }

  if (googleLoading || emailLoading) {
    return <div>Loading...</div>
  }

  if (googleError || emailError) {
    return <div>Error: {googleError?.message || emailError?.message}</div>
  }

  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
      <div className="w-full max-w-sm">
        <LoginForm
          signInWithEmailAndPassword={(email, password) =>
            signInWithEmailAndPassword(email, password)
          }
          signInWithGoogle={() => signInWithGoogle()}
          switchToSignup={switchToSignup}
        />
      </div>
    </div>
  )
}

export default LoginPage
