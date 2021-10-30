package usuario

import (
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

type Colaborador struct {
	Nombre             string
	CancionesFavoritas []cancion.Cancion
}

type Buscador struct {
	Nombre string
}

type Usuario interface {
	Like(c cancion.Cancion) error
	Dislike(c cancion.Cancion) error
	Recomendaciones() ([]cancion.Cancion, error)
}

func (col *Colaborador) Like(c cancion.Cancion) {
	c.Like()
}

func (col *Colaborador) Dislike(c cancion.Cancion) {
	c.Dislike()
}

func (col *Colaborador) Recomendaciones() ([]cancion.Cancion, error) {
	return nil, nil
}
