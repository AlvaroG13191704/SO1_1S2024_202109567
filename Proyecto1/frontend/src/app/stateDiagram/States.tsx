"use client";
import { DataSet, Network } from "vis-network/standalone/esm/vis-network";
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
        { id: 1, label: "New", color: "#FFFFFF", shape: "circle", },
        { id: 2, label: "Ready", color: "#FFFFFF" ,shape: "circle"},
        { id: 3, label: "Running", color: "#FFFFFF", shape: "circle" },
      ];
      const baseEdges = [
        { id: 1, from: 1, to: 2, },
        { id: 2, from: 2, to: 3 },
      ];

      let nodes = new DataSet([...baseNodes]);
      let edges = new DataSet([...baseEdges]);

      switch (process.status) {
        case "started":
          nodes.update({ id: 3, color: process.color });
          break;
        case "stopped":
          nodes.update({ id: 2, color: process.color });
          edges.add({ id: 4, from: 3, to: 2 });
          edges.add({ id: 5, from: 2, to: 3, });
          break;
        case "resumed":
          nodes.update({ id: 3, color: process.color });
          break;
        case "killed":
          nodes.add({ id: 4, label: "Killed", color: process.color ,shape: "circle"});
          edges.add({ id: 3, from: 3, to: 4 });
          break;
      }

      const data = {
        nodes: nodes,
        edges: edges,
      };

      const options = {
        layout: {
          hierarchical: {
            direction: "LR",
            levelSeparation: 300,
            nodeSpacing: 300,
            blockShifting: true,
          },
        },
      };

      new Network(networkContainer.current, data, options);
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
    return fetch(`http://127.0.0.1:8080/api/${url}`, {
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
