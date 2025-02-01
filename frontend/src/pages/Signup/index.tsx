import { useEffect } from 'react'

import {
  useAuthState,
  useCreateUserWithEmailAndPassword,
  useSignInWithGoogle,
} from 'react-firebase-hooks/auth'
import { useNavigate } from 'react-router'

import { auth } from '@/libs/firebase'
import { SignupForm } from '@/pages/signup/components/signup-form'

const SignupPage = () => {
  const [user, loading, error] = useAuthState(auth)
  const [signInWithGoogle] = useSignInWithGoogle(auth)
  const [createUserWithEmailAndPassword] =
    useCreateUserWithEmailAndPassword(auth)
  const navigate = useNavigate()

  useEffect(() => {
    if (user) {
      navigate('/')
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
        <SignupForm
          createUserWithEmailAndPassword={(email, password) =>
            createUserWithEmailAndPassword(email, password)
          }
          signInWithGoogle={() => signInWithGoogle()}
        />
      </div>
    </div>
  )
}

export default SignupPage
