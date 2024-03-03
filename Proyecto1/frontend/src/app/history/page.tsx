import LineGraph from "@/components/graphs/Line";
import History from "./History.client";

export default function HistoryPage() {
  return (
    <main className="flex  flex-col items-center justify-start p-6">
      <h1 className="text-white text-2xl font-bold">Monitoreo Histórico</h1>
      <div className="flex flex-col justify-evenly w-full">
      <History />
      </div>
    </main>
  );
}
