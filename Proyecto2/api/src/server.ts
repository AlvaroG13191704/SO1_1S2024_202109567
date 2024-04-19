import express, { Express, Request, Response } from "express";
import cors from 'cors';
import mongoose from "mongoose";

const MONGO_IP = process.env.MONGO_IP || "130.211.113.94";
const app: Express = express();
const port = 3000;

interface Log {
  data: Data;
  createdat: string;
}

interface Data {
  name: string;
  album: string;
  year: string;
  rank: string;
}

// Enable All CORS Requests
app.use(cors());

// Connect to MongoDB
mongoose
  .connect(`mongodb://${MONGO_IP}:27017/proyecto2`)
  .then(() => console.log("Connected to MongoDB"))
  .catch((err: any) => console.error("Error connecting to MongoDB", err));

// Define a simple route
app.get("/", (req: Request, res: Response) => {
  res.send("Hello, World!");
});

// get all logs from logs collection
app.get("/logs", async (req: Request, res: Response) => {
  // get all logs from logs collection
  const coll = mongoose.connection.collection("logs");
  const logs = await coll.find({}).toArray();

  const logsArray: Log[] = [];

  logs.forEach((log: any) => {
    logsArray.push({
      data: log.data,
      createdat: log.createdat,
    });
  });

  res.json(logsArray);
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
