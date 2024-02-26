import { DoughnutGraph } from "@/components/graphs/Doughnut";

export default function Home() {
  return (
    <main className="flex  flex-col items-center justify-start p-6 gap-8">
      <h1 className="text-white text-2xl font-bold">Monitoreo en Tiempo Real</h1>
      <div className="flex flex-row justify-evenly w-full">

        <div>
          <h2 className="text-white text-lg font-bold">Memoria RAM</h2>
          <p className="text-white text-sm">Porcentaje de uso actual: 25%</p>
          <DoughnutGraph />
        </div>
        <div>
          <h2 className="text-white text-lg font-bold">CPU</h2>
          <p className="text-white text-sm">Porcentaje de uso actual: 25%</p>
          <DoughnutGraph />
        </div>
      </div>
    </main>
  );
}
