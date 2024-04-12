### Pods
Los pods en Kubernetes son usados de dos formas
1. Pods que corren un solo contenedor "one-container-per-pod" es el modelo más comun.
2. Pods que corren varios contenedores que necesitan trabajar en conjunto.

### Deployments
Los deployments son una forma de declarar el estado deseado de una aplicación en un cluster de Kubernetes. Un deployment se encarga de crear y actualizar los pods de una aplicación.

### Namespaces
Los namespaces son una forma de dividir los recursos de un cluster en grupos lógicos. Por ejemplo, si tenemos un cluster con multiples aplicaciones, podemos crear un namespace para cada una de ellas.

### Ingress
Un ingress es un recurso de Kubernetes que permite exponer servicios HTTP y HTTPS a través de un balanceador de carga. Un ingress se encarga de redirigir el tráfico a los servicios correspondientes.

### Service
Un service es un recurso de Kubernetes que permite exponer un conjunto de pods como un servicio. Un service se encarga de redirigir el tráfico a los pods correspondientes.

Cuando un servicio es de tipo 
- ClusterIP: El servicio solo es accesible desde dentro del cluster.
- NodePort: El servicio es accesible desde fuera del cluster a través de un puerto específico en cada nodo.
- LoadBalancer: El servicio es accesible desde fuera del cluster a través de un balanceador de carga.

### Buenas prácticas
- Utilizar namespaces para dividir los recursos de un cluster en grupos lógicos.
- Al crear los deployments, construir la imagen del contenedor principal, los demas contenedores son de apoyo para el principal.

### Kubelet
El kubelet es un agente que se ejecuta en cada nodo de un cluster de Kubernetes. El kubelet se encarga de gestionar los pods y los contenedores en un nodo.

### Kubernetes API
Es el corazón del control plane. El API server expone un HTTP API que permite a los usuarios y a los componentes de Kubernetes comunicarse con el cluster.

