package obra

import (
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

type Videojuego struct {
	titulo    string
	canciones []cancion.Cancion
}

type Pelicula struct {
	titulo    string
	canciones []cancion.Cancion
}

type Serie struct {
	titulo    string
	canciones []cancion.Cancion
}
