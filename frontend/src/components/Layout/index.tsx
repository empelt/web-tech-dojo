import { useAuthState } from 'react-firebase-hooks/auth'
import { Navigate, Outlet } from 'react-router'

import { Footer } from '../Footer'
import { Header } from '../Header'
import LoadingScreen from '../LoadingScreen'

import { auth } from '@/libs/firebase'

type Props = {
  loginGuard?: boolean
  showFooter?: boolean
  showHeader?: boolean
}

const Layout = ({
  showHeader = false,
  showFooter = false,
  loginGuard = false,
}: Props) => {
  const [user, loading] = useAuthState(auth)

  if (loading) {
    return <LoadingScreen />
  }

  if (loginGuard) {
    if (!user) {
      return <Navigate to="/login" />
    }
  }

  return (
    <div className="flex flex-col justify-between min-h-screen">
      <div>
        {showHeader && <Header />}
        <div className={showHeader ? 'mt-16' : ''}>
          <Outlet />
        </div>
      </div>
      {showFooter && <Footer />}
    </div>
  )
}

export default Layout
