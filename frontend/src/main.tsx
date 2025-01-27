import { StrictMode } from 'react'

import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'

import App from './App'
import { Footer } from './components/Footer'
import { Header } from './components/Header'
import { Toaster } from './components/ui/toaster'
import ContactPage from './pages/contact'
import PrivacyPage from './pages/privacy'
import QuestionsPage from './pages/questions'
import TermsPage from './pages/terms'
import './index.css'

const root = document.getElementById('root')

if (!root) {
  throw new Error('Root element not found')
}

ReactDOM.createRoot(root).render(
  <StrictMode>
    <BrowserRouter>
      <div className="flex flex-col h-screen justify-between">
        <div className="mb-4">
          <Header />
          <Routes>
            <Route element={<App />} path="/" />
            <Route element={<QuestionsPage />} path="/questions" />
            <Route element={<App />} path="/" />
            <Route element={<TermsPage />} path="/terms" />
            <Route element={<PrivacyPage />} path="/privacy" />
            <Route element={<ContactPage />} path="/contact" />
          </Routes>
          <Toaster />
        </div>
        <Footer />
      </div>
    </BrowserRouter>
  </StrictMode>,
)
