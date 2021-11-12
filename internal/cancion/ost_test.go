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
			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})

			It("No debe estar vacía", func() {
				Expect(ost_s).NotTo(Equal(cancion.BandaSonora{}))
				Expect(ost_p).NotTo(Equal(cancion.BandaSonora{}))
				Expect(ost_v).NotTo(Equal(cancion.BandaSonora{}))
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
})
