### Para el cliente y servidor hechos en Golang que usan gRPC instalar lo siguiente
```bash
sudo apt install protobuf-compiler # en linux

go get google.golang.org/grpc

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH=$PATH:$(go env GOPATH)/bin  # agregar al path cuando se quiera compilar los archivos .proto

protoc --go_out=. --go-grpc_out=. [name].proto # compilar el archivo .proto  
```

### Para el cliente instalar lo siguiente
```bash
go get github.com/gofiber/fiber/v2 # fiber
```

### Para el servidor instalar lo siguiente
```bash
go get github.com/confluentinc/confluent-kafka-go # kafka
```

### crear imagen de client

```bash
docker build -t alvarog1318/s1p2_grcp_producer_client:v0.1 . # crear imagen de cliente

docker build -t alvarog1318/s1p2_grcp_producer_server:v0.1 . # crear imagen de servidor
```