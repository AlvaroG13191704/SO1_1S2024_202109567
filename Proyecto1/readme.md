# Proyecto 1 
# Alvaro Norberto García Meza - 202109567

## Introducción
Este proyecto consiste en la creación de una aplicación web que muestra 
información del estado de la VM en la que se encuentra alojada. La aplicación, como por ejemplo el uso de CPU, memoria, vista de los procesos y simulación. Para esto se consruyerón distintos contenedores que se comunican entre sí.

## Modulos
Para obtener la información de la VM se crearon distintos módulos que se encargan de obtener la información necesaria. Estos módulos se encuentran en la carpeta `modules` y son los siguientes:

### CPU
El módulo de CPU se encarga de obtener la información del uso de CPU. Para esto se utilizó el archivo `cpu.c` que se encarga de obtener la información y enviarla al backend. Para compilar el archivo se utilizó el siguiente comando:
```bash
make
sudo insmod cpu.ko
```

### Memoria
El módulo de memoria se encarga de obtener la información del uso de memoria. Para esto se utilizó el archivo `ram.c` que se encarga de obtener la información y enviarla al backend. Para compilar el archivo se utilizó el siguiente comando:
```bash
make
sudo insmod ram.ko
```

## Contenedores
Se crearon 3 contenedores, uno para el cliente, otro para el servidor y otro para la base de datos. El cliente es el encargado de mostrar la información, el servidor de obtenerla y la base de datos de almacenarla.
### Backend
El backend fue construido con Fiber, un framework de Go. Se encarga de obtener la información de la `/proc` y de enviarla al cliente. La información que se obtiene es la siguiente:

- CPU
- Memoria
- Procesos

*Para saber más sobre la información extraida se puede consultar los archivos dentro de `/modules` y ver el código escrito en C para `ram.c` y `cpu.c`*

### DockerFile
El DockerFile del backend se encuentra en la carpeta `backend`
Este dockerfile se encarga de construir la imagen del backend, para esto se utiliza una imagen de `golang:alpine` para compilar el código y luego se copia el binario a una imagen de `alpine:latest` para reducir el tamaño de la imagen.
```bash
#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v
RUN go build -o /go/bin/app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
LABEL Name=proyecto1 Version=0.0.1
EXPOSE 8080
```

### Frontend
El frontend fue construido con Next js. Se encarga de mostrar la información que le envía el backend. La información que se muestra es la siguiente:

### DockerFile
El DockerFile del frontend se encuentra en la carpeta `frontend`
Este dockerfile se encarga de construir la imagen del frontend, para esto se utiliza una imagen de `node:18-alpine`, luego un segundo stage para instalar las dependencias y finalmente un tercer stage para copiar los archivos necesarios y correr el servidor.
```bash
FROM node:18-alpine AS base

# Install dependencies only when needed
FROM base AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

# Install dependencies based on the preferred package manager
COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* ./
RUN \
    if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
    elif [ -f package-lock.json ]; then npm ci; \
    elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm i --frozen-lockfile; \
    else echo "Lockfile not found." && exit 1; \
    fi


# Rebuild the source code only when needed
FROM base AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

RUN yarn build

# Production image, copy all the files and run next
FROM base AS runner
WORKDIR /app

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public

# Set the correct permission for prerender cache
RUN mkdir .next
RUN chown nextjs:nodejs .next

# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 80

ENV PORT 80
# set hostname to localhost
ENV HOSTNAME "0.0.0.0"

# server.js is created by next build from the standalone output
# https://nextjs.org/docs/pages/api-reference/next-config-js/output
CMD ["node", "server.js"]
```

### Base de datos
La base de datos fue construida con MongoDB. Se encarga de almacenar la información que le envía el backend. La imagen utilizada fue la de mySQL


### Configuración del nginx
Se utilizó un reverse proxy para poder comunicar los contenedores entre sí. El archivo de configuración se encuentra en `nginx.conf`. El archivo se encarga de redirigir las peticiones a los contenedores correspondientes. Evitando el uso de la IP de la máquina y utilizando los nombres de los contenedores.
```bash
events {}

http {
    server {
        listen 80;

        location / {
            proxy_pass http://frontend:3000;
        }

        location /api/ {
            proxy_pass http://backend:8080;
        }
    }
}
```

### Docker-compose
El archivo de docker-compose se encarga de levantar los contenedores y de configurar la red entre ellos. Además de configurar el volumen para la base de datos.
```bash
version: '3'
services:
  nginx:
    image: nginx:latest
    ports:
      - "80:80"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
    depends_on:
      - frontend
      - backend
  frontend:
    image: alvarog1318/s1p1_frontend:v0.2
    expose:
      - 3000
    environment:
      - PORT=3000
    depends_on:
      - backend
  backend:
    image: alvarog1318/s1p1_backend:v0.2
    privileged: true
    pid: host
    restart: always
    expose:
      - 8080
    volumes:
      - /proc:/host_proc
    depends_on:
      - db
  db:
    image: mysql:latest
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: admin
      MYSQL_DATABASE: Proyecto1
    ports:
      - 3306:3306
    volumes:
      - db_data:/var/lib/mysql
volumes:
  db_data:
```

### Docker push
Se subieron las imágenes a docker hub para poder ser utilizadas en cualquier máquina. Para esto se utilizó el comando `docker push` para subir las imágenes.

*Construir imagen del backend*
```bash
# correr el comando en la carpeta del backend
docker build -t alvarog1318/s1p1_backend:v0.2 .
```

*Construir imagen del frontend*
```bash
# correr el comando en la carpeta del frontend
docker build -t alvarog1318/s1p1_frontend:v0.2 .
```