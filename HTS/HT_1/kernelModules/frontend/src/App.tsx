import React, { useEffect, useState } from 'react';
import './App.css';
import {GetMemory} from "../wailsjs/go/main/App";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);

interface Memory {
  Total: number;
  Free: number;
  Used: number;
  PorcentUsed: number;
}


function App() {

  const [memory, setMemory] = useState<Memory>({Free: 0, Total: 0, Used: 0, PorcentUsed: 0});
  let data = {
    labels: ['Libre', 'En uso'],
    datasets: [
      {
        label: ' GB de memoria',
        data: [memory.Free, memory.Used],
        backgroundColor: [
          'rgba(255, 99, 132, 0.2)',
          'rgba(54, 162, 235, 0.2)'
        ],
        borderColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(54, 162, 235, 1)'
        ],
        borderWidth: 1,
      },
    ],  
  };


  useEffect(() => {
    const interval = setInterval(() => {
      GetMemory().then((result: Memory) => {
        setMemory(result);
        data.datasets[0].data = [result.Free, result.Used];
      });
    },500);

    return () => clearInterval(interval);
  }
  , []);


    return (
      <div id="App">
        <div id="container">
          <h2>RAM</h2>
          <h3>{memory.PorcentUsed}% en uso</h3>
          <Doughnut data={data} />
        </div>
      </div>
    )
}

export default App
