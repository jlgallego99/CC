package obra

import (
	"errors"
)

type ObraBase struct {
	titulo string
}

type Videojuego struct {
	ObraBase
}

type Pelicula struct {
	ObraBase
}

type Serie struct {
	ObraBase
	temporada int
	capitulo  int
}

type Obra interface {
	Titulo() string
	Temporada() int
	Capitulo() int
}

func (o ObraBase) Titulo() string {
	return o.titulo
}

func (s Serie) Temporada() int {
	return s.temporada
}

func (s Serie) Capitulo() int {
	return s.capitulo
}

func NewVideojuego(titulo string) (Videojuego, error) {
	if titulo == "" {
		return Videojuego{}, errors.New("el videojuego no tiene título")
	}

	return Videojuego{
		ObraBase: ObraBase{titulo: titulo},
	}, nil
}

func NewPelicula(titulo string) (Pelicula, error) {
	if titulo == "" {
		return Pelicula{}, errors.New("la película no tiene título")
	}

	return Pelicula{
		ObraBase: ObraBase{titulo: titulo},
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
		ObraBase:  ObraBase{titulo: titulo},
		temporada: temporada,
		capitulo:  capitulo,
	}, nil
}
