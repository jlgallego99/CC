package obra

import (
	"errors"
	"reflect"

	"github.com/jlgallego99/OSTfind/internal/cancion"
)

type Obra interface {
	Canciones() []cancion.Cancion_info
	Momento() string
	NuevaCancion(c cancion.Cancion_info) error
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

func (v *Videojuego) NuevaCancion(c cancion.Cancion_info) error {
	if existe, _ := c.ExisteEn(v.OST); existe {
		return errors.New("La canción ya existe en la OST")
	}

	if reflect.DeepEqual(cancion.Cancion_info{}, c) {
		return errors.New("La canción está vacía")
	}

	v.OST = append(v.OST, c)

	return nil
}

func (p *Pelicula) NuevaCancion(c cancion.Cancion_info) error {
	if existe, _ := c.ExisteEn(p.OST); existe {
		return errors.New("La canción ya existe en la OST")
	}

	if reflect.DeepEqual(cancion.Cancion_info{}, c) {
		return errors.New("La canción está vacía")
	}

	p.OST = append(p.OST, c)

	return nil
}

func (s *Serie) NuevaCancion(c cancion.Cancion_info) error {
	if existe, _ := c.ExisteEn(s.OST); existe {
		return errors.New("La canción ya existe en la OST")
	}

	if reflect.DeepEqual(cancion.Cancion_info{}, c) {
		return errors.New("La canción está vacía")
	}

	s.OST = append(s.OST, c)

	return nil
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
