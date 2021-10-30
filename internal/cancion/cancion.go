package cancion

import (
	"errors"
	"math"
)

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
	Ansiedad
	Diversion
	Energizante
	Miedo
	Relajacion
	Triunfo
	Sueño
	Epicidad
	Desafio
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

func (s *Sensacion) Valid() error {
	switch *s {
	case Alegria, Tristeza, Ansiedad, Diversion, Energizante, Miedo, Relajacion,
		Triunfo, Sueño, Epicidad, Desafio:
		return nil
	default:
		return errors.New("Sensación no válida")
	}
}

func NewCancionInfo(titulo string, compositor string, genero Genero, momento Momento, momento_minutos string) {
}

type Cancion interface {
	PorcentajeLikeDislike() (float64, float64)
	PorcentajeSensaciones() ([]float64, []Sensacion)
	NuevaSensacion(s Sensacion)
	CancionesRelacionadas(num int) []Cancion_info
}

func (c *Cancion_info) PorcentajeLikeDislike() (float64, float64) {
	return 0, 0
}

func (c *Cancion_info) PorcentajeSensaciones() []float64 {
	ocurrencias := make([]int, 11)
	porcentajes := make([]float64, 0)

	for _, v := range c.Sensaciones {
		switch v {
		case Alegria:
			ocurrencias[0]++
		case Tristeza:
			ocurrencias[1]++
		case Ansiedad:
			ocurrencias[2]++
		case Diversion:
			ocurrencias[3]++
		case Energizante:
			ocurrencias[4]++
		case Miedo:
			ocurrencias[5]++
		case Relajacion:
			ocurrencias[6]++
		case Triunfo:
			ocurrencias[7]++
		case Sueño:
			ocurrencias[8]++
		case Epicidad:
			ocurrencias[9]++
		case Desafio:
			ocurrencias[10]++
		}
	}

	for _, v := range ocurrencias {
		p := float64(v) / float64(len(c.Sensaciones)) * 100
		p = math.Round(p*100) / 100
		porcentajes = append(porcentajes, p)
	}

	return porcentajes
}

func (c *Cancion_info) NuevaSensacion(s Sensacion) error {
	err := s.Valid()

	if err == nil {
		c.Sensaciones = append(c.Sensaciones, s)
	}

	return err
}

func (c *Cancion_info) CancionesRelacionadas(num int) []Cancion_info {
	return nil
}
