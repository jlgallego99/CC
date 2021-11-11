package cancion

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/google/uuid"
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
	Id             uuid.UUID
	Titulo         string
	Compositor     string
	Genero         Genero
	Likes          int
	Dislikes       int
	Sensaciones    []Sensacion
	Momento        Momento
	Momento_exacto string
}

func (s *Sensacion) Valid() error {
	switch *s {
	case Alegria, Tristeza, Ansiedad, Diversion, Energizante, Miedo, Relajacion,
		Triunfo, Sueño, Epicidad, Desafio:
		return nil
	default:
		return errors.New("sensación no válida")
	}
}

func (g *Genero) Valid() error {
	switch *g {
	case Genero_Desconocido, Rock, Pop, Ambiental, Electronica, Funk, Jazz, Orquesta, Vocal:
		return nil
	default:
		return errors.New("género no válido")
	}
}

func NewCancion(titulo string, compositor string, genero Genero) (*Cancion_info, error) {
	if titulo == "" {
		return &Cancion_info{}, errors.New("título de la canción vacío")
	}

	if compositor == "" {
		return &Cancion_info{}, errors.New("compositor vacío")
	}

	err := genero.Valid()

	if err != nil {
		return &Cancion_info{}, fmt.Errorf("error al añadir el género: %s", err)
	}

	return &Cancion_info{
		Id:             uuid.New(),
		Titulo:         titulo,
		Compositor:     compositor,
		Genero:         genero,
		Likes:          0,
		Dislikes:       0,
		Sensaciones:    make([]Sensacion, 0),
		Momento:        Momento_Desconocido,
		Momento_exacto: "",
	}, nil
}

type Cancion interface {
	PorcentajeLikeDislike() (float64, float64)
	PorcentajeSensaciones() ([]float64, []Sensacion)
	NuevaSensacion(s Sensacion) error
	CancionesRelacionadas(num int) []Cancion_info
	Like()
	Dislike()
	QuitarLike()
	QuitarDislike()
	ExisteEn(canciones []*Cancion_info) bool
}

func (c *Cancion_info) PorcentajeLikeDislike() (float64, float64) {
	var likes float64
	var dislikes float64

	total := c.Likes + c.Dislikes
	if total == 0 {
		likes = 0
		dislikes = 0
	} else {
		likes = float64(c.Likes) / float64(total) * 100.0
		likes = math.Round(likes*100) / 100
		dislikes = float64(c.Dislikes) / float64(total) * 100.0
		dislikes = math.Round(dislikes*100) / 100
	}

	return likes, dislikes
}

func (c *Cancion_info) PorcentajeSensaciones() []float64 {
	ocurrencias := make([]int, 11)
	porcentajes := make([]float64, 11)
	var p float64

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

	for i, v := range ocurrencias {
		if v == 0 {
			p = 0
		} else {
			p = float64(v) / float64(len(c.Sensaciones)) * 100
			p = math.Round(p*100) / 100
		}

		porcentajes[i] = p
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

func (c *Cancion_info) QuitarSensacion(s Sensacion) error {
	err := s.Valid()

	if err == nil {
		for i, v := range c.Sensaciones {
			if reflect.DeepEqual(s, v) {
				c.Sensaciones = append(c.Sensaciones[:i], c.Sensaciones[i+1:]...)

				return nil
			}
		}
	}

	return err
}

func (c *Cancion_info) CancionesRelacionadas(num int) []Cancion_info {
	return nil
}

func (c *Cancion_info) Like() {
	c.Likes++
}

func (c *Cancion_info) Dislike() {
	c.Dislikes++
}

func (c *Cancion_info) QuitarLike() {
	c.Likes--

	if c.Likes < 0 {
		c.Likes = 0
	}
}

func (c *Cancion_info) QuitarDislike() {
	c.Dislikes--

	if c.Dislikes < 0 {
		c.Dislikes = 0
	}
}

func (c *Cancion_info) ExisteEn(canciones []*Cancion_info) (bool, int) {
	for i, v := range canciones {
		if c == v {
			return true, i
		}
	}

	return false, -1
}
