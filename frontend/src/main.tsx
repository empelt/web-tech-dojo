import { StrictMode } from 'react'

import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'

import App from './App'
import { Toaster } from './components/ui/toaster'
import ChatPage from './pages/chat'
import ContactPage from './pages/contact'
import LoginPage from './pages/login'
import PrivacyPage from './pages/privacy'
import QuestionsPage from './pages/questions'
import SignupPage from './pages/signup'
import TermsPage from './pages/terms'

import './index.css'

import Layout from '@/components/Layout'

const root = document.getElementById('root')

if (!root) {
  throw new Error('Root element not found')
}

ReactDOM.createRoot(root).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route element={<Layout showFooter showHeader />}>
          <Route element={<App />} path="/" />
          <Route element={<TermsPage />} path="/terms" />
          <Route element={<PrivacyPage />} path="/privacy" />
          <Route element={<ContactPage />} path="/contact" />
        </Route>
        <Route element={<Layout loginGuard showFooter showHeader />}>
          <Route element={<QuestionsPage />} path="/questions" />
        </Route>
        <Route element={<Layout loginGuard showHeader />}>
          <Route element={<ChatPage />} path="/questions/:id" />
        </Route>
        <Route element={<LoginPage />} path="/login" />
        <Route element={<SignupPage />} path="/signup" />
      </Routes>
      <Toaster />
    </BrowserRouter>
  </StrictMode>,
)
