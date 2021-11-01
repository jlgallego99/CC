# Tareas gestionadas por Task
Mediante el gestor de tareas **Task** se han definido una serie de tareas en el fichero [Taskfile](../Taskfile.yml). Todas ellas se ejecutan estando en la raíz del proyecto en un terminal, y son las siguientes:
- **task**: Indica todas las tareas definidas, con su descripción.
- **task check**: Compila los paquetes que hay en el directorio *internal* para comprobar que no hay errores de compilación. No genera ningún ejecutable, sirve a modo de comprobación de que está todo bien.
- **task run**: Ejecuta directamente OSTfind, sin instalar (actualmente no hace nada puesto que no se ha desarrollado aún el punto de entrada de la aplicación).
- **task install**: Instala OSTfind, generando un ejecutable (actualmente genera un programa que no hace nada puesto que no se ha desarrollado aún el punto de entrada de la aplicación).
- **task test**: Ejecuta todos los tests de OSTfind y muestra sus resultados.