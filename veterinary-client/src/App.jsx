
import './App.css'

function App() {

  const handlePetitions = async() => {
    const response = await fetch('http://127.0.0.1:3000/users')
    const data = await response.json()

    console.log(data)
  }

  return (
    <div className="App">
     <h1>Hello world </h1> 
      <button onClick={handlePetitions}>Obtener data</button>
    </div>
  )
}

export default App
