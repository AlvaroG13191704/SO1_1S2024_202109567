'use client';

import { Process, Processes } from "@/interfaces/processes.interface";
import { useEffect, useState } from "react";

export function Tree() {
  const [selectedProcess, setSelectedProcess] = useState<Process>({
    pid: 0,
    name: "",
    user: 0,
    children: []
  });
  const [processes, setProcesses] = useState<Processes>({
    processes: []
  });

  useEffect(() => {
    fetch('http://127.0.0.1:8080/get-processes')
      .then(response => response.json())
      .then(data => {
        setProcesses(data);
        console.log({processes});
      })
      .catch(error => console.error('Error:', error));
  }, []); // Empty array ensures this runs only on mount

  const handleSelectChange = (event: any) => {
    setSelectedProcess(event.target.value);

    console.log({selectedProcess});
  };
  return (
<div className="flex flex-col w-full">
  <h1 className="text-white text-2xl font-bold mb-2">Seleccione un proceso</h1>
  <select 
    value={selectedProcess.pid} 
    onChange={handleSelectChange}
    className="w-64 h-10 pl-3 pr-6 text-base placeholder-gray-600 border rounded-lg appearance-none focus:shadow-outline"
  >
    {processes.processes.map((process : Process, index: number) => (
      <option key={index} value={process.name}>
        {process.pid}
      </option>
    ))}
  </select>
</div>
  );
}
