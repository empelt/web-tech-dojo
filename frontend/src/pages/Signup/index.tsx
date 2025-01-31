import { useEffect } from 'react'

import {
  useCreateUserWithEmailAndPassword,
  useSignInWithGoogle,
} from 'react-firebase-hooks/auth'
import { useNavigate } from 'react-router'

import { useAuth } from '@/hooks/useAuth'
import auth from '@/libs/firebase'
import { SignupForm } from '@/pages/Signup/components/signup-form'

const SignupPage = () => {
  const { user } = useAuth()
  const [signInWithGoogle, , googleLoading, googleError] =
    useSignInWithGoogle(auth)
  const [createUserWithEmailAndPassword, , emailLoading, emailError] =
    useCreateUserWithEmailAndPassword(auth)
  const navigate = useNavigate()
  useEffect(() => {
    if (user) {
      console.log(user)
      navigate('/')
    }
  }, [user, navigate])
  const switchToLogin = () => {
    console.log('switchToLogin')
    navigate('/login')
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
        <SignupForm
          createUserWithEmailAndPassword={(email, password) =>
            createUserWithEmailAndPassword(email, password)
          }
          signInWithGoogle={() => signInWithGoogle()}
          switchToLogin={switchToLogin}
        />
      </div>
    </div>
  )
}

export default SignupPage
