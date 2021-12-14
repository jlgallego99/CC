# Tareas gestionadas por Task
Mediante el gestor de tareas **Task** se han definido una serie de tareas en el fichero [Taskfile](../Taskfile.yml). Todas ellas se ejecutan estando en la raíz del proyecto en un terminal, y son las siguientes:
- **task**: Indica todas las tareas definidas, con su descripción.
- **task install**: Instala por defecto el programa en la ruta GOPATH, sin embargo actualmente esto no se puede hacer porque no hay un archivo main y ni se va a desplegar el servidor aún.
- **task build**: Compila los paquetes que hay en el directorio *internal* para comprobar que no hay errores de compilación. No genera ningún ejecutable, sirve a modo de comprobación de que está todo bien.
- **task test**: Ejecuta todos los tests de OSTfind y muestra sus resultados.
- **taks installdeps**: Mira el fichero go.mod con todas las dependencias que necesita el proyecto y las instala en el sistema para que este pueda funcionar.
- **task dockertest**: Ejecuta el contenedor docker alojado en DockerHub que permite pasar los tests del proyecto.