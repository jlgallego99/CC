package obra

import (
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

type Obra interface {
	Canciones() []cancion.Cancion
	Momento() string
	nuevaCancion(c cancion.Cancion)
}

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
	temporada int
	capitulo  int
	canciones []cancion.Cancion
}

func (v *Videojuego) Canciones() []cancion.Cancion {
	return nil
}

func (p *Pelicula) Canciones() []cancion.Cancion {
	return nil
}

func (s *Serie) Canciones() []cancion.Cancion {
	return nil
}

func (v *Videojuego) Momento() string {
	return ""
}

func (p *Pelicula) Momento() string {
	return ""
}

func (s *Serie) Momento() string {
	return ""
}

func (v *Videojuego) nuevaCancion(c cancion.Cancion) {

}

func (p *Pelicula) nuevaCancion(c cancion.Cancion) {

}

func (s *Serie) nuevaCancion(c cancion.Cancion) {

}

func newVideojuego(titulo string, canciones []cancion.Cancion) *Videojuego {
	return &Videojuego{
		titulo:    titulo,
		canciones: canciones,
	}
}

func newPelicula(titulo string, canciones []cancion.Cancion) *Pelicula {
	return &Pelicula{
		titulo:    titulo,
		canciones: canciones,
	}
}

func newSerie(titulo string, temporada, capitulo int, canciones []cancion.Cancion) *Serie {
	return &Serie{
		titulo:    titulo,
		temporada: temporada,
		capitulo:  capitulo,
		canciones: canciones,
	}
}
