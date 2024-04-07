# Tarea 4  - Alvaro García 202109567

## [Video explicación](https://youtu.be/5cz6gRAAOYw)


## Locust y gRPC

### Instalar lo necesario para grpc en golang

```bash
sudo apt install protobuf-compiler # en linux

go get google.golang.org/grpc

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH=$PATH:$(go env GOPATH)/bin  # agregar al path cuando se quiera compilar los archivos .proto

protoc --go_out=. --go-grpc_out=. [name].proto # compilar el archivo .proto    
```

### Instalar locust
```bash
pip install locust
```


### Levantar los serviodores
```bash
locust -f locustfile.py # levantar locust
go run main.go # levantar el servidor y client grpc
```

