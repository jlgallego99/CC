package server

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

// Guardar temporalmente las OSTs en una variable global
var osts []*cancion.BandaSonora

type Cancion_msg struct {
	Titulo     string `json:"titulo"`
	Compositor string `json:"compositor"`
	Genero     string `json:"genero"`
}

type Ost_msg struct {
	Nombre    string        `json:"nombre"`
	Temporada int           `json:"temporada"`
	Capitulo  int           `json:"capitulo"`
	Canciones []Cancion_msg `json:"canciones"`
}

func newOST(c *gin.Context) {
	var ost *cancion.BandaSonora
	var canciones []*cancion.Cancion_info
	var err error

	// Leer cuerpo de la petición
	ostmsg := new(Ost_msg)
	err = c.BindJSON(ostmsg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	obra := c.Param("obra")
	switch obra {
	case "videojuego":
		ost, err = cancion.NewVideojuegoOST(ostmsg.Nombre, make([]*cancion.Cancion_info, 0))

	case "serie":
		ost, err = cancion.NewSerieOST(ostmsg.Nombre, ostmsg.Temporada, ostmsg.Capitulo, make([]*cancion.Cancion_info, 0))

	case "pelicula":
		ost, err = cancion.NewPeliculaOST(ostmsg.Nombre, make([]*cancion.Cancion_info, 0))

	default:
		err = errors.New("no se reconoce el tipo de OST")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	// Añadir canciones de la ost
	for _, cmsg := range ostmsg.Canciones {
		can, err := cancion.NewCancion(cmsg.Titulo, cmsg.Compositor, cancion.StringToGenero[cmsg.Genero])

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}

		canciones = append(canciones, can)
	}

	ost.ActualizarOST(canciones)
	osts = append(osts, ost)

	c.JSON(http.StatusOK, gin.H{
		"message": "OST creada",
		"ost": gin.H{
			"id":        ost.Id,
			"nombre":    ost.Obra.Titulo(),
			"canciones": ost.Canciones,
		},
	})
}

func getOST(c *gin.Context) {
	var err error

	obra := c.Param("obra")
	ostId := c.Param("ostid")

	switch obra {
	case "videojuego", "serie", "pelicula":
		err = nil

	default:
		err = errors.New("no se reconoce el tipo de OST")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	for _, ost := range osts {
		if ost.Id == ostId && strings.EqualFold("obra."+obra, reflect.TypeOf(ost.Obra).String()) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OST encontrada",
				"ost": gin.H{
					"id":        ost.Id,
					"nombre":    ost.Obra.Titulo(),
					"canciones": ost.Canciones,
				},
			})

			return
		} else {
			err = errors.New("No existe esa OST para " + obra)
		}
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}
}

func updateOST(c *gin.Context) {

}

func allOsts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Todas las OSTs del sistema",
		"osts":    osts,
	})
}

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Ruta inexistente"})
}
