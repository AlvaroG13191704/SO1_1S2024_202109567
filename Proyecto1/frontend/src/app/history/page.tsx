import LineGraph from "@/components/graphs/Line";

export default function HistoryPage() {
  return (
    <main className="flex  flex-col items-center justify-start p-6">
      <h1 className="text-white text-2xl font-bold">Monitoreo Histórico</h1>
      <div className="flex flex-col justify-evenly w-full">

        <div>
          <h2 className="text-white text-lg font-bold">Memoria RAM</h2>
          <LineGraph />
        </div>
        <div>
          <h2 className="text-white text-lg font-bold">CPU</h2>
          <LineGraph />
        </div>
      </div>
    </main>
  );
}
