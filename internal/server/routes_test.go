package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jlgallego99/OSTfind/internal/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Routes", func() {
	var router *gin.Engine
	var w_v, w_p, w_s *httptest.ResponseRecorder
	var req_v, req_p, req_s *http.Request
	var nuevaOst_pv, nuevaOst_s gin.H
	var body_v, body_p, body_s gin.H
	var canciones []gin.H

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
				json.Unmarshal(w_v.Body.Bytes(), &body_v)
				json.Unmarshal(w_p.Body.Bytes(), &body_p)
				json.Unmarshal(w_s.Body.Bytes(), &body_s)
			})

			It("Deberia crear la OST", func() {
				Expect(body_v["message"]).To(Equal("OST creada"))
				Expect(body_p["message"]).To(Equal("OST creada"))
				Expect(body_s["message"]).To(Equal("OST creada"))

				Expect(body_v["ost"].(map[string]interface{})["canciones"].([]interface{})[0].(map[string]interface{})["Titulo"]).To(Equal("Cancion 1"))
				Expect(body_p["ost"].(map[string]interface{})["canciones"].([]interface{})[0].(map[string]interface{})["Titulo"]).To(Equal("Cancion 1"))
				Expect(body_s["ost"].(map[string]interface{})["canciones"].([]interface{})[0].(map[string]interface{})["Titulo"]).To(Equal("Cancion 1"))
			})

			It("La OST debería tener dos canciones", func() {
				Expect(len(body_v["ost"].(map[string]interface{})["canciones"].([]interface{}))).To(Equal(2))
				Expect(len(body_p["ost"].(map[string]interface{})["canciones"].([]interface{}))).To(Equal(2))
				Expect(len(body_s["ost"].(map[string]interface{})["canciones"].([]interface{}))).To(Equal(2))
			})

			It("El código HTTP debe ser 200", func() {
				Expect(w_v.Code).To(Equal(200))
				Expect(w_p.Code).To(Equal(200))
				Expect(w_s.Code).To(Equal(200))
			})
		})
	})
})
