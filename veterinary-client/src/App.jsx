import React, {useState} from 'react';
import './App.css'

function App() {

  const [name, setName] = useState("")

  const handlePetitions = async() => {
    const response = await fetch('http://127.0.0.1:3000/users')
    const data = await response.json()

    console.log(data)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const response = await fetch('http://127.0.0.1:3000/users', {
      method: "POST",
      body: JSON.stringify({name})
    })

    const data = await response.json()
    console.log(data)

  }

  return (
    <div className="App">
     <h1>Hello world </h1> 
      <button onClick={handlePetitions}>Obtener data</button>
      <form onSubmit={handleSubmit}>
        <input type="name" placeholder="coloca tu nombre" value={name} onChange={(e) => setName(e.target.value)}/>
        <button type="submit">Guardar</button>

      </form>
    </div>
  )
}

export default App
