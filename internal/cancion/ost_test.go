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
		canciones = make([]*cancion.Cancion_info, 5)
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
})
