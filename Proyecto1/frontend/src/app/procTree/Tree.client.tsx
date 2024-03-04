'use client';
import { DataSet, Network } from 'vis-network/standalone/esm/vis-network';
import { Process, Processes } from "@/interfaces/processes.interface";
import { useEffect, useRef, useState } from "react";

export function Tree() {
  const networkContainer = useRef(null);
  const [selectedProcess, setSelectedProcess] = useState<Process>({
    pid: 0,
    name: "",
    user: 0,
    children: []
  });

  const [processes, setProcesses] = useState({ processes: [] });

  useEffect(() => {
    fetch('http://127.0.0.1:8080/get-processes')
      .then(response => response.json())
      .then(data => {
        setProcesses(data);
      })
      .catch(error => console.error('Error:', error));
  }, []); // Empty array ensures this runs only on mount
  
  // useEffect(() => {
  //   // console.log(processes);
  // }, [processes]); // This useEffect runs whenever `processes` changes

  useEffect(() => {
    if (selectedProcess.children.length > 0 && networkContainer.current) {
      const nodes = new DataSet(selectedProcess.children.map((child, index) => ({ id: index, label: child.name + " (" + child.pid + ")"})));
      const edges = new DataSet(selectedProcess.children.map((child, index) => ({ from: selectedProcess.pid, to: index })));

      const data = {
        nodes: nodes,
        edges: edges
      };

      const options = {
        layout: {
          hierarchical: {
            direction: "UD",
            sortMethod: "directed",
            levelSeparation: 500,
            nodeSpacing: 500
          }
        }
      };

      new Network(networkContainer.current, data, options);
    }
  }, [selectedProcess]); // This useEffect runs whenever `selectedProcess` changes

  
  const handleSelectChange = (event:any) => {
    console.log(event.target.value);
    const selectedPid = parseInt(event.target.value);
    processes.processes.forEach((process: Process) => {
      if (process.pid === selectedPid) {
        console.log(process);
        setSelectedProcess(process);
      }
    });
  };
  return (
<div className="flex flex-col w-full">
  <h1 className="text-white text-2xl font-bold mb-2">Seleccione un proceso</h1>
  <select 
    value={selectedProcess.pid} 
    onChange={handleSelectChange}
    className="w-64 h-10 pl-3 pr-6 text-base text-black placeholder-gray-950 border rounded-lg appearance-none focus:shadow-outline"
  >
  {processes.processes && processes.processes.map((process : Process, index: number) => (
    <option key={index} value={process.pid} className="text-black">
      {process.pid}
    </option>
  ))}
  </select>
  <div ref={networkContainer} style={{ height: "500px", width: "100%" }}></div>
</div>
  );
}
