package handlers

import (
	"net/http"

	"github.com/Rishav-R03/RESTful_tut/models"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Album)
}
