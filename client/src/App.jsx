import { useState } from 'react'
import './App.css'


function App() {

  const [name, setName] = useState('')


  const handleSubmit = async (e) => {
    e.preventDefault()
    const response = await fetch(import.meta.env.VITE_API + '/users', {
      method: 'POST',
      body: JSON.stringify({name}),
      headers: {
        "Content-Type": "application/json"
      }
    })
    const data = await response.json()
    console.log(data);
    
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input 
        type="name" 
        placeholder='Set your name' 
        onChange={(e) => setName(e.target.value)} />
        <button>Save</button>
      </form>
    </div>
  )
}

export default App
