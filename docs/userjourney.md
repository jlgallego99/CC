# Usuarios
Existen dos tipos de roles en los **usuarios**:
- **Buscadores**: Usuarios no registrados que buscan información de las canciones y obras relacionadas en la aplicación. También pueden ver las recomendaciones que se ofrecen por cada canción y obra (otras relacionadas que tengan similitud con las buscadas).
- **Colaboradores**: Usuarios registrados que pueden contribuir al sistema aportando información a las canciones, además de recibir recomendaciones personalizadas. También pueden hacer todo lo que hacen los buscadores.

# User journeys
A continuación se describe lo que pueden hacer estos usuarios en la aplicación mediante **user journeys**:
- **Buscadores**:
    - **UJ.1** 
    > Una persona está viendo la serie The Boys, y por la mitad del capítulo que está viendo escucha una canción que le gusta mucho y quiere saber cual es.   
    > Este buscador busca en OSTfind el capítulo que estaba viendo de la serie The Boys.  
    > La serie tiene una la lista de canciones, en ella salen todas las canciones que suenan en el capítulo.
    > Ve el minuto en el que suena la canción, que es Wannabe de las Spice Girls.

    - **UJ.2**
    > Una persona acaba de terminar el videojuego Celeste y, puesto que una canción en concreto le ha gustado mucho, quiere conocer otras más de ese tipo.  
    > Este buscador busca en OSTfind la canción de este juego introduciendo su título.  
    > Puede ver la información de esta canción: su título, género, momento en el que suena, si le ha gustado a la gente y las sensaciones que les ha transmitido.  
    > El sistema le muestra como recomendación una lista de canciones que están relacionadas con esta.  
    > Esta persona además puede ver los juegos en los que suenan estas nuevas canciones que ha descubierto.  

    - **UJ.3**
    > Una persona está triste y quiere ver películas con canciones melancólicas.  
    > Busca en OSTfind canciones cuya sensación más predominante sea la tristeza.  
    > El sistema le ofrece una lista con canciones del tipo que ha buscado.  
    > Puede ver la información de cada una de estas canciones y a qué película pertenecen.  

- **Colaboradores**:
    - **UJ.4**
    > Una persona se ha visto la película El Castillo Ambulante y le encanta su música.  
    > Llega a OSTfind y quiere aportar información ya que conoce muy bien su música.  
    > Añade las canciones que existen en su banda sonora indicando su título, género, compositor, las sensaciones que le transmite y el minuto en el que suenan.  
    > Ahora otros usuarios, colaboradores o buscadores, podrán tener mejor información de la banda sonora de esta película.  
    > Además indica cuales son las canciones de esa banda sonora que más le gustan.   
    > El sistema le recomendará nuevas canciones basándose en las que ha indicado que le gustan.  

    - **UJ.5**
    > Una persona se ha jugado un videojuego en el que en una ciudad suena una canción que le encanta.  
    > Quiere indicar en OSTfind que le gusta esa canción, ya que quiere saber qué otras canciones parecidas que suenen en ciudades de videojuegos puede escuchar.  
    > Descubre que esa canción no está registrada, por lo que la añade él mismo indicando su título, videojuego al que pertenece, género, sensaciones, momento en el que suena (indicando que es una ciudad, y además indicando exactamente en qué ciudad) y si le gusta.  
    > Ahora esa canción figura en la banda sonora del videojuego, y ya puede saber qué otras canciones están relacionadas con esta.  

Podemos ver que el punto central de OSTfind son las canciones, y todas las búsquedas y operaciones se realizan en base a esta entidad.