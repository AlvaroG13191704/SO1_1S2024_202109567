## Creación de los Cluster en GPC

### Comandos para crear un cluster en GPC

```bash
# Inicializa la configuración de gcloud
gcloud init

# Escgoer un usuario y proyecto

# Establecer zona central us-central1-a

gcloud container clusters create cluster-name --num-nodes=3 --machine-type=n1-standard-2 --zone us-central1

```

### Deployments
Los deployments son una forma de declarar el estado deseado de una aplicación en un cluster de Kubernetes. Un deployment se encarga de crear y actualizar los pods de una aplicación.

### Namespaces
Los namespaces son una forma de dividir los recursos de un cluster en grupos lógicos. Por ejemplo, si tenemos un cluster con multiples aplicaciones, podemos crear un namespace para cada una de ellas.

### Ingress
Un ingress es un recurso de Kubernetes que permite exponer servicios HTTP y HTTPS a través de un balanceador de carga. Un ingress se encarga de redirigir el tráfico a los servicios correspondientes.

### Service
Un service es un recurso de Kubernetes que permite exponer un conjunto de pods como un servicio. Un service se encarga de redirigir el tráfico a los pods correspondientes.

### Buenas prácticas
- Utilizar namespaces para dividir los recursos de un cluster en grupos lógicos.
- Al crear los deployments, construir la imagen del contenedor principal, los demas contenedores son de apoyo para el principal.

### Kubelet
El kubelet es un agente que se ejecuta en cada nodo de un cluster de Kubernetes. El kubelet se encarga de gestionar los pods y los contenedores en un nodo.

### Kubernetes API
Es el corazón del control plane. El API server expone un HTTP API que permite a los usuarios y a los componentes de Kubernetes comunicarse con el cluster.


To create a Deployment.yaml file to deploy a Kafka pod using Strimzi, you can use the following example as a template:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-kafka-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: my-kafka
  template:
    metadata:
      labels:
        app: my-kafka
    spec:
      containers:
      - name: kafka
        image: quay.io/strimzi/kafka:0.29.0-kafka-3.2.3
        resources:
          requests:
            cpu: 500m
            memory: 1Gi
          limits:
            cpu: 1
            memory: 2Gi
        env:
        - name: KAFKA_BROKER_ID
          value: "0"
        - name: KAFKA_ADVERTISED_LISTENERS
          value: PLAINTEXT://:9092
        - name: KAFKA_LISTENER_SECURITY_PROTOCOL_MAP
          value: PLAINTEXT:PLAINTEXT
        - name: KAFKA_INTER_BROKER_LISTENER_NAME
          value: PLAINTEXT
        - name: KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR
          value: "1"
        - name: KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR
          value: "1"
        - name: KAFKA_TRANSACTION_STATE_LOG_MIN_ISR
          value: "1"
        - name: KAFKA_LOG_DIRS
          value: /var/lib/kafka/data
        volumeMounts:
        - name: kafka-data
          mountPath: /var/lib/kafka/data
      volumes:
      - name: kafka-data
        emptyDir: {}
```

Here's a breakdown of the key components in this Deployment.yaml file:

1. `apiVersion`: Specifies the Kubernetes API version to use, in this case, `apps/v1`.
2. `kind`: Defines the type of Kubernetes object, which is a `Deployment`.
3. `metadata`: Provides the name of the Deployment.
4. `spec.replicas`: Specifies the number of replicas (Kafka broker pods) to deploy.
5. `spec.selector`: Defines the label selector for the Pods that the Deployment manages.
6. `spec.template`: Defines the Pod template, including the container specification.
7. `spec.template.metadata.labels`: Assigns the label `app: my-kafka` to the Pods.
8. `spec.template.spec.containers`: Defines the container(s) to be deployed. In this case, it's a single container with the Kafka image.
9. `spec.template.spec.containers.resources`: Specifies the CPU and memory requests and limits for the Kafka container.
10. `spec.template.spec.containers.env`: Sets various Kafka configuration environment variables.
11. `spec.template.spec.containers.volumeMounts`: Mounts the `kafka-data` volume to the `/var/lib/kafka/data` directory inside the container.
12. `spec.template.spec.volumes`: Defines the `kafka-data` volume as an `emptyDir` volume.

To use this Deployment.yaml file, you'll need to have the Strimzi Operator installed in your Kubernetes cluster. Once the Operator is set up, you can apply the Deployment.yaml file using the following command:

```
kubectl apply -f Deployment.yaml
```

This will create the Kafka Deployment in your cluster, and the Strimzi Operator will handle the provisioning of the Kafka broker pod.