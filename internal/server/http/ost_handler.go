package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlgallego99/OSTfind/internal/cancion"
)

func newOST(c *gin.Context) {
	obra := c.Param("obra")
	ostName := c.Param("ost")
	var ost *cancion.BandaSonora
	var err error

	switch obra {
	case "videojuego":
		ost, err = cancion.NewVideojuegoOST(ostName, make([]*cancion.Cancion_info, 0))

	case "serie":
		ost, err = cancion.NewSerieOST(ostName, 1, 1, make([]*cancion.Cancion_info, 0))

	case "pelicula":
		ost, err = cancion.NewPeliculaOST(ostName, make([]*cancion.Cancion_info, 0))
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OST creada",
		"ost": gin.H{
			"nombre":    ost.Obra.Titulo(),
			"canciones": ost.Canciones,
		},
	})
}

func getOST(c *gin.Context) {

}
