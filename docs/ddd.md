Según el domain driven design se desarrollarán varios módulos y clases en base a las siguientes entidades, objetos valor y agregados que representan el dominio del problema con el que trabaja el proyecto.

# Entidades
- **Canción**
- **Obra**:
    - **Película**
    - **Videojuego**
    - **Serie**
- **Usuario**

# Objetos valor
- **Like/Dislike**
- **Sensación (mood)**
- **Título**
- **Compositor o artista**
- **Género**
- **Momento**
- **Momento en minutos**

# Agregados
- **Todas las canciones de una obra**: Cada obra tiene un número de canciones que componen su banda sonora.
- **Título de una canción**
- **Género de una canción**
- **Compositor o artista de una canción**
- **Porcentaje like/dislike de una cancion**: Cada canción puede tener un porcentaje que indique cuanto ha gustado o no a los usuarios.
- **Porcentaje de sensaciones que transmite una cancion**: Cada canción puede tener varias sensaciones aportadas por los usuarios, agrupadas en porcentajes.
- **Momento en minutos en el que suena una canción de una serie o película**
- **Momento en el que suena una canción de un videojuego**: Esto puede ser una zona, ciudad, batalla, etc.
- **Canciones que le gustan a un usuario**: Cada usuario tiene unas canciones que le gustan.
- **Canciones recomendadas a un usuario**: Cada usuario tiene una serie de recomendaciones en base a las canciones que le gustan.
- **Canciones relacionadas**: Cada canción tiene una serie de canciones que son parecidas a ella, a modo de recomendación según género, compositor y sensaciones.