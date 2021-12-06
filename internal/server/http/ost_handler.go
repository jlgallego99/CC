package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

// Guardar temporalmente las OSTs en una variable global
var osts []*cancion.BandaSonora

func newOST(c *gin.Context) {
	var ost *cancion.BandaSonora
	var err error

	obra := c.Param("obra")
	ostName := c.Param("ost")

	switch obra {
	case "videojuego":
		ost, err = cancion.NewVideojuegoOST(ostName, make([]*cancion.Cancion_info, 0))

	case "serie":
		ost, err = cancion.NewSerieOST(ostName, 1, 1, make([]*cancion.Cancion_info, 0))

	case "pelicula":
		ost, err = cancion.NewPeliculaOST(ostName, make([]*cancion.Cancion_info, 0))

	default:
		err = errors.New("no se reconoce el tipo de OST")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	osts = append(osts, ost)

	c.JSON(http.StatusOK, gin.H{
		"message": "OST creada",
		"ost": gin.H{
			"nombre":    ost.Obra.Titulo(),
			"canciones": ost.Canciones,
		},
	})
}

func getOST(c *gin.Context) {
	var err error

	obra := c.Param("obra")
	ostName := c.Param("ost")

	switch obra {
	case "videojuego", "serie", "pelicula":
		err = nil

	default:
		err = errors.New("no se reconoce el tipo de OST")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for _, ost := range osts {
		if ost.Obra.Titulo() == ostName {
			err = nil

			c.JSON(http.StatusOK, gin.H{
				"message": "OST encontrada",
				"ost": gin.H{
					"nombre":    ost.Obra.Titulo(),
					"canciones": ost.Canciones,
				},
			})

			return
		} else {
			err = errors.New("no existe esa OST")
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
