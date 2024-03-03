'use client'
'use client'
import { RAM } from '@/interfaces/ram.interface';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);


export function DoughnutGraph({freeRam, usedRam} : RAM) {

let data = {
  labels: ['En uso', 'Libre'],
  datasets: [
    {
      label: '% de uso',
      data: [usedRam, freeRam],
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

  return <Doughnut data={data} width={400} />;

}

