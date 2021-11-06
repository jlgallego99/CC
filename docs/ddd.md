Según el domain driven design se desarrollarán varios módulos y clases en base a las siguientes entidades, objetos valor y agregados que representan el dominio del problema con el que trabaja el proyecto.

# Entidades
- **Canción**: Cada canción tiene identidad propia, es única por su identificador, y además es mutable puesto que tiene parámetros que van variando en el tiempo como las sensaciones y los likes/dislikes, que se van modificando.
- **Usuario**: Cada usuario tiene identidad propia y es único por su nombre de usuario. Es mutable puesto que tiene parámetros que van variando en el tiempo, como las canciones que le gustan, las que no le gustan o las canciones en las que ha colaborado de alguna manera (añadiéndolas o aportando sensaciones).

# Objetos valor
- **Obra**: Estos tres objetos valor representan los tipos de obras que existen en el proyecto. Son inmutables, simplemente definen métodos para acceder a sus atributos y realizan validaciones para que sean objetos correctos.
    - **Película**
    - **Videojuego**
    - **Serie**

# Agregados
- **OST**: Todas las canciones de una obra (película, videojuego o serie) componen su banda sonora (OST).