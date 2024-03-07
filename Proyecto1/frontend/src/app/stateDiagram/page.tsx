import States from "./States";

export default function StateDiagramPage() {
  return (
    <main className="flex  flex-col items-center justify-start p-6">
      <h1 className="text-white text-2xl font-bold">Diagrama de estados</h1>
      <States />
    </main>
  );
}
