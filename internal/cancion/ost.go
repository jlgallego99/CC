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
	Cancion(titulo string) error
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
