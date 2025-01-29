import { Outlet } from 'react-router'

import { Footer } from '../Footer'
import { Header } from '../Header'

type Props = {
  showFooter?: boolean
  showHeader?: boolean
}

const Layout = ({ showHeader = false, showFooter = false }: Props) => {
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
