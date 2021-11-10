package usuario

import (
	"errors"
	"fmt"

	"github.com/jlgallego99/OSTfind/internal/cancion"
	"github.com/jlgallego99/OSTfind/internal/obra"
)

type Colaborador struct {
	Nombre               string
	CancionesFavoritas   []cancion.Cancion_info
	CancionesOdiadas     []cancion.Cancion_info
	CancionesColaboradas []cancion.Cancion_info
}

type Buscador struct {
	Nombre string
}

func NewColaborador(nombre string) (*Colaborador, error) {
	if nombre == "" {
		return &Colaborador{}, errors.New("nombre del usuario colaborador vacío")
	}

	return &Colaborador{
		Nombre:               nombre,
		CancionesFavoritas:   make([]cancion.Cancion_info, 0),
		CancionesOdiadas:     make([]cancion.Cancion_info, 0),
		CancionesColaboradas: make([]cancion.Cancion_info, 0),
	}, nil
}

func NewBuscador(nombre string) (*Buscador, error) {
	if nombre == "" {
		return &Buscador{}, errors.New("nombre del usuario buscador vacío")
	}

	return &Buscador{
		Nombre: nombre,
	}, nil
}

type Usuario interface {
	Like(c *cancion.Cancion) error
	Dislike(c *cancion.Cancion) error
	Recomendaciones() ([]cancion.Cancion, error)
	ActualizarSensaciones(c *cancion.Cancion_info, sensaciones []cancion.Sensacion) error
	CrearSerie(titulo string, temporada int, capitulo int, canciones []cancion.Cancion_info) (obra.Serie, error)
	CrearPelicula(titulo string, canciones []cancion.Cancion_info) (obra.Pelicula, error)
	CrearVideojuego(titulo string, canciones []cancion.Cancion_info) (obra.Videojuego, error)
}

func (col *Colaborador) Like(c *cancion.Cancion_info) error {
	if existe, _ := c.ExisteEn(col.CancionesFavoritas); existe {
		return errors.New("el usuario ya le ha dado like a esta canción")
	}

	if existe, i := c.ExisteEn(col.CancionesOdiadas); existe {
		col.CancionesOdiadas = append(col.CancionesOdiadas[:i], col.CancionesOdiadas[i+1:]...)
	}

	c.Like()
	col.CancionesFavoritas = append(col.CancionesFavoritas, *c)

	return nil
}

func (col *Colaborador) Dislike(c *cancion.Cancion_info) error {
	if existe, _ := c.ExisteEn(col.CancionesOdiadas); existe {
		return errors.New("el usuario ya le ha dado dislike a esta canción")
	}

	if existe, i := c.ExisteEn(col.CancionesFavoritas); existe {
		col.CancionesFavoritas = append(col.CancionesFavoritas[:i], col.CancionesFavoritas[i+1:]...)
	}

	c.Dislike()
	col.CancionesOdiadas = append(col.CancionesOdiadas, *c)

	return nil
}

func (col *Colaborador) Recomendaciones() ([]cancion.Cancion, error) {
	return nil, nil
}

func (col *Colaborador) ActualizarSensaciones(c *cancion.Cancion_info, sensaciones []cancion.Sensacion) error {
	if c == nil {
		return errors.New("no existe la canción")
	}

	sensacionesUsuario := make([]cancion.Sensacion, len(c.Sensaciones))
	copy(sensacionesUsuario, c.Sensaciones)

	// Buscar la canción colaborada
	var pos int
	for i, v := range col.CancionesColaboradas {
		if v.Titulo == c.Titulo {
			pos = i
		}
	}

	// Quitar solo las sensaciones que ha aportado el usuario
	if len(sensaciones) == 0 {
		col.CancionesColaboradas = append(col.CancionesColaboradas[:pos], col.CancionesColaboradas[pos+1:]...)

		for _, s := range sensacionesUsuario {
			err := c.QuitarSensacion(s)

			if err != nil {
				return fmt.Errorf("no se ha podido eliminar la sensación repetida: %s", err)
			}
		}

		return nil
	}

	// El usuario ya ha aportado sensaciones antes, se actualizan con las nuevas
	var existe bool
	if existe, _ = c.ExisteEn(col.CancionesColaboradas); existe {
		col.CancionesColaboradas[pos].Sensaciones = sensaciones

		for _, s_us := range sensacionesUsuario {
			for _, s := range c.Sensaciones {
				if s == s_us {
					err := c.QuitarSensacion(s)

					if err != nil {
						return fmt.Errorf("no se ha podido eliminar la sensación repetida: %s", err)
					}
				}
			}
		}
	}

	// Añadir las nuevas sensaciones a la canción
	for _, s := range sensaciones {
		err := c.NuevaSensacion(s)

		if err != nil {
			return fmt.Errorf("no se ha podido registrar la nueva sensación: %s", err)
		}
	}

	if !existe {
		col.CancionesColaboradas = append(col.CancionesColaboradas, *c)
	}

	return nil
}

func (col *Colaborador) CrearSerie(titulo string, temporada int, capitulo int, canciones []cancion.Cancion_info) (obra.Serie, error) {
	return obra.NewSerie(titulo, temporada, capitulo)
}

func (col *Colaborador) CrearPelicula(titulo string, canciones []cancion.Cancion_info) (obra.Pelicula, error) {
	return obra.NewPelicula(titulo)
}

func (col *Colaborador) CrearVideojuego(titulo string, canciones []cancion.Cancion_info) (obra.Videojuego, error) {
	return obra.NewVideojuego(titulo)
}
