package usuario_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlgallego99/OSTfind/internal/cancion"
	"github.com/jlgallego99/OSTfind/internal/obra"
	"github.com/jlgallego99/OSTfind/internal/usuario"
)

var _ = Describe("Usuario", func() {
	var colaborador usuario.Colaborador
	var cancionCorrecta cancion.Cancion_info

	var serie *obra.Serie
	var pelicula *obra.Pelicula
	var videojuego *obra.Videojuego

	var cancionesVacio []cancion.Cancion_info
	var canciones []cancion.Cancion_info

	var err_s, err_p, err_v error

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

		serie = obra.NewSerie("SeriePrueba", 1, 1, make([]cancion.Cancion_info, 0))
		pelicula = obra.NewPelicula("PeliculaPrueba", make([]cancion.Cancion_info, 0))
		videojuego = obra.NewVideojuego("VideojuegoPrueba", make([]cancion.Cancion_info, 0))
		cancionesVacio = make([]cancion.Cancion_info, 5)
		canciones = make([]cancion.Cancion_info, 0)
		canciones = append(canciones, cancionCorrecta)
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

	Describe("Actualizar la OST de una obra", func() {
		BeforeEach(func() {
			err_s = colaborador.ActualizarOST(serie, canciones)
			err_p = colaborador.ActualizarOST(pelicula, canciones)
			err_v = colaborador.ActualizarOST(videojuego, canciones)
		})

		Context("La obra no existe", func() {
			It("Debe dar error", func() {
				err := colaborador.ActualizarOST(nil, cancionesVacio)

				Expect(err).To(HaveOccurred())
			})
		})

		Context("La obra existe", func() {
			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})
		})

		Context("La nueva OST está vacía", func() {
			It("Debe dar error", func() {
				err_s = colaborador.ActualizarOST(serie, cancionesVacio)
				err_p = colaborador.ActualizarOST(pelicula, cancionesVacio)
				err_v = colaborador.ActualizarOST(videojuego, cancionesVacio)

				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})

		Context("Se añaden canciones que ya existen en la OST", func() {
			It("Debe dar error", func() {
				err_s = colaborador.ActualizarOST(serie, canciones)
				err_p = colaborador.ActualizarOST(pelicula, canciones)
				err_v = colaborador.ActualizarOST(videojuego, canciones)

				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})
	})
})
