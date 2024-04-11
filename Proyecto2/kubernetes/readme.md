## Acá estan los deployments para las dos bases de datos a utilizar en el proyecto


### mongo-deployment.yaml

Comandos para crear un deployment de mongo en kubernetes

```bash
kubectl apply -f mongo-deployment.yaml # deploy mongo pod

kubectl apply -f mongo-service.yaml # deploy mongo service

```

### Comandos esenciales

```bash
kubectl get namespaces # Ver namespaces

kubectl get nodes # Ver nodos

kubectl get pods # Ver pods

kubectl get services -n <nombre_servicio>
kubectl delete services -n <nombre_servicio> 

kubectl get deployments -n <nombre_deployment>
kubectl delete deployments -n <nombre_deployment>

# Archivos de configuracion de Kubernetes
kubectl get [nodes|deployments|services|pods] nombre -o yaml > pod.yaml
kubectl create -f archivo.yaml
kubectl delete -f archivo.yaml
kubectl apply -f archivo.yaml

# Monitoreo
kubectl logs -f pod/name
kubectl logs -f deployment/name
kubectl describe deployments name
```