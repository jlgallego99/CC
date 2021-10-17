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
