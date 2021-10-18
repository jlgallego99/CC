# Lógica de negocio

Mediante un algoritmo Knn se le dotará al sistema de la lógica necesaria para un sistema de clasificación y recomendación personalizada para cada usuario. En él, cada usuario podrá tener recomendaciones de obras en las cuales figuren canciones como las que le gustan y estén acorde a su estado anímico. Este estado anímico estará comprendido entre varios valores como alegre, triste, reflexivo, cansado, etc, definidos en el propio perfil del usuario.

El sistema será multiusuario, todos ellos colaborando para aportar más información al sistema sobre las canciones. Los usuarios pueden:
- Añadir información sobre una canción asociada a una serie, película o videojuego.
- Buscar por canción o por obra (serie, videojuego o película) y así poder ver la banda sonora asociada y su información.
- Indicar que una canción le gusta.
- Recibir recomendaciones personalizadas.

La información de cada canción de la banda sonora incluye:
- Título
- Compositor
- Género
- El estado que le transmite al usuario (tristeza, alegría, epicidad, reflexión, cansancio, etc). Se tendrá un porcentaje global de las sensaciones que le ha transmitido a los distintos usuarios del sistema.
- Momento en el que suena:
    - En el caso de series y películas, el minuto en el que suena, o un momento concreto como el inicio o los créditos. 
    - En el caso de videojuegos, se asociará a un apartado fácilmente reconocible, como puede ser temas de personajes, de lugares, ciudades, cinemáticas, etc. Se acepta que una canción pertenezca a más de un momento, ya que es posible.

En general, los usuarios contribuyen a dar información sobre la canción y con la de todos se forma el sistema con lo necesario para dar recomendaciones.

Se validará que tanto canción como compositor existan, y que esta esté creada por el compositor o artista indicado. También que la serie, película o videojuego existan, y que el minuto o momento en el que suena la canción sea real y posible. 