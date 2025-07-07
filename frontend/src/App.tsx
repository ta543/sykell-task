import React from 'react'
import Dashboard from './pages/Dashboard'
import NoiseOverlay from './components/NoiseOverlay'
import StarField from './components/StarField'
import './App.css'

function App() {
  return (
    <>
      <NoiseOverlay />
      <div className="container">
        <h1>Web Crawler</h1>
        <Dashboard />
      </div>
    </>
  )
}

export default App
