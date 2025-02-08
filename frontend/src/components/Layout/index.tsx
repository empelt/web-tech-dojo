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
    <div className={ showFooter ? 'h-[100vh] grid grid-rows-[1fr_auto]' : 'h-[100vh]' }>
      <div className= { showHeader ? 'h-full mx-auto grid grid-rows-[auto_1fr] overflow-hidden' : '' }>
        {showHeader && <Header />}
        <div className='overflow-scroll hidden-scrollbar pb-4'>
          <Outlet />
        </div>
      </div>
      {showFooter && <Footer />}
    </div>
  )
}

export default Layout
