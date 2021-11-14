# Creación de contenedores para un entorno de pruebas

## Docker
Para crear el Dockerfile se ha seguido un proceso ordenado en el cual lo primero y más importante es elegir el contenedor base a partir del cual se va a construir el contenedor necesario para ejecutar el proyecto, y luego se instalarán los paquetes y bibliotecas necesarias para poder ejecutar los tests.

Se ha usado como guía las buenas prácticas de la [documentación de Docker](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) y los ejemplos de optimización de [erseco](https://github.com/erseco/dockerfile-optimization-examples), en el cual se pone de manifiesto que Alpine es una de las mejores opciones para crear contenedores Docker.

### Elección del contenedor base

Para elegir el contenedor base se ha optado por usar Alpine Linux debido a su bajo tamaño, lo cual hará que se tenga un contenedor ligero con lo mínimo necesario para ejecutar los tests y que no se tarde mucho en crear y ejecutar. Sin embargo, existen muchas versiones de [Alpine](https://hub.docker.com/_/alpine?tab=tags), y aún más, existen contenedores bases oficiales de [Go](https://hub.docker.com/_/golang?tab=tags) que también usan Alpine. 

En general, buscamos crear un contenedor que use la última versión estable de Alpine, la *3.14*, ya que nos garantiza tener la versión más moderna y sin problemas ya que es estable. Lo mismo con Go, buscamos la versión *1.17.3* que es la última versión estable y la que se ha usado a lo largo de este proyecto, asegurándonos de que todo va a funcionar bien en esa versión (aunque hay que tener en cuenta que serviría cualquiera de las dos versiones importantes anteriores, la 1.16 y la 1.15 porque son las que siguen recibiendo soporte y son estables). Los contenedores base de Alpine para la versión que queremos tienen un tamaño de menos de 3MB, mientras que los contenedores oficiales de Go que usan Alpine tienen un tamaño de más de 100MB, una diferencia bastante grande. 

Puesto que solo vamos a necesitar cosas básicas como el lenguaje, el gestor de tareas, una biblioteca de tests y nada más, es mejor crear un contenedor desde una imagen básica de Alpine.

Hay que tener en cuenta que el proyecto está en el lenguaje de programación Go, por lo que necesitaremos tener instalado ese lenguaje, además de las bibliotecas necesarias para testear, en este caso *Ginkgo*, y el gestor de tareas *Task*. Se ha buscado en el [repositorio de paquetes de Alpine](https://pkgs.alpinelinux.org/packages) a ver si se podía instalar usando *apk*, y efectivamente existe un paquete con la versión *1.17.3*, que es la que buscamos, por lo que podremos instalar así el lenguaje.

Sin embargo, el gestor de tareas Task no existe en el repositorio de paquetes de alpine, por lo que habría que instalarlo directamente usando curl, como dice en la [documentación oficial](https://taskfile.dev/#/installation?id=build-from-source). El resto de dependencias se instalan al ejecutar los tests, por lo que no es necesario pasar en el dockerfile ningún fichero de dependencias.

Con todo esto ya se tiene un [Dockerfile](../Dockerfile) completo y listo para poder ejecutar los tests de OSTfind. Dentro de él se encuentran comentarios sobre las decisiones tomadas en cada línea, que vienen a representar lo estudiado en este documento. Principalmente se ha buscado reducir el número de capas usando una sola instrucción *RUN* para instalar múltiples paquetes y se ha usado un usuario no privilegiado para ejecutar las tareas. Tanto el lenguaje como el gestor de tareas los debe instalar el root puesto que un usuario no root no tiene permisos para instalar paquetes con apk.

Se probó a usar la imagen de alpine oficial de Golang, sin embargo quedaba una imagen de mas de 500MB, teniendo ademas que instalar el paquete *build-base* en el dockerfile puesto que sin gcc no se puede ejecutar go. Con la imagen actual se tiene un tamaño de unos 450MB (140MB comprimida tal y como se puede ver en [DockerHub](https://hub.docker.com/r/jlgallego99/ostfind/tags)), que no he podido reducir más debido a que la mayoría de ese espacio pertenece a la instalación del lenguaje, y es necesario para ejecutar los tests. Si lo que hiciese el docker fuese ejecutar un binario, se podría hacer con una imagen super ligera en la que primero se compile y se genere un ejecutable, y luego este se copie a una imagen de scratch vacia donde simplemente se ejecute, pero no es el caso.

Es importante tener en cuenta que este contenedor no contiene los fuentes del proyecto, si no que se ejecutará a partir del repositorio de código, redirigiendo los fuentes en una ejecución del contenedor docker (que está en DockerHub), con la siguiente orden:
```
docker run -t -v `pwd`:/app/test jlgallego99/OSTfind
```
Que se puede ejecutar de forma más fácil con la orden del gestor de tareas:
```
task dockertest
```

### Actualización automática del contenedor
El contenedor se ha subido a DockerHub.

## Github Container Registry