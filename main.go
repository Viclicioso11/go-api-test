package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"entities"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run()
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entities.GetAlbums())
}
