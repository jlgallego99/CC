package cancion_test

import (
	"github.com/jlgallego99/OSTfind/internal/cancion"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OST", func() {
	var err_s, err_p, err_v error
	var ost_s, ost_p, ost_v *cancion.BandaSonora
	var cancionCorrecta *cancion.Cancion_info
	var cancionObtenida *cancion.Cancion_info
	var canciones []*cancion.Cancion_info
	var cancionesVacio []*cancion.Cancion_info

	BeforeEach(func() {
		cancionCorrecta = &cancion.Cancion_info{
			Titulo:         "a",
			Compositor:     "b",
			Genero:         cancion.Ambiental,
			Likes:          0,
			Dislikes:       0,
			Sensaciones:    make([]cancion.Sensacion, 0),
			Momento:        cancion.Ciudad,
			Momento_exacto: "",
		}

		cancionesVacio = make([]*cancion.Cancion_info, 0)
		canciones = make([]*cancion.Cancion_info, 0)
		canciones = append(canciones, cancionCorrecta)

		ost_v, err_v = cancion.NewVideojuegoOST("VideojuegoPrueba", cancionesVacio)
		ost_p, err_p = cancion.NewPeliculaOST("PeliculaPrueba", cancionesVacio)
		ost_s, err_s = cancion.NewSerieOST("SeriePrueba", 1, 1, cancionesVacio)
	})

	Describe("Crear OSTs", func() {
		Context("La OST es correcta", func() {
			BeforeEach(func() {
				ost_v, err_v = cancion.NewVideojuegoOST("VideojuegoPrueba", canciones)
				ost_p, err_p = cancion.NewPeliculaOST("PeliculaPrueba", canciones)
				ost_s, err_s = cancion.NewSerieOST("SeriePrueba", 1, 1, canciones)
			})

			It("No debe estar vacía", func() {
				Expect(ost_s).NotTo(Equal(cancion.BandaSonora{}))
				Expect(ost_p).NotTo(Equal(cancion.BandaSonora{}))
				Expect(ost_v).NotTo(Equal(cancion.BandaSonora{}))
			})

			It("La obra y lista de canciones deben ser las introducidas", func() {
				Expect(ost_s.Canciones).To(Equal(canciones))
				Expect(ost_p.Canciones).To(Equal(canciones))
				Expect(ost_v.Canciones).To(Equal(canciones))

				Expect(ost_s.Obra.Titulo()).To(Equal("SeriePrueba-1-1"))
				Expect(ost_p.Obra.Titulo()).To(Equal("PeliculaPrueba"))
				Expect(ost_v.Obra.Titulo()).To(Equal("VideojuegoPrueba"))
			})

			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})
		})

		Context("La OST es incorrecta", func() {
			BeforeEach(func() {
				ost_v, err_v = cancion.NewVideojuegoOST("VideojuegoPrueba", nil)
				ost_p, err_p = cancion.NewPeliculaOST("PeliculaPrueba", nil)
				ost_s, err_s = cancion.NewSerieOST("SeriePrueba", 1, 1, nil)
			})

			It("La lista de cancion debe estar vacía", func() {
				Expect(ost_s).To(Equal(&cancion.BandaSonora{}))
				Expect(ost_p).To(Equal(&cancion.BandaSonora{}))
				Expect(ost_v).To(Equal(&cancion.BandaSonora{}))
			})

			It("Debe dar error", func() {
				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})

		Context("La obra es incorrecta", func() {
			BeforeEach(func() {
				ost_v, err_v = cancion.NewVideojuegoOST("", nil)
				ost_p, err_p = cancion.NewPeliculaOST("", nil)
				ost_s, err_s = cancion.NewSerieOST("", -1, 1, nil)
			})

			It("La lista de cancion debe estar vacía", func() {
				Expect(ost_s).To(Equal(&cancion.BandaSonora{}))
				Expect(ost_p).To(Equal(&cancion.BandaSonora{}))
				Expect(ost_v).To(Equal(&cancion.BandaSonora{}))
			})

			It("Debe dar error", func() {
				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})
	})

	Describe("Añadir una nueva canción", func() {
		BeforeEach(func() {
			err_s = ost_s.NuevaCancion(cancionCorrecta)
			err_p = ost_p.NuevaCancion(cancionCorrecta)
			err_v = ost_v.NuevaCancion(cancionCorrecta)
		})

		Context("Se añade una canción no repetida", func() {
			It("La canción debe existir en la OST", func() {
				Expect(ost_s.Canciones[0]).To(Equal(cancionCorrecta))
				Expect(ost_p.Canciones[0]).To(Equal(cancionCorrecta))
				Expect(ost_v.Canciones[0]).To(Equal(cancionCorrecta))
			})

			It("Debe existir exactamente una canción", func() {
				Expect(len(ost_s.Canciones)).To(Equal(1))
				Expect(len(ost_p.Canciones)).To(Equal(1))
				Expect(len(ost_v.Canciones)).To(Equal(1))
			})

			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})
		})

		Context("Se añade una canción repetida", func() {
			BeforeEach(func() {
				err_s = ost_s.NuevaCancion(cancionCorrecta)
				err_p = ost_p.NuevaCancion(cancionCorrecta)
				err_v = ost_v.NuevaCancion(cancionCorrecta)
			})

			It("Debe seguir existiendo exactamente una canción", func() {
				Expect(len(ost_s.Canciones)).To(Equal(1))
				Expect(len(ost_p.Canciones)).To(Equal(1))
				Expect(len(ost_v.Canciones)).To(Equal(1))
			})

			It("Debe dar error", func() {
				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})

		Context("Se añade una canción vacía", func() {
			BeforeEach(func() {
				err_s = ost_s.NuevaCancion(&cancion.Cancion_info{})
				err_p = ost_p.NuevaCancion(&cancion.Cancion_info{})
				err_v = ost_v.NuevaCancion(&cancion.Cancion_info{})
			})

			It("Debe dar error", func() {
				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})
	})

	Describe("Actualizar la OST de una obra", func() {
		BeforeEach(func() {
			err_s = ost_s.ActualizarOST(canciones)
			err_p = ost_p.ActualizarOST(canciones)
			err_v = ost_v.ActualizarOST(canciones)
		})

		Context("La OST es correcta", func() {
			It("Se tiene exactamente una canción", func() {
				Expect(len(ost_s.Canciones)).To(Equal(1))
				Expect(len(ost_p.Canciones)).To(Equal(1))
				Expect(len(ost_v.Canciones)).To(Equal(1))
			})

			It("La canción introducida es correcta", func() {
				Expect(ost_s.Canciones[0]).To(Equal(cancionCorrecta))
				Expect(ost_p.Canciones[0]).To(Equal(cancionCorrecta))
				Expect(ost_v.Canciones[0]).To(Equal(cancionCorrecta))
			})

			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})
		})

		Context("La OST es incorrecta por tener canciones vacías", func() {
			BeforeEach(func() {
				cancionesVacio = append(cancionesVacio, &cancion.Cancion_info{})

				err_s = ost_s.ActualizarOST(cancionesVacio)
				err_p = ost_p.ActualizarOST(cancionesVacio)
				err_v = ost_v.ActualizarOST(cancionesVacio)
			})

			It("Se tiene exactamente una canción", func() {
				Expect(len(ost_s.Canciones)).To(Equal(1))
				Expect(len(ost_p.Canciones)).To(Equal(1))
				Expect(len(ost_v.Canciones)).To(Equal(1))
			})

			It("La lista de canciones no debe haberse modificado", func() {
				Expect(ost_s.Canciones[0]).To(Equal(cancionCorrecta))
				Expect(ost_p.Canciones[0]).To(Equal(cancionCorrecta))
				Expect(ost_v.Canciones[0]).To(Equal(cancionCorrecta))
			})

			It("Debe dar error", func() {
				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})

		Context("La nueva OST no tiene canciones", func() {
			BeforeEach(func() {
				err_s = ost_s.ActualizarOST(cancionesVacio)
				err_p = ost_p.ActualizarOST(cancionesVacio)
				err_v = ost_v.ActualizarOST(cancionesVacio)
			})

			It("La lista de canciones debe estar vacía", func() {
				Expect(ost_s.Canciones).To(BeEmpty())
				Expect(ost_p.Canciones).To(BeEmpty())
				Expect(ost_v.Canciones).To(BeEmpty())
			})

			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})
		})
	})

	Describe("Actualizar los datos de la obra", func() {
		Context("Los datos de la obra son correctos", func() {
			BeforeEach(func() {
				err_v = ost_v.ActualizarObra("NuevoTitulo")
				err_p = ost_p.ActualizarObra("NuevoTitulo")
				err_s = ost_s.ActualizarObra("NuevoTitulo", 2, 2)
			})

			It("Se ha modificado el título de la obra correctamente", func() {
				Expect(ost_v.Obra.Titulo()).To(Equal("NuevoTitulo"))
				Expect(ost_p.Obra.Titulo()).To(Equal("NuevoTitulo"))
				Expect(ost_s.Obra.Titulo()).To(Equal("NuevoTitulo-2-2"))
			})

			It("No debe dar error", func() {
				Expect(err_v).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_s).NotTo(HaveOccurred())
			})
		})

		Context("Los datos de obra son incorrectos", func() {
			BeforeEach(func() {
				err_v = ost_v.ActualizarObra("")
				err_p = ost_p.ActualizarObra("")
				err_s = ost_s.ActualizarObra("NuevoTitulo", 0)
			})

			It("El título de la obra debe permanecer igual", func() {
				Expect(ost_s.Obra.Titulo()).To(Equal("SeriePrueba-1-1"))
				Expect(ost_p.Obra.Titulo()).To(Equal("PeliculaPrueba"))
				Expect(ost_v.Obra.Titulo()).To(Equal("VideojuegoPrueba"))
			})

			It("Debe dar error", func() {
				Expect(err_v).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_s).To(HaveOccurred())

				Expect(err_v.Error()).To(Equal("no se ha podido actualizar la obra: el videojuego no tiene título"))
				Expect(err_p.Error()).To(Equal("no se ha podido actualizar la obra: la película no tiene título"))
				Expect(err_s.Error()).To(Equal("no se ha podido actualizar la obra: no se ha especificado tanto capítulo como temporada"))
			})
		})
	})

	Describe("Devolver una canción de la OST por su nombre", func() {
		BeforeEach(func() {
			err_s = ost_s.NuevaCancion(cancionCorrecta)
			err_p = ost_p.NuevaCancion(cancionCorrecta)
			err_v = ost_v.NuevaCancion(cancionCorrecta)
		})

		Context("La canción existe en la OST", func() {
			BeforeEach(func() {
				cancionObtenida, err_s = ost_s.Cancion("a")
				cancionObtenida, err_p = ost_p.Cancion("a")
				cancionObtenida, err_v = ost_v.Cancion("a")
			})

			It("La canción devuelta debe ser la que se quiere", func() {
				Expect(cancionObtenida).To(Equal(cancionCorrecta))
			})

			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})
		})

		Context("La canción no existe en la OST", func() {
			BeforeEach(func() {
				cancionObtenida, err_s = ost_s.Cancion("Noexiste")
				cancionObtenida, err_p = ost_p.Cancion("Noexiste")
				cancionObtenida, err_v = ost_v.Cancion("Noexiste")
			})

			It("La canción devuelta debe ser nula", func() {
				Expect(cancionObtenida).To(BeNil())
			})

			It("Debe dar error", func() {
				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})
	})
})
