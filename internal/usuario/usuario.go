package usuario

import (
	"errors"
	"fmt"

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
	CrearSerie(titulo string, temporada int, capitulo int, canciones []cancion.Cancion_info) (*obra.Serie, error)
	CrearPelicula(titulo string, canciones []cancion.Cancion_info) (*obra.Pelicula, error)
	CrearVideojuego(titulo string, canciones []cancion.Cancion_info) (*obra.Videojuego, error)
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
		err := o.NuevaCancion(v)

		if err != nil {
			return fmt.Errorf("No se ha podido añadir la canción: %s", err)
		}
	}

	return nil
}

func (col *Colaborador) CrearSerie(titulo string, temporada int, capitulo int, canciones []cancion.Cancion_info) (*obra.Serie, error) {
	return obra.NewSerie(titulo, temporada, capitulo, canciones)
}

func (col *Colaborador) CrearPelicula(titulo string, canciones []cancion.Cancion_info) (*obra.Pelicula, error) {
	return obra.NewPelicula(titulo, canciones)
}

func (col *Colaborador) CrearVideojuego(titulo string, canciones []cancion.Cancion_info) (*obra.Videojuego, error) {
	return obra.NewVideojuego(titulo, canciones)
}
