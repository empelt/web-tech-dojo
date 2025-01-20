import { StrictMode } from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'
import App from './App.tsx'

const root = document.getElementById('root')!

ReactDOM.createRoot(root).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
