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
gcloud container clusters create proyecto2 --num-nodes=1 --tags=allin --tags=allout --enable-legacy-authorization --issue-client-certificate --preemptible --machine-type=n1-standard-4

# port forward
kubectl port-forward -n so1-proyecto2 --address 0.0.0.0 svc/grafana 3000:3000


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
