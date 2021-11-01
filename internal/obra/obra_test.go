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
	var err_ns, err_np, err_nv error
	var err_s, err_p, err_v error

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

		serie, err_ns = obra.NewSerie("SeriePrueba", 1, 1, []cancion.Cancion_info{cancionCorrecta, {}})
		pelicula, err_np = obra.NewPelicula("PeliculaPrueba", []cancion.Cancion_info{cancionCorrecta, {}})
		videojuego, err_nv = obra.NewVideojuego("VideojuegoPrueba", []cancion.Cancion_info{cancionCorrecta, {}})
	})

	Describe("Crear nueva obra", func() {
		Context("Se crea una obra con todos sus campos", func() {
			It("No debe dar error", func() {
				Expect(err_ns).NotTo(HaveOccurred())
				Expect(err_np).NotTo(HaveOccurred())
				Expect(err_nv).NotTo(HaveOccurred())
			})

			It("Debe tener todos los campos iguales", func() {
				Expect(serie.Titulo).To(Equal("SeriePrueba"))
				Expect(serie.Temporada).To(Equal(1))
				Expect(serie.Capitulo).To(Equal(1))
				Expect(serie.OST).To(Equal([]cancion.Cancion_info{cancionCorrecta, {}}))

				Expect(pelicula.Titulo).To(Equal("PeliculaPrueba"))
				Expect(pelicula.OST).To(Equal([]cancion.Cancion_info{cancionCorrecta, {}}))

				Expect(videojuego.Titulo).To(Equal("VideojuegoPrueba"))
				Expect(videojuego.OST).To(Equal([]cancion.Cancion_info{cancionCorrecta, {}}))
			})
		})

		Context("Se crea una obra con algún campo incorrecto", func() {
			BeforeEach(func() {
				serie, err_ns = obra.NewSerie("SeriePrueba", -1, 1, []cancion.Cancion_info{cancionCorrecta, {}})
				pelicula, err_np = obra.NewPelicula("", []cancion.Cancion_info{cancionCorrecta, {}})
				videojuego, err_nv = obra.NewVideojuego("", []cancion.Cancion_info{cancionCorrecta, {}})
			})

			It("Debe dar error", func() {
				Expect(err_ns).To(HaveOccurred())
				Expect(err_np).To(HaveOccurred())
				Expect(err_nv).To(HaveOccurred())
			})

			It("La obra debe estar vacía", func() {
				Expect(serie).To(BeNil())
				Expect(pelicula).To(BeNil())
				Expect(videojuego).To(BeNil())
			})
		})
	})

	Describe("Añadir una canción a una obra", func() {
		BeforeEach(func() {
			serie.OST = []cancion.Cancion_info{}
			pelicula.OST = []cancion.Cancion_info{}
			videojuego.OST = []cancion.Cancion_info{}

			err_s = serie.NuevaCancion(cancionCorrecta)
			err_p = pelicula.NuevaCancion(cancionCorrecta)
			err_v = videojuego.NuevaCancion(cancionCorrecta)
		})

		Context("Se añade una canción nueva", func() {
			It("Debe existir la canción en la OST", func() {
				Expect(serie.OST[0]).To(Equal(cancionCorrecta))
				Expect(pelicula.OST[0]).To(Equal(cancionCorrecta))
				Expect(videojuego.OST[0]).To(Equal(cancionCorrecta))
			})

			It("No debe dar error", func() {
				Expect(err_s).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_v).NotTo(HaveOccurred())
			})
		})

		Context("Se añade una canción vacía", func() {
			It("Debe dar error", func() {
				err_s = serie.NuevaCancion(cancion.Cancion_info{})
				err_p = pelicula.NuevaCancion(cancion.Cancion_info{})
				err_v = videojuego.NuevaCancion(cancion.Cancion_info{})

				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})

		Context("Se añade una canción que ya existe en la OST", func() {
			BeforeEach(func() {
				err_s = serie.NuevaCancion(cancionCorrecta)
				err_p = pelicula.NuevaCancion(cancionCorrecta)
				err_v = videojuego.NuevaCancion(cancionCorrecta)
			})

			It("No debe haber más de una canción en la OST", func() {
				Expect(len(serie.OST)).To(Equal(1))
				Expect(len(pelicula.OST)).To(Equal(1))
				Expect(len(videojuego.OST)).To(Equal(1))
			})

			It("Debe dar error", func() {
				Expect(err_s).To(HaveOccurred())
				Expect(err_p).To(HaveOccurred())
				Expect(err_v).To(HaveOccurred())
			})
		})
	})
})
