package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jlgallego99/OSTfind/internal/cancion"
	"github.com/jlgallego99/OSTfind/internal/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Respuesta struct {
	Message string        `json:"message"`
	OST     OST_Respuesta `json:"ost"`
}

type OST_Respuesta struct {
	Id        string                 `json:"id"`
	Nombre    string                 `json:"nombre"`
	Canciones []cancion.Cancion_info `json:"canciones"`
}

var _ = Describe("Routes", func() {
	var router *gin.Engine
	var w_v, w_p, w_s *httptest.ResponseRecorder
	var req_v, req_p, req_s *http.Request
	var nuevaOst_pv, nuevaOst_s gin.H
	var res_v, res_p, res_s *Respuesta
	var canciones []gin.H
	var err_v, err_p, err_s error

	BeforeEach(func() {
		router = server.SetupRoutes()
		w_v = httptest.NewRecorder()
		w_p = httptest.NewRecorder()
		w_s = httptest.NewRecorder()

		canciones = []gin.H{{
			"titulo":     "Cancion 1",
			"compositor": "Compositor 1",
			"genero":     "Rock",
		}, {
			"titulo":     "Cancion 2",
			"compositor": "Compositor 2",
			"genero":     "Ambiental",
		}}
		nuevaOst_pv = gin.H{
			"nombre":    "OST Prueba",
			"canciones": canciones,
		}
		nuevaOst_s = gin.H{
			"nombre":    "OST Prueba",
			"temporada": 1,
			"capitulo":  1,
			"canciones": canciones,
		}
	})

	Describe("Crear una banda sonora con POST", func() {
		BeforeEach(func() {
			body_v, _ := json.Marshal(nuevaOst_pv)
			req_v, _ = http.NewRequest("POST", "/osts/videojuego", bytes.NewReader(body_v))
			router.ServeHTTP(w_v, req_v)

			body_p, _ := json.Marshal(nuevaOst_pv)
			req_p, _ = http.NewRequest("POST", "/osts/videojuego", bytes.NewReader(body_p))
			router.ServeHTTP(w_p, req_p)

			body_s, _ := json.Marshal(nuevaOst_s)
			req_s, _ = http.NewRequest("POST", "/osts/videojuego", bytes.NewReader(body_s))
			router.ServeHTTP(w_s, req_s)
		})

		Context("La OST es correcta", func() {
			BeforeEach(func() {
				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)

				res_p = &Respuesta{}
				err_p = json.Unmarshal(w_p.Body.Bytes(), res_p)

				res_s = &Respuesta{}
				err_s = json.Unmarshal(w_s.Body.Bytes(), res_s)
			})

			It("El JSON de respuesta no debe tener errores", func() {
				Expect(err_v).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_s).NotTo(HaveOccurred())
			})

			It("Deberia crear la OST", func() {
				Expect(res_v.Message).To(Equal("OST creada"))
				Expect(res_p.Message).To(Equal("OST creada"))
				Expect(res_s.Message).To(Equal("OST creada"))

				Expect(res_v.OST.Nombre).To(Equal("OST Prueba"))
				Expect(res_p.OST.Nombre).To(Equal("OST Prueba"))
				Expect(res_s.OST.Nombre).To(Equal("OST Prueba"))
			})

			It("La OST debería tener dos canciones", func() {
				Expect(len(res_v.OST.Canciones)).To(Equal(2))
				Expect(len(res_p.OST.Canciones)).To(Equal(2))
				Expect(len(res_s.OST.Canciones)).To(Equal(2))
			})

			It("El código HTTP debe ser 200", func() {
				Expect(w_v.Code).To(Equal(200))
				Expect(w_p.Code).To(Equal(200))
				Expect(w_s.Code).To(Equal(200))
			})
		})
	})
})
