import { StrictMode } from 'react'

import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'

import App from './App'
import { Header } from './components/Header'
import Questions from './pages/questions'
import './index.css'

const root = document.getElementById('root')

if (!root) {
  throw new Error('Root element not found')
}

ReactDOM.createRoot(root).render(
  <StrictMode>
    <BrowserRouter>
      <Header />
      <Routes>
        <Route element={<App />} path="/" />
        <Route element={<Questions />} path="/questions" />
        <Route element={<App />} path="/" />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
