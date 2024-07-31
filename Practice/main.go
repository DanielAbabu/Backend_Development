package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = map[string]album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var new album

	if err := c.BindJSON(&new); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "Post is not created"})
		return
	}

	albums[new.ID] = new
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Post is created"})
}

func getAlbumID(c *gin.Context) {
	id := c.Param("id")
	newalbum, err := albums[id]

	if !err {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(200, newalbum)
}

func putAlbum(c *gin.Context) {
	var add album
	id := c.Param("id")

	if err := c.BindJSON(&add); err != nil {
		c.IndentedJSON(http.StatusNotModified, gin.H{"message": "Album is not sent properly"})
		return
	}

	old, err := albums[id]
	if !err {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album is not Available"})
		return
	}

	old.Artist = add.Artist
	old.Title = add.Title
	old.Price = add.Price

	albums[id] = old
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album is  updated"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	delete(albums, id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album is deleted"})

}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbum)
	router.GET("/albums/:id", getAlbumID)
	router.POST("/albums", postAlbum)
	router.PUT("/albums/:id", putAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")

}
