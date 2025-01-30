import { useNavigate } from 'react-router'

import { LoginForm } from '@/pages/Login/components/login-form'
const LoginPage = () => {
  const navigate = useNavigate()
  const switchToSignup = () => {
    console.log('switchToSignup')
    navigate('/signup')
  }
  const loginWithGoogle = () => {
    console.log('loginWithGoogle')
  }
  const loginFunction = (email: string, password: string) => {
    console.log(email, password)
  }
  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
      <div className="w-full max-w-sm">
        <LoginForm
          loginFunction={loginFunction}
          loginWithGoogle={loginWithGoogle}
          switchToSignup={switchToSignup}
        />
      </div>
    </div>
  )
}

export default LoginPage
