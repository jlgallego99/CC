package obra_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlgallego99/OSTfind/internal/cancion"
	"github.com/jlgallego99/OSTfind/internal/obra"
)

var _ = Describe("Obra", func() {
	var serie *obra.Serie
	var pelicula *obra.Pelicula
	var videojuego *obra.Videojuego
	var cancionCorrecta cancion.Cancion_info

	BeforeEach(func() {
		cancionCorrecta = cancion.Cancion_info{
			Titulo:          "a",
			Compositor:      "b",
			Genero:          cancion.Ambiental,
			Likes:           0,
			Dislikes:        0,
			Sensaciones:     make([]cancion.Sensacion, 0),
			Momento:         cancion.Ciudad,
			Momento_exacto:  "",
			Momento_minutos: "",
		}

		serie = obra.NewSerie("SeriePrueba", 1, 1, []cancion.Cancion_info{cancionCorrecta, {}})
		pelicula = obra.NewPelicula("PeliculaPrueba", []cancion.Cancion_info{cancionCorrecta, {}})
		videojuego = obra.NewVideojuego("VideojuegoPrueba", []cancion.Cancion_info{cancionCorrecta, {}})
	})

	Describe("Crear nueva obra", func() {
		Context("Se crea una serie", func() {
			It("Debe tener todos los campos iguales", func() {
				Expect(serie.Titulo).To(Equal("SeriePrueba"))
				Expect(serie.Temporada).To(Equal(1))
				Expect(serie.Capitulo).To(Equal(1))
				Expect(serie.OST).To(Equal([]cancion.Cancion_info{cancionCorrecta, {}}))
			})
		})

		Context("Se crea una película", func() {
			It("Debe tener todos los campos iguales", func() {
				Expect(pelicula.Titulo).To(Equal("PeliculaPrueba"))
				Expect(pelicula.OST).To(Equal([]cancion.Cancion_info{cancionCorrecta, {}}))
			})
		})

		Context("Se crea una serie", func() {
			It("Debe tener todos los campos iguales", func() {
				Expect(videojuego.Titulo).To(Equal("VideojuegoPrueba"))
				Expect(videojuego.OST).To(Equal([]cancion.Cancion_info{cancionCorrecta, {}}))
			})
		})
	})
})
