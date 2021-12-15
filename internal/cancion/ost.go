package cancion

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/jlgallego99/OSTfind/internal/obra"
)

type BandaSonora struct {
	Id        string
	Canciones []*Cancion_info
	Obra      obra.Obra
}

type OST interface {
	ActualizarOST(ost []*Cancion_info) error
	ActualizarObra(titulo string, captemp ...int) error
	NuevaCancion(c *Cancion_info) error
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
		Id:        uuid.New().String(),
		Obra:      v,
		Canciones: canciones,
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
		Id:        uuid.New().String(),
		Obra:      p,
		Canciones: canciones,
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
		Id:        uuid.New().String(),
		Obra:      p,
		Canciones: canciones,
	}, nil
}

func (b *BandaSonora) ActualizarOST(ost []*Cancion_info) error {
	ostAntigua := b.Canciones
	b.Canciones = make([]*Cancion_info, 0)

	for _, v := range ost {
		err := b.NuevaCancion(v)

		if err != nil {
			b.Canciones = ostAntigua
			return fmt.Errorf("no se ha podido añadir la canción: %s", err)
		}
	}

	return nil
}

func (b *BandaSonora) ActualizarObra(titulo string, captemp ...int) error {
	var err error
	var o obra.Obra

	switch b.Obra.(type) {
	case obra.Videojuego:
		o, err = obra.NewVideojuego(titulo)

	case obra.Pelicula:
		o, err = obra.NewPelicula(titulo)

	case obra.Serie:
		if len(captemp) != 2 {
			err = errors.New("no se ha especificado tanto capítulo como temporada")
		} else {
			o, err = obra.NewSerie(titulo, captemp[0], captemp[1])
		}
	}

	if err != nil {
		return fmt.Errorf("no se ha podido actualizar la obra: %s", err)
	} else {
		b.Obra = o
	}

	return nil
}

func (b *BandaSonora) NuevaCancion(c *Cancion_info) error {
	if existe, _ := c.ExisteEn(b.Canciones); existe {
		return errors.New("la canción ya existe en la OST")
	}

	if reflect.DeepEqual(Cancion_info{}, *c) {
		return errors.New("la canción está vacía")
	}

	b.Canciones = append(b.Canciones, c)

	return nil
}

func (b *BandaSonora) Cancion(titulo string) (*Cancion_info, error) {
	for _, v := range b.Canciones {
		if v.Titulo == titulo {
			return v, nil
		}
	}

	return nil, errors.New("no se ha encontrado la canción especificada")
}
