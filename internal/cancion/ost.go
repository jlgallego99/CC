package cancion

import (
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
