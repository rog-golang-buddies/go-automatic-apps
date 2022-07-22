import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import './styles/styles.scss'
import Models from './pages/Models';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
      <BrowserRouter>
      <Routes>
      <Route path="/" element={<App />} />
      <Route path="models" element={<Models />} />
    </Routes>
      </BrowserRouter>
  </React.StrictMode>
)
