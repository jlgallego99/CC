package obra_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jlgallego99/OSTfind/internal/obra"
)

var _ = Describe("Obra", func() {
	var serie obra.Serie
	var pelicula obra.Pelicula
	var videojuego obra.Videojuego
	var err_ns, err_np, err_nv error

	Describe("Crear nueva obra", func() {
		Context("Se crea una obra con todos sus campos", func() {
			BeforeEach(func() {
				serie, err_ns = obra.NewSerie("SeriePrueba", 1, 1)
				pelicula, err_np = obra.NewPelicula("PeliculaPrueba")
				videojuego, err_nv = obra.NewVideojuego("VideojuegoPrueba")
			})

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
				serie, err_ns = obra.NewSerie("SeriePrueba", -1, 1)
				pelicula, err_np = obra.NewPelicula("")
				videojuego, err_nv = obra.NewVideojuego("")
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
