package obra

import (
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

type Obra interface {
	Canciones() []cancion.Cancion_info
	Momento() string
	NuevaCancion(c cancion.Cancion_info)
}

type Videojuego struct {
	Titulo string
	OST    []cancion.Cancion_info
}

type Pelicula struct {
	Titulo string
	OST    []cancion.Cancion_info
}

type Serie struct {
	Titulo    string
	Temporada int
	Capitulo  int
	OST       []cancion.Cancion_info
}

func (v *Videojuego) Canciones() []cancion.Cancion_info {
	return nil
}

func (p *Pelicula) Canciones() []cancion.Cancion_info {
	return nil
}

func (s *Serie) Canciones() []cancion.Cancion_info {
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

func (v *Videojuego) NuevaCancion(c cancion.Cancion_info) {

}

func (p *Pelicula) NuevaCancion(c cancion.Cancion_info) {

}

func (s *Serie) NuevaCancion(c cancion.Cancion_info) {

}

func NewVideojuego(titulo string, canciones []cancion.Cancion_info) *Videojuego {
	return &Videojuego{
		Titulo: titulo,
		OST:    canciones,
	}
}

func NewPelicula(titulo string, canciones []cancion.Cancion_info) *Pelicula {
	return &Pelicula{
		Titulo: titulo,
		OST:    canciones,
	}
}

func NewSerie(titulo string, temporada, capitulo int, canciones []cancion.Cancion_info) *Serie {
	return &Serie{
		Titulo:    titulo,
		Temporada: temporada,
		Capitulo:  capitulo,
		OST:       canciones,
	}
}
