package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type Respuesta_multi struct {
	Message string          `json:"message"`
	OSTs    []OST_Respuesta `json:"osts"`
}

type OST_Respuesta struct {
	Id        string                 `json:"id"`
	Nombre    string                 `json:"nombre"`
	Canciones []cancion.Cancion_info `json:"canciones"`
}

var _ = Describe("Routes", func() {
	var router = server.SetupRoutes()
	var w_v, w_p, w_s *httptest.ResponseRecorder
	var req_v, req_p, req_s *http.Request
	var nuevaOst_pv, nuevaOst_s gin.H
	var res_v, res_p, res_s *Respuesta
	var resm_v *Respuesta_multi
	var canciones []gin.H
	var err_v, err_p, err_s error
	var id_v, id_p, id_s string

	BeforeEach(func() {
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
		Context("La OST es correcta", func() {
			BeforeEach(func() {
				body_v, _ := json.Marshal(nuevaOst_pv)
				req_v, _ = http.NewRequest("POST", "/osts/videojuego", bytes.NewReader(body_v))
				router.ServeHTTP(w_v, req_v)

				body_p, _ := json.Marshal(nuevaOst_pv)
				req_p, _ = http.NewRequest("POST", "/osts/pelicula", bytes.NewReader(body_p))
				router.ServeHTTP(w_p, req_p)

				body_s, _ := json.Marshal(nuevaOst_s)
				req_s, _ = http.NewRequest("POST", "/osts/serie", bytes.NewReader(body_s))
				router.ServeHTTP(w_s, req_s)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)

				res_p = &Respuesta{}
				err_p = json.Unmarshal(w_p.Body.Bytes(), res_p)

				res_s = &Respuesta{}
				err_s = json.Unmarshal(w_s.Body.Bytes(), res_s)

				id_v = res_v.OST.Id
				id_p = res_p.OST.Id
				id_s = res_s.OST.Id
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
				Expect(res_s.OST.Nombre).To(Equal("OST Prueba-1-1"))

				Expect(res_v.OST.Canciones[0].Titulo).To(Equal("Cancion 1"))
				Expect(res_p.OST.Canciones[0].Titulo).To(Equal("Cancion 1"))
				Expect(res_s.OST.Canciones[0].Titulo).To(Equal("Cancion 1"))
				Expect(res_v.OST.Canciones[1].Titulo).To(Equal("Cancion 2"))
				Expect(res_p.OST.Canciones[1].Titulo).To(Equal("Cancion 2"))
				Expect(res_s.OST.Canciones[1].Titulo).To(Equal("Cancion 2"))

				Expect(res_v.OST.Canciones[0].Compositor).To(Equal("Compositor 1"))
				Expect(res_p.OST.Canciones[0].Compositor).To(Equal("Compositor 1"))
				Expect(res_s.OST.Canciones[0].Compositor).To(Equal("Compositor 1"))
				Expect(res_v.OST.Canciones[1].Compositor).To(Equal("Compositor 2"))
				Expect(res_p.OST.Canciones[1].Compositor).To(Equal("Compositor 2"))
				Expect(res_s.OST.Canciones[1].Compositor).To(Equal("Compositor 2"))
			})

			It("La OST debería tener dos canciones", func() {
				Expect(len(res_v.OST.Canciones)).To(Equal(2))
				Expect(len(res_p.OST.Canciones)).To(Equal(2))
				Expect(len(res_s.OST.Canciones)).To(Equal(2))
			})

			It("El código HTTP debe ser 200", func() {
				Expect(w_v.Code).To(Equal(http.StatusOK))
				Expect(w_p.Code).To(Equal(http.StatusOK))
				Expect(w_s.Code).To(Equal(http.StatusOK))
			})
		})

		Context("El JSON es incorrecto", func() {
			BeforeEach(func() {
				// JSON que le faltan llaves y json vacio
				nuevaOst_pv := "\"nombre\": \"OST Prueba\", \"canciones\": [{ \"titulo\": \"Cancion 1\", \"compositor\": \"Compositor 1\", \"genero\": \"Rock\" }"
				nuevaOst_s := ""

				raw := json.RawMessage(nuevaOst_pv)
				body_v, _ := json.Marshal(raw)
				req_v, _ = http.NewRequest("POST", "/osts/videojuego", bytes.NewReader(body_v))
				router.ServeHTTP(w_v, req_v)

				body_p, _ := json.Marshal(raw)
				req_p, _ = http.NewRequest("POST", "/osts/pelicula", bytes.NewReader(body_p))
				router.ServeHTTP(w_p, req_p)

				raw = json.RawMessage(nuevaOst_s)
				body_s, _ := json.Marshal(raw)
				req_s, _ = http.NewRequest("POST", "/osts/serie", bytes.NewReader(body_s))
				router.ServeHTTP(w_s, req_s)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)

				res_p = &Respuesta{}
				err_p = json.Unmarshal(w_p.Body.Bytes(), res_p)

				res_s = &Respuesta{}
				err_s = json.Unmarshal(w_s.Body.Bytes(), res_s)
			})

			It("El código HTTP debe ser 400", func() {
				Expect(w_v.Code).To(Equal(http.StatusBadRequest))
				Expect(w_p.Code).To(Equal(http.StatusBadRequest))
				Expect(w_s.Code).To(Equal(http.StatusBadRequest))
			})

			It("El JSON de respuesta no debe tener errores", func() {
				Expect(err_v).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_s).NotTo(HaveOccurred())
			})
		})

		Context("Los campos del cuerpo tienen valores incorrectos", func() {
			BeforeEach(func() {
				nuevaOst_pv := "{\"nombre\": \"\", \"canciones\": [{ \"titulo\": \"Cancion 1\", \"compositor\": \"Compositor 1\", \"genero\": \"Rock\" }]}"
				nuevaOst_s := "{\"nombre\": \"Serie1\", \"canciones\": [{ \"titulo\": \"Cancion 1\", \"compositor\": \"Compositor 1\", \"genero\": \"Generoquenoexiste\" }]}"

				raw := json.RawMessage(nuevaOst_pv)
				body_v, _ := json.Marshal(raw)
				req_v, _ = http.NewRequest("POST", "/osts/videojuego", bytes.NewReader(body_v))
				router.ServeHTTP(w_v, req_v)

				body_p, _ := json.Marshal(raw)
				req_p, _ = http.NewRequest("POST", "/osts/pelicula", bytes.NewReader(body_p))
				router.ServeHTTP(w_p, req_p)

				raw = json.RawMessage(nuevaOst_s)
				body_s, _ := json.Marshal(raw)
				req_s, _ = http.NewRequest("POST", "/osts/serie", bytes.NewReader(body_s))
				router.ServeHTTP(w_s, req_s)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)

				res_p = &Respuesta{}
				err_p = json.Unmarshal(w_p.Body.Bytes(), res_p)

				res_s = &Respuesta{}
				err_s = json.Unmarshal(w_s.Body.Bytes(), res_s)
			})

			It("El código HTTP debe ser 400", func() {
				Expect(w_v.Code).To(Equal(http.StatusBadRequest))
				Expect(w_p.Code).To(Equal(http.StatusBadRequest))
				Expect(w_s.Code).To(Equal(http.StatusBadRequest))
			})

			It("El JSON de respuesta no debe tener errores", func() {
				Expect(err_v).NotTo(HaveOccurred())
				Expect(err_p).NotTo(HaveOccurred())
				Expect(err_s).NotTo(HaveOccurred())
			})
		})
	})

	Describe("Recuperar una OST concreta con GET", func() {
		Context("La OST existe", func() {
			BeforeEach(func() {
				req_v, _ = http.NewRequest("GET", "/osts/videojuego/"+id_v, nil)
				router.ServeHTTP(w_v, req_v)

				req_p, _ = http.NewRequest("GET", "/osts/pelicula/"+id_p, nil)
				router.ServeHTTP(w_p, req_p)

				req_s, _ = http.NewRequest("GET", "/osts/serie/"+id_s, nil)
				router.ServeHTTP(w_s, req_s)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)

				res_p = &Respuesta{}
				err_p = json.Unmarshal(w_p.Body.Bytes(), res_p)

				res_s = &Respuesta{}
				err_s = json.Unmarshal(w_s.Body.Bytes(), res_s)
			})

			It("El código HTTP debe ser 200", func() {
				Expect(w_v.Code).To(Equal(http.StatusOK))
				Expect(w_p.Code).To(Equal(http.StatusOK))
				Expect(w_s.Code).To(Equal(http.StatusOK))
			})

			It("La OST recuperada es correcta", func() {
				Expect(res_v.Message).To(Equal("OST encontrada"))
				Expect(res_p.Message).To(Equal("OST encontrada"))
				Expect(res_s.Message).To(Equal("OST encontrada"))

				Expect(res_v.OST.Nombre).To(Equal(nuevaOst_pv["nombre"]))
				Expect(res_p.OST.Nombre).To(Equal(nuevaOst_pv["nombre"]))
				Expect(res_s.OST.Nombre).To(Equal(nuevaOst_s["nombre"].(string) + "-" + fmt.Sprint(nuevaOst_s["temporada"].(int)) + "-" + fmt.Sprint(nuevaOst_s["capitulo"].(int))))

				Expect(res_v.OST.Canciones[0].Titulo).To(Equal("Cancion 1"))
				Expect(res_p.OST.Canciones[0].Titulo).To(Equal("Cancion 1"))
				Expect(res_s.OST.Canciones[0].Titulo).To(Equal("Cancion 1"))
				Expect(res_v.OST.Canciones[1].Titulo).To(Equal("Cancion 2"))
				Expect(res_p.OST.Canciones[1].Titulo).To(Equal("Cancion 2"))
				Expect(res_s.OST.Canciones[1].Titulo).To(Equal("Cancion 2"))

				Expect(res_v.OST.Canciones[0].Compositor).To(Equal("Compositor 1"))
				Expect(res_p.OST.Canciones[0].Compositor).To(Equal("Compositor 1"))
				Expect(res_s.OST.Canciones[0].Compositor).To(Equal("Compositor 1"))
				Expect(res_v.OST.Canciones[1].Compositor).To(Equal("Compositor 2"))
				Expect(res_p.OST.Canciones[1].Compositor).To(Equal("Compositor 2"))
				Expect(res_s.OST.Canciones[1].Compositor).To(Equal("Compositor 2"))
			})
		})

		Context("La OST existe pero no en ese tipo de obra", func() {
			BeforeEach(func() {
				req_v, _ = http.NewRequest("GET", "/osts/videojuego/"+id_p, nil)
				router.ServeHTTP(w_v, req_v)

				req_p, _ = http.NewRequest("GET", "/osts/pelicula/"+id_s, nil)
				router.ServeHTTP(w_p, req_p)

				req_s, _ = http.NewRequest("GET", "/osts/serie/"+id_v, nil)
				router.ServeHTTP(w_s, req_s)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)

				res_p = &Respuesta{}
				err_p = json.Unmarshal(w_p.Body.Bytes(), res_p)

				res_s = &Respuesta{}
				err_s = json.Unmarshal(w_s.Body.Bytes(), res_s)
			})

			It("El código HTTP debe ser 404", func() {
				Expect(w_v.Code).To(Equal(http.StatusNotFound))
				Expect(w_p.Code).To(Equal(http.StatusNotFound))
				Expect(w_s.Code).To(Equal(http.StatusNotFound))
			})

			It("El mensaje de error debe ser correcto", func() {
				Expect(res_v.Message).To(Equal("No existe esa OST para videojuego"))
				Expect(res_p.Message).To(Equal("No existe esa OST para pelicula"))
				Expect(res_s.Message).To(Equal("No existe esa OST para serie"))
			})
		})

		Context("La OST no existe", func() {
			BeforeEach(func() {
				req_v, _ = http.NewRequest("GET", "/osts/videojuego/123", nil)
				router.ServeHTTP(w_v, req_v)

				req_p, _ = http.NewRequest("GET", "/osts/pelicula/123", nil)
				router.ServeHTTP(w_p, req_p)

				req_s, _ = http.NewRequest("GET", "/osts/serie/123", nil)
				router.ServeHTTP(w_s, req_s)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)

				res_p = &Respuesta{}
				err_p = json.Unmarshal(w_p.Body.Bytes(), res_p)

				res_s = &Respuesta{}
				err_s = json.Unmarshal(w_s.Body.Bytes(), res_s)
			})

			It("El código HTTP debe ser 404", func() {
				Expect(w_v.Code).To(Equal(http.StatusNotFound))
				Expect(w_p.Code).To(Equal(http.StatusNotFound))
				Expect(w_s.Code).To(Equal(http.StatusNotFound))
			})

			It("El mensaje de error debe ser correcto", func() {
				Expect(res_v.Message).To(Equal("No existe esa OST para videojuego"))
				Expect(res_p.Message).To(Equal("No existe esa OST para pelicula"))
				Expect(res_s.Message).To(Equal("No existe esa OST para serie"))
			})
		})
	})

	Describe("Recuperar las OSTs con GET", func() {
		BeforeEach(func() {
			req_v, _ = http.NewRequest("GET", "/osts", nil)
			router.ServeHTTP(w_v, req_v)

			resm_v = &Respuesta_multi{}
			err_v = json.Unmarshal(w_v.Body.Bytes(), resm_v)
		})

		Context("Se han registrado anteriormente OSTs", func() {
			It("El JSON de respuesta no debe tener errores", func() {
				Expect(err_v).NotTo(HaveOccurred())
			})

			It("Se tienen todas las OSTs", func() {
				Expect(len(resm_v.OSTs)).To(Equal(12))
			})

			It("El código HTTP debe ser 200", func() {
				Expect(w_v.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Describe("Cualquier ruta no definida", func() {
		Context("Rutas inexistentes", func() {
			BeforeEach(func() {
				w_v = httptest.NewRecorder()

				req_v, _ = http.NewRequest("GET", "/no/existe", nil)
				router.ServeHTTP(w_v, req_v)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)
			})

			It("Debe dar error 404", func() {
				Expect(w_v.Code).To(Equal(http.StatusNotFound))
			})

			It("El mensaje de error debe ser correcto", func() {
				Expect(res_v.Message).To(Equal("Ruta inexistente"))
			})

			It("El JSON de respuesta no debe tener errores", func() {
				Expect(err_v).NotTo(HaveOccurred())
			})
		})

		Context("Usar método incorrecto en rutas existentes", func() {
			BeforeEach(func() {
				w_v = httptest.NewRecorder()

				req_v, _ = http.NewRequest("POST", "/osts", nil)
				router.ServeHTTP(w_v, req_v)

				res_v = &Respuesta{}
				err_v = json.Unmarshal(w_v.Body.Bytes(), res_v)
			})

			It("Debe dar error 404", func() {
				Expect(w_v.Code).To(Equal(http.StatusNotFound))
			})

			It("El mensaje de error debe ser correcto", func() {
				Expect(res_v.Message).To(Equal("Ruta inexistente"))
			})

			It("El JSON de respuesta no debe tener errores", func() {
				Expect(err_v).NotTo(HaveOccurred())
			})
		})
	})
})
