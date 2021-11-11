package cancion

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jlgallego99/OSTfind/internal/obra"
)

type BandaSonora struct {
	canciones []*Cancion_info
	obra      obra.Obra
}

type OST interface {
	ActualizarOST(ost []Cancion_info) error
	NuevaCancion(c Cancion_info) error
	Cancion(titulo string) (*Cancion_info, error)
}

func NewVideojuegoOST(titulo string, canciones []*Cancion_info) (*BandaSonora, error) {
	if canciones == nil {
		return &BandaSonora{}, errors.New("la lista de canciones es nula")
	}

	v, err := obra.NewVideojuego(titulo)

	if err != nil {
		return &BandaSonora{}, fmt.Errorf("no se ha podido crear la ost de videojuego: %s", err)
	}

	return &BandaSonora{
		obra:      v,
		canciones: canciones,
	}, nil
}

func NewPeliculaOST(titulo string, canciones []*Cancion_info) (*BandaSonora, error) {
	if canciones == nil {
		return &BandaSonora{}, errors.New("la lista de canciones es nula")
	}

	p, err := obra.NewPelicula(titulo)

	if err != nil {
		return &BandaSonora{}, fmt.Errorf("no se ha podido crear la ost de película: %s", err)
	}

	return &BandaSonora{
		obra:      p,
		canciones: canciones,
	}, nil
}

func NewSerieOST(titulo string, temporada, capitulo int, canciones []*Cancion_info) (*BandaSonora, error) {
	if canciones == nil {
		return &BandaSonora{}, errors.New("la lista de canciones es nula")
	}

	p, err := obra.NewSerie(titulo, temporada, capitulo)

	if err != nil {
		return &BandaSonora{}, fmt.Errorf("no se ha podido crear la ost de serie: %s", err)
	}

	return &BandaSonora{
		obra:      p,
		canciones: canciones,
	}, nil
}

func (b *BandaSonora) ActualizarOST(ost []Cancion_info) error {
	b.canciones = make([]*Cancion_info, 0)

	for _, v := range ost {
		err := b.NuevaCancion(v)

		if err != nil {
			return fmt.Errorf("no se ha podido añadir la canción: %s", err)
		}
	}

	return nil
}

func (b *BandaSonora) NuevaCancion(c Cancion_info) error {
	if existe, _ := c.ExisteEn(b.canciones); existe {
		return errors.New("la canción ya existe en la OST")
	}

	if reflect.DeepEqual(Cancion_info{}, c) {
		return errors.New("la canción está vacía")
	}

	b.canciones = append(b.canciones, &c)

	return nil
}

func (b *BandaSonora) Cancion(titulo string) (*Cancion_info, error) {
	for _, v := range b.canciones {
		if v.Titulo == titulo {
			return v, nil
		}
	}

	return nil, errors.New("no se ha encontrado la canción especificada")
}
