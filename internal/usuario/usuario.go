package usuario

import "github.com/jlgallego99/OSTfind/internal/cancion"

type Usuario interface {
	Like(c cancion.Cancion) error
	Dislike(c cancion.Cancion) error
	Recomendaciones() ([]cancion.Cancion, error)
}
