'use client';
import { useEffect, useState } from "react";
import { RAMHistory, CPUHistory, History } from '../../interfaces/history.interface';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import { Line } from 'react-chartjs-2';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

export default function History() {
  const [labelsRAM, setLabelsRAM] = useState([]);
  const [labelsCPU, setLabelsCPU] = useState([]);
  const optionsRAM = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top' as const,
      },
      title: {
        display: true,
        text: 'Historial de Memoria RAM',
      },
    },
  };

  const optionsCPU = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top' as const,
      },
      title: {
        display: true,
        text: 'Historial del CPU',
      },
    },
  };
  
  const [dataRAM, setDataRAM] = useState({
    labels: labelsRAM,
    datasets: [
      {
        label: 'Porcentaje de memoria RAM usada',
        data: [],
        borderColor: 'rgb(255, 99, 132)',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
      } 
    ],
  });

  const [dataCPU, setDataCPU] = useState({
    labels: labelsCPU,
    datasets: [
      {
        label: 'Porcentaje de CPU usada',
        data: [],
        borderColor: 'rgb(53, 162, 235)',
        backgroundColor: 'rgba(53, 162, 235, 0.5)',
      }
    ],
  });

  useEffect(() => {
    fetch('http://127.0.0.1:8080/api/get-history')
      .then(response => response.json())
      .then(data => {
        console.log(data)
        const newLabelsRAM = data.ram.map((ram: RAMHistory) => ram.date);
        const newDataRAM = data.ram.map((ram: RAMHistory) => ram.percent);
        setLabelsRAM(newLabelsRAM);
        setDataRAM(prevData => ({
          ...prevData,
          labels: newLabelsRAM,
          datasets: [
            {
              ...prevData.datasets[0],
              data: newDataRAM
            }
          ]
        }));

        const newLabelsCPU = data.cpu.map((cpu: CPUHistory) => cpu.date);
        const newDataCPU = data.cpu.map((cpu: CPUHistory) => cpu.percentage);
        setLabelsCPU(newLabelsCPU);
        setDataCPU(prevData => ({
          ...prevData,
          labels: newLabelsCPU,
          datasets: [
            {
              ...prevData.datasets[0],
              data: newDataCPU
            }
          ]
        }));
      })
      .catch(error => console.error('Error:', error));
  }, []); // Empty array ensures this runs only on mount
  
  return (
    <>
      <div>
        <Line options={optionsRAM} data={dataRAM} width={900} />
      </div>
      <div>
        <Line options={optionsCPU} data={dataCPU} width={900} />
      </div>
    </>
  );
}