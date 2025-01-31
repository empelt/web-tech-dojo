import { StrictMode } from 'react'

import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'

import App from './App'
import { Header } from './components/Header'
import './index.css'
import { AuthProvider } from './hooks/useAuth'
import LoginPage from './pages/Login'
import SignupPage from './pages/Signup'

const root = document.getElementById('root')

if (!root) {
  throw new Error('Root element not found')
}

ReactDOM.createRoot(root).render(
  <StrictMode>
    <BrowserRouter>
      <AuthProvider>
        <Header />
        <Routes>
          <Route element={<LoginPage />} path="/login" />
          <Route element={<SignupPage />} path="/signup" />
          <Route element={<App />} path="/" />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  </StrictMode>,
)
