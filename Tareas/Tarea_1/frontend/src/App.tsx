import { useState } from 'react'
import './App.css'


interface MyData {
  name: string
  uid: string
}

function App() {
  const [data, setData] = useState<MyData>({
    name:"",
    uid:""
  })

  async function request() {
    // make a request to the api to get the values
    const response = await fetch('http://127.0.0.1:8080/data');
    const data = await response.json();
    setData(data);
  }

  return (
    <>
      <div>
        <h1>Tarea 1 - SO1 - 1s2024</h1>
      </div>
      <div className="card">
        <button onClick={request}>
          Mostrar datos
        </button>
      </div>
      <p className="read-the-docs">
        {data.name} {data.uid}
      </p>
    </>
  )
}

export default App
