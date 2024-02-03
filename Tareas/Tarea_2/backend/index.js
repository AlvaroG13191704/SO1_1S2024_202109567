const express = require('express');
const cors = require('cors');
const MongoClient = require('mongodb').MongoClient;
const app = express();

app.use(cors()); // Enable All CORS Requests
app.use(express.json()); // for parsing application/json

const url = 'mongodb://root:admin@db:27017/';
const dbName = 'lab_sopes_1';

app.post('/upload', async (req, res) => {
  const payload = req.body;

  const client = new MongoClient(url, { useUnifiedTopology: true });

  try {
    await client.connect();
    const db = client.db(dbName);
    const collection = db.collection('tarea_2');

    let result = await collection.insertOne(payload);
    res.status(200).json({ insertedId: result.insertedId });
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'An error occurred while uploading image' });
  } finally {
    await client.close();
  }
});

app.listen(8080, () => {
  console.log('Server is running on port 8080');
});