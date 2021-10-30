package usuario_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlgallego99/OSTfind/internal/cancion"
	"github.com/jlgallego99/OSTfind/internal/usuario"
)

var _ = Describe("Usuario", func() {
	var colaborador usuario.Colaborador
	var cancionCorrecta cancion.Cancion_info

	BeforeEach(func() {
		colaborador = usuario.Colaborador{
			Nombre:             "PepeColabora",
			CancionesFavoritas: make([]cancion.Cancion_info, 0),
			CancionesOdiadas:   make([]cancion.Cancion_info, 0),
		}

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

	Describe("Dar like o dislike a una canción", func() {
		Context("No se le he dado aún el like a la canción", func() {
			It("No debe tener error", func() {
				err := colaborador.Like(cancionCorrecta)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("Se le ha dado ya el like a la canción", func() {
			BeforeEach(func() {
				colaborador.Like(cancionCorrecta)
			})

			It("Debe tener error", func() {
				err := colaborador.Like(cancionCorrecta)
				Expect(err).To(HaveOccurred())
			})

			It("Debe tener la canción en su lista de canciones favoritas", func() {
				Expect(colaborador.CancionesFavoritas[0]).To(Equal(cancionCorrecta))
			})

			It("Si se le da dislike se debe quitar el like", func() {
				colaborador.Dislike(cancionCorrecta)
				Expect(colaborador.CancionesFavoritas).To(BeEmpty())
			})
		})

		Context("Se le ha dado ya el dislike a la canción", func() {
			BeforeEach(func() {
				colaborador.Dislike(cancionCorrecta)
			})

			It("Debe tener error", func() {
				err := colaborador.Dislike(cancionCorrecta)
				Expect(err).To(HaveOccurred())
			})

			It("Debe tener la canción en su lista de canciones favoritas", func() {
				Expect(colaborador.CancionesOdiadas[0]).To(Equal(cancionCorrecta))
			})

			It("Si se le da like se debe quitar el dislike", func() {
				colaborador.Like(cancionCorrecta)
				Expect(colaborador.CancionesOdiadas).To(BeEmpty())
			})
		})
	})
})
