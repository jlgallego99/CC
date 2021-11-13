# El contenedor base es Alpine en su versión 3.14 (menos de 3MB)
# Con este se crea el contenedor con todo lo necesario para ejecutar los tests
FROM alpine:3.14

# Como root, actualizar todos los paquetes básicos de Alpine y crear usuario no root
USER root

# Solo el root puede instalar las aplicaciones, si no da error de permiso denegado
RUN apk update && apk upgrade \
    && adduser -D ostfind \
    && apk add --no-cache go curl \
    && mkdir -p /app/test \
    && chown ostfind /app/test \
    && sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin \
    # Borrar paquetes que no se van a usar más
    && apk del curl perl 

# Se cambia a un usuario que no sea root
USER ostfind

# Se hace cd al directorio donde se va a ejecutar todo
WORKDIR /app/test

# Se copia el archivo que contiene las dependencias del host a dentro del docker
COPY go.mod .
COPY go.sum .
# Se copia el archivo que contiene las instrucciones del gestor de tareas del host a dentro del docker
COPY Taskfile.yml .

# Se instalan las dependencias y se borra lo que no es necesario
RUN go mod download \
    && rm go.mod && rm go.sum

# Ejecuta lo que se quiere cuando se inicia el contenedor, en este caso pasa los tests usando el gestor de tareas
CMD ["task", "test"]