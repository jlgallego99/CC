package cancion

import "errors"

type Genero int

const (
	Genero_Desconocido Genero = iota
	Rock
	Pop
	Ambiental
	Electronica
	Funk
	Jazz
	Orquesta
	Vocal
)

type Sensacion int

const (
	Alegria Sensacion = iota
	Tristeza
	Epicidad
)

type Momento int

const (
	Momento_Desconocido Momento = iota
	Batalla
	Ciudad
	Evento
	Zona
	Creditos
	Opening
	Personaje
	Minijuego
)

type Cancion_info struct {
	Titulo          string
	Compositor      string
	Genero          Genero
	Likes           int
	Dislikes        int
	Sensaciones     []Sensacion
	Momento         Momento
	Momento_exacto  string
	Momento_minutos string
}

func (s *Sensacion) valid() error {
	switch *s {
	case Alegria, Tristeza, Epicidad:
		return nil
	default:
		return errors.New("Sensación no válida")
	}
}

func newCancionInfo(titulo string, compositor string, genero Genero, momento Momento, momento_minutos string) {
}

type Cancion interface {
	porcentajeLikeDislike() (float64, float64)
	porcentajeSensaciones() ([]float64, []Sensacion)
	nuevaSensacion(s Sensacion)
	cancionesRelacionadas(num int) []Cancion_info
}

func (c *Cancion_info) porcentajeLikeDislike() (float64, float64) {
	return 0, 0
}

func (c *Cancion_info) porcentajeSensaciones() ([]float64, []Sensacion) {
	return nil, nil
}

func (c *Cancion_info) nuevaSensacion(s Sensacion) error {
	if err := s.valid(); err != nil {
		c.Sensaciones = append(c.Sensaciones, s)
		return nil
	} else {
		return err
	}
}

func (c *Cancion_info) cancionesRelacionadas(num int) []Cancion_info {
	return nil
}
