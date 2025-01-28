import { Outlet } from 'react-router'

import { Footer } from '../Footer'
import { Header } from '../Header'

type Props = {
  showFooter?: boolean
  showHeader?: boolean
}

const Layout = ({ showHeader = false, showFooter = false }: Props) => {
  return (
    <>
      {showHeader && <Header />}
      <div className={showHeader ? 'mt-16' : ''}>
        <Outlet />
      </div>
      {showFooter && <Footer />}
    </>
  )
}

export default Layout
