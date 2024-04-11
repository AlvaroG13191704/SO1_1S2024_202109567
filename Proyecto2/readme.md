## Creación de los Cluster en GPC

### Comandos para crear un cluster en GPC

```bash
# Inicializa la configuración de gcloud
gcloud init

# Escgoer un usuario y proyecto

# Establecer zona central us-central1-a

# instalar kubectl
gcloud components install kubectl

# Crear un cluster
gcloud container clusters create proyecto2 --num-nodes=1 --tags=allin --tags=allout --enable-legacy-authorization --issue-client-certificate --preemptible --machine-type=n1-standard-2

# conectar kubectl al cluster
gcloud container clusters get-credentials proyecto2 --zone us-central1-a --project sopes1-417522

# External API of mongo-service
NAME            TYPE           CLUSTER-IP     EXTERNAL-IP      PORT(S)           AGE
mongo-service   LoadBalancer   10.33.91.249   34.31.32.118     27017:31001/TCP   101m
redis-service   LoadBalancer   10.33.93.93    35.224.242.190   6379:32096/TCP    47s

```

### Conectarse a mongocompass

```bash
mongodb://34.31.32.118:27017/
```

### Conectarse a redis 

```bash
35.224.242.190:6379
```

Para comunicar los servicios con mongo y redis es tan simple como usar el nombre 
de los servicios al momento de conectarse con código a los servicios.
