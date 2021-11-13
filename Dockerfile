# El contenedor base es Alpine en su versión 3.14 (menos de 3MB)
# Con este se crea el contenedor con todo lo necesario para ejecutar los tests
FROM alpine:3.14

# Como root, actualizar todos los paquetes básicos de Alpine
RUN apk update && \
    apk upgrade &&