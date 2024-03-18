
# Tarea 3  - Alvaro García 202109567

## [Video explicación](https://youtu.be/5cz6gRAAOYw)


## Crear un servidor de Node js que publique mensajes en un canal de Redis

### Instalar Nodejs
Actualizar el sistema y luego instalar Nodejs y npm

```bash
sudo apt-get update
sudo apt-get install nodejs
sudo apt-get install npm
```

### Verificar la instalación
```bash
node -v
npm -v
```

### Crear un proyecto de Nodejs
```bash
mkdir publisher
cd publisher
npm init -y
npm install express
npm install ioredis
```

### Crear un servidor de Nodejs simple

Crear un archivo llamado `index.js` y agregar el siguiente código

```bash
touch index.js
# abrir
vim index.js
```

```javascript
const express = require("express");
const redis = require("ioredis");

// Create a new Express application
const app = express();

// Create a new Redis client
const client = new redis({
  host: "10.233.97.107",
  port: 6379,
});

client.on("error", (err) => {
  console.log("Redis error: ", err);
});

// Define a simple route
app.get("/", async (req, res) => {
  const message = JSON.stringify({ "msg": "Hola a todos" });
  await client.publish("test", message );
  res.send("OK");
});

// Start the server
app.listen(3000, "0.0.0.0", () => {
  console.log("Server running at http://0.0.0.0:3000/");
});
```

### Iniciar el servidor
```bash
node index.js
```

### Acceder al servidor
Agregando la IP del a VM `http://34.125.109.105:3000`


## Crear el subscriptor en Python

### Instalar Python
```bash
sudo apt-get update
sudo apt-get install python3
sudo apt-get install python3-pip
```

### Verificar la instalación
```bash
python3 --version
pip3 --version
```

### Instalar el cliente de Redis
```bash
pip3 install redis
```

### Crear el subscriptor
```bash
touch main.py
# abrir
vim main.py
```

```python
import redis
import json

rdb = redis.Redis(
    host='10.233.97.107',
    port=6379,
    decode_responses=True,
)

# Create a pubsub object
p = rdb.pubsub()

# Subscribe to the 'test' channel
p.subscribe('test')

# Listen for messages
while True:
    message = p.get_message()
    if message:
        data = message.get('data')
        if data and isinstance(data, str):
            data_dict = json.loads(data)
            msg = data_dict.get('msg')
            print('Received msg: ', msg)
```

### Iniciar el subscriptor
```bash
python main.py
```

