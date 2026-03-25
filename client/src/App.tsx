import { useEffect } from 'react'

function App() {
  useEffect(() => {
    fetch('/api/health')
      .then((res) => res.json())
      .then((data) => console.log('Health check:', data))
      .catch((err) => console.error('Health check failed:', err))
  }, [])

  return <h1>GozuTab</h1>
}

export default App
