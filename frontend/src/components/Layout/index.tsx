import { useAuthState } from 'react-firebase-hooks/auth'
import { Navigate, Outlet } from 'react-router'

import { Footer } from '../Footer'
import { Header } from '../Header'
import LoadingScreen from '../LoadingScreen'

import { auth } from '@/lib/firebase'

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
    <div>
      <div className= { showHeader ? 'mx-auto h-[100vh] grid grid-rows-[auto_1fr]' : '' }>
        {showHeader && <Header />}
        <Outlet />
      </div>
      {showFooter && <Footer />}
    </div>
  )
}

export default Layout
