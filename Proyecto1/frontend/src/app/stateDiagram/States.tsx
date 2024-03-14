"use client";
import { Network } from "vis-network";
import { useEffect, useRef, useState } from "react";
import { ProcessState } from "@/interfaces/processes.interface";

export default function States() {
  const networkContainer = useRef(null);
  const [process, setProcess] = useState<ProcessState>({
    pid: 0,
    status: "",
    previousStatus: "",
    color: "",
  });

  useEffect(() => {
    if (process.pid !== 0 && networkContainer.current) {
      const baseNodes = [
        { id: 1, label: "New", color: "white" },
        { id: 2, label: "Ready" , color: "white"},
        { id: 3, label: "Running", color: "white" },
      ];
      const baseEdges = [
        { from: 1, to: 2 , color : "white"},
        { from: 2, to: 3 , color : "white"},
      ];

      let nodes = [...baseNodes];
      let edges = [...baseEdges];

      switch (process.status) {
        case "started":
          nodes[2].color = process.color;
          break;
        case "stopped":
          nodes[1].color = process.color;
          edges.push({ from: 3, to: 2 , color : process.color});
          break;
        case "resumed":
          nodes[2].color = process.color;
          break;
        case "killed":
          nodes.push({
            id: 4,
            label: "Terminated",
            color: process.color,
          });
          edges.push({ from: 3, to: 4 , color : process.color});
          break;
      }


      const options = {

        nodes: {
          font: {
            size: 20, 
            color: "black",
          },
          color: {
            background: "white",
            border: "white",
          },
        },
        borderColor: "white",
        edges: {
          arrows: {
            to: {
              enabled: true,
              scaleFactor: 0.8,
            },
          },
        },
      };
      new Network(networkContainer.current, {nodes, edges}, options);
    } else {
      // clear
      if (networkContainer.current) {
        (networkContainer.current as HTMLElement).innerHTML = `
          <div class="text-white text-2xl font-bold mb-2 mt-4">
            No se esta corriendo ningún proceso
          </div>
        `;
      }
    }
  }, [process]); // This useEffect runs whenever `processes` changes

  async function handleNewProcess() {
    const response = await makeRequest("process/start");

    setProcess({
      pid: response.pid,
      status: "started",
      previousStatus: "started",
      color: "#8fce00",
    });
  }

  async function handleStopProcess() {
    const response = await makeRequest(`process/stop?pid=${process.pid}`);

    setProcess({
      pid: process.pid,
      status: "stopped",
      previousStatus: process.status,
      color: "#8fce00",
    });
  }

  async function handleReadyProcess() {
    const response = await makeRequest(`process/resume?pid=${process.pid}`);

    setProcess({
      pid: process.pid,
      status: "resumed",
      previousStatus: process.status,
      color: "#8fce00",
    });
  }

  async function handleKillProcess() {
    const response = await makeRequest(`process/kill?pid=${process.pid}`);

    setProcess({
      pid: process.pid,
      status: "killed",
      previousStatus: process.status,
      color: "#8fce00",
    });

    setTimeout(() => {
      setProcess({
        pid: 0,
        status: "",
        previousStatus: "",
        color: "",
      });
    }, 2000);
  }

  function makeRequest(url: string) {
    return fetch(`/api/${url}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  return (
    <div className="flex flex-col w-full justify-center items-center mt-4 gap-24">
      <div className="flex flex-row justify-around w-full items-center ">
        <p className="basis-2/4 text-2xl font-medium">PID: {process.pid}</p>
        <div className="flex flex-row justify-between w-full">
          <button
            onClick={handleNewProcess}
            className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
          >
            New
          </button>
          {process.pid !== 0 && (
            <>
              <button
                onClick={handleStopProcess}
                className="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded"
              >
                Stop
              </button>
              <button
                onClick={handleReadyProcess}
                className="bg-cyan-500 hover:bg-cyan-700 text-white font-bold py-2 px-4 rounded"
              >
                Ready
              </button>
              <button
                onClick={handleKillProcess}
                className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
              >
                Kill
              </button>
            </>
          )}
        </div>
      </div>
      <div
        className="flex w-full justify-center items-center"
        ref={networkContainer}
        style={{ height: "500px", width: "100%" }}
      ></div>
    </div>
  );
}
