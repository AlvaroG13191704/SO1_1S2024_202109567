"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const cors_1 = __importDefault(require("cors"));
const mongoose_1 = __importDefault(require("mongoose"));
const MONGO_IP = process.env.MONGO_IP || "130.211.113.94";
const app = (0, express_1.default)();
const port = 3000;
// Enable All CORS Requests
app.use((0, cors_1.default)());
// Connect to MongoDB
mongoose_1.default
    .connect(`mongodb://${MONGO_IP}:27017/proyecto2`)
    .then(() => console.log("Connected to MongoDB"))
    .catch((err) => console.error("Error connecting to MongoDB", err));
// Define a simple route
app.get("/", (req, res) => {
    res.send("Hello, World!");
});
// get all logs from logs collection
app.get("/logs", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    // get all logs from logs collection
    const coll = mongoose_1.default.connection.collection("logs");
    const logs = yield coll.find({})
        .sort({ createdat: -1 })
        .limit(20)
        .toArray();
    const logsArray = [];
    logs.forEach((log) => {
        logsArray.push({
            data: log.data,
            createdat: log.createdat,
        });
    });
    res.json(logsArray);
}));
app.listen(port, () => {
    console.log(`Server is running on port http://localhost:${port}`);
});
