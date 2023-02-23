import React, {useState, useEffect} from 'react';
import { ChakraProvider } from '@chakra-ui/react'
import Home from './Home.jsx'
import './App.css'

function App() {

  const [name, setName] = useState("")
  const [users, setUsers] = useState([])

  const handlePetitions = async() => {
    const response = await fetch('http://127.0.0.1:3000/users')
    const data = await response.json()

    console.log(data)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const response = await fetch('http://127.0.0.1:3000/users', {
      method: "POST",
      body: JSON.stringify({name}),
      headers: {
        "Content-Type": "form-data"
      }
    })

    const data = await response.json()
    console.log(data)

  }

  useEffect(() => {
    async function loadUsers() {
      const response = await fetch('http://127.0.0.1:3000/api/users')
      const data = await response.json()

      setUsers(data)
    }
    loadUsers()
  }, [])

  return (
    <ChakraProvider>
    <Home handlePetitions={handlePetitions} handleSubmit={handleSubmit} name={name}/>
    </ChakraProvider>
  )
}

export default App
