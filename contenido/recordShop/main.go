package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

var albums = []Album{
	{ID: "45aeaf77-6a14-41d2-9a6f-58223d2e836a", Title: "Oooh", Artist: "Maju", Price: 14.99},
	{ID: "a1c85ff7-5599-47b8-8d92-e4253b3964d7", Title: "Aaaah", Artist: "Verg", Price: 14.50},
	{ID: "bec2ee4d-9196-4cf8-bb68-ee4e5d140cd4", Title: "Eeee", Artist: "Bault", Price: 14.10},
}

// Controllers
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	err := c.BindJSON(&newAlbum)

	if err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}
