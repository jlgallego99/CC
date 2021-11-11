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

	var serie obra.Serie
	var pelicula obra.Pelicula
	var videojuego obra.Videojuego

	var err error
	var err_ns, err_np, err_nv error

	var sensaciones []cancion.Sensacion
	var sensacionesNuevas []cancion.Sensacion

	BeforeEach(func() {
		colaborador = usuario.Colaborador{
			Nombre:             "PepeColabora",
			CancionesFavoritas: make([]*cancion.Cancion_info, 0),
			CancionesOdiadas:   make([]*cancion.Cancion_info, 0),
		}

		cancionCorrecta = cancion.Cancion_info{
			Titulo:         "a",
			Compositor:     "b",
			Genero:         cancion.Ambiental,
			Likes:          0,
			Dislikes:       0,
			Sensaciones:    make([]cancion.Sensacion, 0),
			Momento:        cancion.Ciudad,
			Momento_exacto: "",
		}

		serie, err_ns = colaborador.CrearSerie("SeriePrueba", 1, 1, make([]cancion.Cancion_info, 0))
		pelicula, err_np = colaborador.CrearPelicula("PeliculaPrueba", make([]cancion.Cancion_info, 0))
		videojuego, err_nv = colaborador.CrearVideojuego("VideojuegoPrueba", make([]cancion.Cancion_info, 0))

		sensaciones = []cancion.Sensacion{cancion.Alegria, cancion.Ansiedad, cancion.Ansiedad, cancion.Miedo, cancion.Miedo, cancion.Desafio, cancion.Tristeza}
		sensacionesNuevas = []cancion.Sensacion{cancion.Alegria, cancion.Ansiedad, cancion.Desafio, cancion.Sueño}
	})

	Describe("Dar like o dislike a una canción", func() {
		Context("No se le he dado aún el like a la canción", func() {
			BeforeEach(func() {
				err = colaborador.Like(&cancionCorrecta)
			})

			It("La canción debe tener exactamente un like", func() {
				Expect(cancionCorrecta.Likes).To(Equal(1))
			})

			It("No debe tener error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("No se le he dado aún el dislike a la canción", func() {
			BeforeEach(func() {
				err = colaborador.Dislike(&cancionCorrecta)
			})

			It("La canción debe tener exactamente un dislike", func() {
				Expect(cancionCorrecta.Dislikes).To(Equal(1))
			})

			It("No debe tener error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("Se le ha dado ya el like a la canción", func() {
			BeforeEach(func() {
				colaborador.Like(&cancionCorrecta)
			})

			It("Debe tener error", func() {
				err := colaborador.Like(&cancionCorrecta)
				Expect(err).To(HaveOccurred())
			})

			It("Debe tener la canción en su lista de canciones favoritas", func() {
				Expect(colaborador.CancionesFavoritas[0]).To(Equal(&cancionCorrecta))
			})

			It("Si se le da dislike se debe quitar el like", func() {
				colaborador.Dislike(&cancionCorrecta)
				Expect(colaborador.CancionesFavoritas).To(BeEmpty())
			})
		})

		Context("Se le ha dado ya el dislike a la canción", func() {
			BeforeEach(func() {
				colaborador.Dislike(&cancionCorrecta)
			})

			It("Debe tener error", func() {
				err := colaborador.Dislike(&cancionCorrecta)
				Expect(err).To(HaveOccurred())
			})

			It("Debe tener la canción en su lista de canciones favoritas", func() {
				Expect(colaborador.CancionesOdiadas[0]).To(Equal(&cancionCorrecta))
			})

			It("Si se le da like se debe quitar el dislike", func() {
				colaborador.Like(&cancionCorrecta)
				Expect(colaborador.CancionesOdiadas).To(BeEmpty())
			})
		})
	})

	Describe("Actualizar las sensaciones de una obra", func() {
		BeforeEach(func() {
			err = colaborador.ActualizarSensaciones(&cancionCorrecta, sensaciones)
		})

		Context("La canción no existe", func() {
			BeforeEach(func() {
				err = colaborador.ActualizarSensaciones(nil, sensaciones)
			})

			It("Debe dar error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Se quitan todas las sensaciones aportadas", func() {
			BeforeEach(func() {
				err = colaborador.ActualizarSensaciones(&cancionCorrecta, []cancion.Sensacion{})
			})

			It("No debe estar la canción en la lista de colaboradas", func() {
				Expect(len(cancionCorrecta.Sensaciones)).To(Equal(0))
			})

			It("No debe dar error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("No se había aportado anteriormente a la canción", func() {
			It("Deben haberse añadido todas las sensaciones", func() {
				Expect(len(cancionCorrecta.Sensaciones)).To(Equal(7))
				Expect(cancionCorrecta.Sensaciones[0]).To(Equal(cancion.Alegria))
				Expect(cancionCorrecta.Sensaciones[1]).To(Equal(cancion.Ansiedad))
				Expect(cancionCorrecta.Sensaciones[2]).To(Equal(cancion.Ansiedad))
				Expect(cancionCorrecta.Sensaciones[3]).To(Equal(cancion.Miedo))
				Expect(cancionCorrecta.Sensaciones[4]).To(Equal(cancion.Miedo))
				Expect(cancionCorrecta.Sensaciones[5]).To(Equal(cancion.Desafio))
				Expect(cancionCorrecta.Sensaciones[6]).To(Equal(cancion.Tristeza))
			})

			It("La canción está en la lista de colaboradas del usuario", func() {
				Expect(colaborador.CancionesColaboradas[0]).To(Equal(&cancionCorrecta))
			})

			It("No debe dar error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("Se había aportado anteriormente a la canción", func() {
			BeforeEach(func() {
				err = colaborador.ActualizarSensaciones(&cancionCorrecta, sensacionesNuevas)
			})

			It("Deben haberse añadido todas las sensaciones", func() {
				Expect(len(cancionCorrecta.Sensaciones)).To(Equal(4))
				Expect(cancionCorrecta.Sensaciones[0]).To(Equal(cancion.Alegria))
				Expect(cancionCorrecta.Sensaciones[1]).To(Equal(cancion.Ansiedad))
				Expect(cancionCorrecta.Sensaciones[2]).To(Equal(cancion.Desafio))
				Expect(cancionCorrecta.Sensaciones[3]).To(Equal(cancion.Sueño))
			})

			It("La canción debe seguir en la lista de colaboradas del usuario", func() {
				Expect(colaborador.CancionesColaboradas[0]).To(Equal(&cancionCorrecta))
			})

			It("No debe dar error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("Crear una nueva obra", func() {
		Context("Se crea una obra con todos sus campos", func() {
			It("No debe dar error", func() {
				Expect(err_ns).NotTo(HaveOccurred())
				Expect(err_np).NotTo(HaveOccurred())
				Expect(err_nv).NotTo(HaveOccurred())
			})

			It("Debe tener todos los campos iguales", func() {
				Expect(serie.Titulo()).To(Equal("SeriePrueba"))
				Expect(serie.Temporada()).To(Equal(1))
				Expect(serie.Capitulo()).To(Equal(1))

				Expect(pelicula.Titulo()).To(Equal("PeliculaPrueba"))

				Expect(videojuego.Titulo()).To(Equal("VideojuegoPrueba"))
			})
		})

		Context("Se crea una obra con algún campo incorrecto", func() {
			BeforeEach(func() {
				serie, err_ns = colaborador.CrearSerie("SeriePrueba", -1, 1, []cancion.Cancion_info{})
				pelicula, err_np = colaborador.CrearPelicula("", []cancion.Cancion_info{})
				videojuego, err_nv = colaborador.CrearVideojuego("", []cancion.Cancion_info{})
			})

			It("Debe dar error", func() {
				Expect(err_ns).To(HaveOccurred())
				Expect(err_np).To(HaveOccurred())
				Expect(err_nv).To(HaveOccurred())
			})

			It("La obra debe estar vacía", func() {
				Expect(serie).To(Equal(obra.Serie{}))
				Expect(pelicula).To(Equal(obra.Pelicula{}))
				Expect(videojuego).To(Equal(obra.Videojuego{}))
			})
		})
	})
})
