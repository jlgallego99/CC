package cancion

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
	titulo          string
	compositor      string
	genero          Genero
	likes           int
	dislikes        int
	sensacion       []Sensacion
	momento         Momento
	momento_exacto  string
	momento_minutos string
}

func newCancionInfo(titulo string, compositor string, genero Genero, momento Momento, momento_minutos string) {
}

type Cancion interface {
	porcentajeLikeDislike() (float64, float64)
	porcentajeSensaciones() ([]float64, []Sensacion)
	nuevaSensacion(s Sensacion)
	cancionesRelacionadas() []Cancion_info
}

func (c *Cancion_info) porcentajeLikeDislike() (float64, float64) {
	return 0, 0
}

func (c *Cancion_info) porcentajeSensaciones() ([]float64, []Sensacion) {
	return nil, nil
}

func (c *Cancion_info) nuevaSensacion(s Sensacion) {

}

func (c *Cancion_info) cancionesRelacionadas() []Cancion_info {
	return nil
}
