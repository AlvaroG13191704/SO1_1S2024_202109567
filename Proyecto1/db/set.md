## Levantar un contenedor de MySQL
```bash
docker pull mysql

docker run --name gerenciales -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root123 -d mysql:latest 
```

## Para conectarse con datagrip
```bash
# En la pestaña de MySQL
# Host: localhost
# Port: 3306
# User: root
# Password: root123
```

## Resolver problema de “Host ‘172.17.0.1’ 
```bash
# abrir una terminal y ejecutar
docker exec -it db_db_1 mysql -u root -p
# ingresar la contraseña
tu_contraseña
# correr este comando para permitir acceso desde cualquier host
CREATE USER 'newuser'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'newuser'@'%';
FLUSH PRIVILEGES;
```
