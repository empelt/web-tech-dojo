import { StrictMode } from 'react'

import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'

import App from './App'
import { Toaster } from './components/ui/toaster'
import ChatPage from './pages/chat'
import ContactPage from './pages/contact'
import PrivacyPage from './pages/privacy'
import QuestionsPage from './pages/questions'
import TermsPage from './pages/terms'

import './index.css'
import { AuthProvider } from './hooks/useAuth'
import LoginPage from './pages/Login'
import SignupPage from './pages/Signup'
import Layout from '@/components/Layout'

const root = document.getElementById('root')

if (!root) {
  throw new Error('Root element not found')
}

ReactDOM.createRoot(root).render(
  <StrictMode>
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          <Route element={<Layout showFooter showHeader />}>
            <Route element={<App />} path="/" />
            <Route element={<QuestionsPage />} path="/questions" />
            <Route element={<App />} path="/" />
            <Route element={<TermsPage />} path="/terms" />
            <Route element={<PrivacyPage />} path="/privacy" />
            <Route element={<ContactPage />} path="/contact" />
          </Route>
          <Route element={<Layout showHeader />}>
            <Route element={<ChatPage />} path="/questions/:id" />
          </Route>
          <Route element={<LoginPage />} path="/login" />
          <Route element={<SignupPage />} path="/signup" />
        </Routes>
        <Toaster />
      </AuthProvider>
    </BrowserRouter>
  </StrictMode>,
)
