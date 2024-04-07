## Creación de los Cluster en GPC

### Comandos para crear un cluster en GPC

```bash
gcloud init # Inicializa la configuración de gcloud

# Escgoer un usuario y proyecto

# Establecer zona central us-central1-a

gcloud container clusters create cluster-name --num-nodes=3 --machine-type=n1-standard-2 --zone us-central1-a
```

