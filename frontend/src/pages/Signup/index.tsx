import { useEffect } from 'react'

import { useSignInWithGoogle } from 'react-firebase-hooks/auth'
import { useNavigate } from 'react-router'

import auth from '@/libs/firebase'
import { SignupForm } from '@/pages/Signup/components/signup-form'

const SignupPage = () => {
  const [signInWithGoogle, user, loading, error] = useSignInWithGoogle(auth)
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
  const signupFunction = (email: string, password: string) => {
    console.log(email, password)
  }
  if (loading) {
    return <div>Loading...</div>
  }
  if (error) {
    return <div>Error: {error.message}</div>
  }

  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
      <div className="w-full max-w-sm">
        <SignupForm
          signInWithGoogle={() => signInWithGoogle()}
          signupFunction={signupFunction}
          switchToLogin={switchToLogin}
        />
      </div>
    </div>
  )
}

export default SignupPage
