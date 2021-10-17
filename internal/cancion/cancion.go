package cancion

type Genero int

const (
	Rock Genero = iota
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
	Batalla Momento = iota
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

type Cancion interface {
	porcentajeLikeDislike() (float64, float64)
	porcentajeSensaciones() ([]float64, []Sensacion)
	nuevaSensacion(s Sensacion)
	cancionesRelacionadas() []Cancion_info
}
