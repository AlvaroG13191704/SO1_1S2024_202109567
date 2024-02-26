import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Link from "next/link";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Proyecto 1 Sistemas Operativos 1 ",
  description: "Proyecto 1 Sistemas Operativos 1 que muestra información de la CPU, Memoria y Procesos de una vm",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="px-10 py-6 flex flex-col justify-center items-center gap-6">
          <nav className="p-4 w-full">
            <div className="flex flex-row justify-around">
              <h1 className="text-white text-2xl font-bold basis-1/4">SO1</h1>
              <h1 className="text-white text-2xl font-bold">Proyecto 1 Modulos de Kernel </h1>
              <h1 className="text-white text-2xl font-bold">Alvaro García - 202109567</h1>
            </div>
          </nav>

          <div className="flex flex-row w-full justify-around ">
            <Link href="/" className="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded">Monitoreo en Tiempo Real</Link>
            <Link href="/history" className="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded">Monitoreo Histórico</Link>
            <Link href="/procTree" className="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded">Árbol de procesos</Link>
            <Link href="/stateDiagram" className="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded">Diagrama de estados</Link>
          </div>
          <div className="h-full w-full">
            {children}
          </div>
        </div>
      </body>
    </html>
  );
}
