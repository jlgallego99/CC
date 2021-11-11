package obra

import (
	"errors"
	"strconv"
)

type Videojuego struct {
	titulo string
}

type Pelicula struct {
	titulo string
}

type Serie struct {
	titulo    string
	temporada int
	capitulo  int
}

type Obra interface {
	Titulo() string
}

func (v Videojuego) Titulo() string {
	return v.titulo
}

func (p Pelicula) Titulo() string {
	return p.titulo
}

func (s Serie) Titulo() string {
	return s.titulo + "-" + strconv.Itoa(s.temporada) + "-" + strconv.Itoa(s.capitulo)
}

func NewVideojuego(titulo string) (Videojuego, error) {
	if titulo == "" {
		return Videojuego{}, errors.New("el videojuego no tiene título")
	}

	return Videojuego{
		titulo: titulo,
	}, nil
}

func NewPelicula(titulo string) (Pelicula, error) {
	if titulo == "" {
		return Pelicula{}, errors.New("la película no tiene título")
	}

	return Pelicula{
		titulo: titulo,
	}, nil
}

func NewSerie(titulo string, temporada, capitulo int) (Serie, error) {
	if titulo == "" {
		return Serie{}, errors.New("la serie no tiene título")
	}

	if capitulo <= 0 || temporada <= 0 {
		return Serie{}, errors.New("el capítulo y temporada de la serie no es correcta")
	}

	return Serie{
		titulo:    titulo,
		temporada: temporada,
		capitulo:  capitulo,
	}, nil
}
