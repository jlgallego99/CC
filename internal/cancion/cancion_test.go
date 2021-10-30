package cancion_test

import (
	"github.com/jlgallego99/OSTfind/internal/cancion"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cancion", func() {
	var cancionCorrecta cancion.Cancion_info
	var sensacionCorrecta cancion.Sensacion

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
		sensacionCorrecta = cancion.Tristeza
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
})
