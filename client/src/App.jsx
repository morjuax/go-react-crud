import { useEffect, useState } from 'react'
import './App.css'


function App() {

  const [name, setName] = useState('')
  const [users, setUsers] = useState([])

  async function loadUsers() {
    const response = await fetch(import.meta.env.VITE_API + '/users')
    const data = await response.json()
    console.log(data)
    setUsers(data.users);
   }

  useEffect(() => {
   loadUsers()
  }, [])


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
    loadUsers()
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
      <ul>
        {users.map(user => (
          <li key={user._id}>{user.name}</li>
        ))}
      </ul>
    </div>
  )
}

export default App
