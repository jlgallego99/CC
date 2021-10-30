package cancion_test

import (
	"github.com/jlgallego99/OSTfind/internal/cancion"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cancion", func() {
	var cancionCorrecta cancion.Cancion_info
	sensacionCorrecta := cancion.Tristeza

	sensacionesOrdenadas := []cancion.Sensacion{cancion.Alegria, cancion.Ansiedad, cancion.Ansiedad, cancion.Miedo, cancion.Miedo, cancion.Desafio}
	sensacionesDesordenadas := []cancion.Sensacion{cancion.Ansiedad, cancion.Miedo, cancion.Ansiedad, cancion.Alegria, cancion.Desafio, cancion.Miedo}

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
	})

	Describe("Añadir nueva sensación", func() {
		Context("La sensación es correcta", func() {
			It("Debe existir esa nueva sensación", func() {
				cancionCorrecta.NuevaSensacion(sensacionCorrecta)
				s := cancionCorrecta.Sensaciones[len(cancionCorrecta.Sensaciones)-1]

				Expect(s).To(Equal(sensacionCorrecta))
			})

			It("Debe devolver un error", func() {
				Expect(cancionCorrecta.NuevaSensacion(sensacionCorrecta)).NotTo(HaveOccurred())
			})
		})

		Context("La sensación no existe", func() {
			It("No debe existir ninguna nueva sensación", func() {
				cancionCorrecta.NuevaSensacion(100)
				length := len(cancionCorrecta.Sensaciones)

				Expect(length).To(Equal(0))
			})

			It("Debe devolver un error", func() {
				Expect(cancionCorrecta.NuevaSensacion(100)).To(HaveOccurred())
			})
		})
	})

	Describe("Calcular porcentaje de sensaciones", func() {
		Context("Las sensaciones están ordenadas", func() {
			var porcentajes []float64
			BeforeEach(func() {
				cancionCorrecta.Sensaciones = sensacionesOrdenadas
				porcentajes = cancionCorrecta.PorcentajeSensaciones()
			})

			It("Deben ser correctos todos los porcentajes", func() {
				Expect(porcentajes[cancion.Alegria]).To(Equal(16.67))
				Expect(porcentajes[cancion.Ansiedad]).To(Equal(33.33))
				Expect(porcentajes[cancion.Miedo]).To(Equal(33.33))
				Expect(porcentajes[cancion.Desafio]).To(Equal(16.67))
			})

			It("El porcentaje de las sensaciones que no aparecen debe ser 0", func() {
				Expect(porcentajes[cancion.Tristeza]).To(BeZero())
				Expect(porcentajes[cancion.Diversion]).To(BeZero())
				Expect(porcentajes[cancion.Energizante]).To(BeZero())
				Expect(porcentajes[cancion.Relajacion]).To(BeZero())
				Expect(porcentajes[cancion.Triunfo]).To(BeZero())
				Expect(porcentajes[cancion.Sueño]).To(BeZero())
				Expect(porcentajes[cancion.Epicidad]).To(BeZero())
			})
		})

		Context("Las sensaciones están desordenadas", func() {
			var porcentajes []float64
			BeforeEach(func() {
				cancionCorrecta.Sensaciones = sensacionesDesordenadas
				porcentajes = cancionCorrecta.PorcentajeSensaciones()
			})

			It("Deben ser correctos todos los porcentajes", func() {
				Expect(porcentajes[cancion.Alegria]).To(Equal(16.67))
				Expect(porcentajes[cancion.Ansiedad]).To(Equal(33.33))
				Expect(porcentajes[cancion.Miedo]).To(Equal(33.33))
				Expect(porcentajes[cancion.Desafio]).To(Equal(16.67))
			})

			It("El porcentaje de las sensaciones que no aparecen debe ser 0", func() {
				Expect(porcentajes[cancion.Tristeza]).To(BeZero())
				Expect(porcentajes[cancion.Diversion]).To(BeZero())
				Expect(porcentajes[cancion.Energizante]).To(BeZero())
				Expect(porcentajes[cancion.Relajacion]).To(BeZero())
				Expect(porcentajes[cancion.Triunfo]).To(BeZero())
				Expect(porcentajes[cancion.Sueño]).To(BeZero())
				Expect(porcentajes[cancion.Epicidad]).To(BeZero())
			})
		})

		Context("No hay ninguna sensación", func() {
			var porcentajes []float64
			BeforeEach(func() {
				cancionCorrecta.Sensaciones = make([]cancion.Sensacion, 0)
				porcentajes = cancionCorrecta.PorcentajeSensaciones()
			})

			It("El porcentaje de todas las sensaciones debe ser 0", func() {
				Expect(porcentajes[cancion.Alegria]).To(BeZero())
				Expect(porcentajes[cancion.Ansiedad]).To(BeZero())
				Expect(porcentajes[cancion.Miedo]).To(BeZero())
				Expect(porcentajes[cancion.Desafio]).To(BeZero())
				Expect(porcentajes[cancion.Tristeza]).To(BeZero())
				Expect(porcentajes[cancion.Diversion]).To(BeZero())
				Expect(porcentajes[cancion.Energizante]).To(BeZero())
				Expect(porcentajes[cancion.Relajacion]).To(BeZero())
				Expect(porcentajes[cancion.Triunfo]).To(BeZero())
				Expect(porcentajes[cancion.Sueño]).To(BeZero())
				Expect(porcentajes[cancion.Epicidad]).To(BeZero())
			})
		})
	})

	Describe("Se añaden likes o dislikes", func() {
		Context("No hay ningún like", func() {
			It("Debe registrar un nuevo like y tener uno", func() {
				cancionCorrecta.Like()

				Expect(cancionCorrecta.Likes).To(Equal(1))
			})

			It("Debe añadir y quitar el like y quedarse a 0", func() {
				cancionCorrecta.Like()
				cancionCorrecta.QuitarLike()

				Expect(cancionCorrecta.Likes).To(Equal(0))
			})

			It("No se pueden tener likes negativos", func() {
				cancionCorrecta.QuitarLike()

				Expect(cancionCorrecta.Likes).To(Equal(0))
			})
		})

		Context("No hay ningún dislike", func() {
			It("Debe registrar un nuevo dislike", func() {
				cancionCorrecta.Dislike()

				Expect(cancionCorrecta.Dislikes).To(Equal(1))
			})

			It("Debe añadir y quitar el dislike y quedarse a 0", func() {
				cancionCorrecta.Dislike()
				cancionCorrecta.QuitarDislike()

				Expect(cancionCorrecta.Dislikes).To(Equal(0))
			})

			It("No se pueden tener dislikes negativos", func() {
				cancionCorrecta.QuitarDislike()

				Expect(cancionCorrecta.Dislikes).To(Equal(0))
			})
		})
	})
})
