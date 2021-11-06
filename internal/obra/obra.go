package obra

import (
	"errors"
)

type Videojuego struct {
	Titulo string
}

type Pelicula struct {
	Titulo string
}

type Serie struct {
	Titulo    string
	Temporada int
	Capitulo  int
}

func NewVideojuego(titulo string) (Videojuego, error) {
	if titulo == "" {
		return Videojuego{}, errors.New("El videojuego no tiene título")
	}

	return Videojuego{
		Titulo: titulo,
	}, nil
}

func NewPelicula(titulo string) (Pelicula, error) {
	if titulo == "" {
		return Pelicula{}, errors.New("La película no tiene título")
	}

	return Pelicula{
		Titulo: titulo,
	}, nil
}

func NewSerie(titulo string, temporada, capitulo int) (Serie, error) {
	if titulo == "" {
		return Serie{}, errors.New("La serie no tiene título")
	}

	if capitulo <= 0 || temporada <= 0 {
		return Serie{}, errors.New("El capítulo y temporada de la serie no es correcta")
	}

	return Serie{
		Titulo:    titulo,
		Temporada: temporada,
		Capitulo:  capitulo,
	}, nil
}
