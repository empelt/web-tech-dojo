import { useNavigate } from 'react-router'

import { SignupForm } from '@/pages/Signup/components/signup-form'
const SignupPage = () => {
  const navigate = useNavigate()
  const switchToLogin = () => {
    console.log('switchToLogin')
    navigate('/login')
  }
  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
      <div className="w-full max-w-sm">
        <SignupForm switchToLogin={switchToLogin} />
      </div>
    </div>
  )
}

export default SignupPage
