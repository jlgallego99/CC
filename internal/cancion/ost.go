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
}
