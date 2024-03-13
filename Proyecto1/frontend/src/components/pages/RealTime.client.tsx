'use client';


import { RealTime } from "@/interfaces/ram.interface";
import { useEffect, useState } from "react";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);


export default function RealTimePage() {

  const [realTime, setRealTime] = useState<RealTime>({
    cpu: {
      percentage: ""
    },
    ram: {
      totalRam: "",
      freeRam: "",
      usedRam: "",
      percent: "",
      date: ""
    }
  })

  let dataRAM = {
    labels: ['En uso', 'Libre'],
    datasets: [
      {
        label: 'GB',
        data: [realTime.ram.usedRam, realTime.ram.freeRam],
        backgroundColor: [
          'rgba(255, 99, 132, 0.2)',
          'rgba(255, 159, 64, 0.2)',
        ],
        borderColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(255, 159, 64, 1)',
        ],
        borderWidth: 1,
      },
    ],
  };

  let dataCPU = {
    labels: ['En uso', 'Libre'],
    datasets: [
      {
        label: '% de uso',
        data: [Number(realTime.cpu.percentage), (100 - Number(realTime.cpu.percentage))],
        backgroundColor: [
          'rgba(255, 99, 132, 0.2)',
          'rgba(255, 159, 64, 0.2)',
        ],
        borderColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(255, 159, 64, 1)',
        ],
        borderWidth: 1,
      },
    ],
  };

  useEffect(() => {
    const interval = setInterval(() => {
      fetch('http://127.0.0.1:8080/api/real-time')
      .then(response => response.json())
      .then(data => {
        setRealTime(data);
        console.log({data});

        dataRAM.datasets[0].data = [data.ram.usedRam, data.ram.freeRam];
        dataCPU.datasets[0].data = [Number(data.cpu.percentage), (100 - Number(data.cpu.percentage))];
      })
      .catch(error => console.error('Error:', error));
  }, 500);

    return () => clearInterval(interval);
  }
  , []);

  return (
    <main className="flex  flex-col items-center justify-start p-6 gap-8">
      <h1 className="text-white text-2xl font-bold">Monitoreo en Tiempo Real</h1>
      <div className="flex flex-row justify-evenly w-full">

        <div>
          <h2 className="text-white text-lg font-bold">Memoria RAM</h2>
          <p className="text-white text-sm">Porcentaje de uso actual: {realTime.ram.percent}%</p>
          <Doughnut data={dataRAM} width={400} />
        </div>
        <div>
          <h2 className="text-white text-lg font-bold">CPU</h2>
          <p className="text-white text-sm">Porcentaje de uso actual: {realTime.cpu.percentage}%</p>
          <Doughnut data={dataCPU} width={400} />
        </div>
      </div>
    </main>
  );
}
