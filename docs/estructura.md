# Estructura de directorios del proyecto
Siguiendo las guías de diseño de [GoApp](https://github.com/bnkamalesh/goapp) y [project-layout](https://github.com/golang-standards/project-layout), enfocadas a la creación de aplicaciones web en Go siguiendo las prácticas más seguidas, se ha estructurado la aplicación en una serie de directorios esenciales. Cada directorio contiene un paquete de Go, es decir, un conjunto de archivos con un mismo propósito.

En el directorio **internal** se tiene todo el código del proyecto. Este directorio está definido como estándar por la [especificación de Go](https://go.dev/doc/go1.4#internalpackages), y su objetivo es que todo este código sea "privado" a la propia aplicación, no pudiendo importarse fuera del subárbol en el que residen. Así, se asegura que todo el código que se tiene aquí no se pueda importar en lugares que no convenga. Dentro de ello se hace la siguiente división:
- Un directorio por cada paquete que englobe un componente concreto de la lógica de negocio (agregados, entidades y objetos valor), en este caso se tiene **cancion**, **obra** y **usuario**. Estos tienen ficheros con la definición de las funciones y los tests para cada fichero (terminado en *_test*).
- En el directorio **server** se tiene todo el código relativo al API REST. Se definen en un fichero las rutas, y en un fichero los handlers asociados a esas rutas, que gestionan las peticiones HTTP. Estos handlers llaman a las funciones de los paquetes definidos anteriormente, haciendo así una separación entre el API REST y la lógica de negocio.
- En el directorio **config** se tiene el código y diversos ficheros necesarios para la configuración de cualquiera de los servicios usados en el proyecto. En concreto, el fichero *config.go* define funciones para iniciar los servicios de configuración como etcd y dotenv, y recuperar las variables que puedan necesitar otros paquetes para iniciar servicios.