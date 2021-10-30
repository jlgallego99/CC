package usuario

import (
	"errors"

	"github.com/jlgallego99/OSTfind/internal/cancion"
	"github.com/jlgallego99/OSTfind/internal/obra"
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
	ActualizarOST(o obra.Obra, ost []cancion.Cancion_info) error
}

func (col *Colaborador) Like(c cancion.Cancion_info) error {
	if existe, _ := c.ExisteEn(col.CancionesFavoritas); existe {
		return errors.New("El usuario ya le ha dado like a esta canción")
	}

	if existe, i := c.ExisteEn(col.CancionesOdiadas); existe {
		col.CancionesOdiadas = append(col.CancionesOdiadas[:i], col.CancionesOdiadas[i+1:]...)
	}

	col.CancionesFavoritas = append(col.CancionesFavoritas, c)
	c.Like()

	return nil
}

func (col *Colaborador) Dislike(c cancion.Cancion_info) error {
	if existe, _ := c.ExisteEn(col.CancionesOdiadas); existe {
		return errors.New("El usuario ya le ha dado dislike a esta canción")
	}

	if existe, i := c.ExisteEn(col.CancionesFavoritas); existe {
		col.CancionesFavoritas = append(col.CancionesFavoritas[:i], col.CancionesFavoritas[i+1:]...)
	}

	col.CancionesOdiadas = append(col.CancionesOdiadas, c)
	c.Dislike()

	return nil
}

func (col *Colaborador) Recomendaciones() ([]cancion.Cancion, error) {
	return nil, nil
}

func (col *Colaborador) ActualizarOST(o obra.Obra, ost []cancion.Cancion_info) error {
	if o == nil {
		return errors.New("No existe la obra")
	}

	for _, v := range ost {
		o.NuevaCancion(v)
	}

	return nil
}
