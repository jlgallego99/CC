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
				porcentajes, _ = cancionCorrecta.PorcentajeSensaciones()
			})

			It("Deben ser correctos todos los porcentajes", func() {
				Expect(porcentajes[0]).To(Equal(16.67))
				Expect(porcentajes[2]).To(Equal(33.33))
				Expect(porcentajes[5]).To(Equal(33.33))
				Expect(porcentajes[10]).To(Equal(16.67))
			})

			It("El porcentaje de las sensaciones que no aparecen debe ser 0", func() {
				Expect(porcentajes[1]).To(BeZero())
				Expect(porcentajes[3]).To(BeZero())
				Expect(porcentajes[4]).To(BeZero())
				Expect(porcentajes[6]).To(BeZero())
				Expect(porcentajes[7]).To(BeZero())
				Expect(porcentajes[8]).To(BeZero())
				Expect(porcentajes[9]).To(BeZero())
			})
		})

		Context("Las sensaciones están desordenadas", func() {
			var porcentajes []float64
			BeforeEach(func() {
				cancionCorrecta.Sensaciones = sensacionesDesordenadas
				porcentajes, _ = cancionCorrecta.PorcentajeSensaciones()
			})

			It("Deben ser correctos todos los porcentajes", func() {
				Expect(porcentajes[0]).To(Equal(16.67))
				Expect(porcentajes[2]).To(Equal(33.33))
				Expect(porcentajes[5]).To(Equal(33.33))
				Expect(porcentajes[10]).To(Equal(16.67))
			})

			It("El porcentaje de las sensaciones que no aparecen debe ser 0", func() {
				Expect(porcentajes[1]).To(BeZero())
				Expect(porcentajes[3]).To(BeZero())
				Expect(porcentajes[4]).To(BeZero())
				Expect(porcentajes[6]).To(BeZero())
				Expect(porcentajes[7]).To(BeZero())
				Expect(porcentajes[8]).To(BeZero())
				Expect(porcentajes[9]).To(BeZero())
			})
		})
	})
})
