package usuario

import (
	"errors"

	"github.com/jlgallego99/OSTfind/internal/cancion"
)

type Colaborador struct {
	Nombre             string
	CancionesFavoritas []cancion.Cancion_info
	CancionesOdiadas   []cancion.Cancion_info
}

type Buscador struct {
	Nombre string
}

type Usuario interface {
	Like(c cancion.Cancion) error
	Dislike(c cancion.Cancion) error
	Recomendaciones() ([]cancion.Cancion, error)
}

func (col *Colaborador) Like(c cancion.Cancion_info) error {
	col.CancionesFavoritas = append(col.CancionesFavoritas, c)

	if c.ExisteEn(col.CancionesFavoritas) {
		return errors.New("El usuario ya le ha dado like a esta canción")
	}

	c.Like()

	return nil
}

func (col *Colaborador) Dislike(c cancion.Cancion_info) error {
	if c.ExisteEn(col.CancionesOdiadas) {
		return errors.New("El usuario ya le ha dado dislike a esta canción")
	}

	c.Dislike()

	return nil
}

func (col *Colaborador) Recomendaciones() ([]cancion.Cancion, error) {
	return nil, nil
}
