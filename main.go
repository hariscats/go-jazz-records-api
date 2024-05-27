package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Kind of Blue", Artist: "Miles Davis", Price: 32.99},
	{ID: "5", Title: "The Shape of Jazz to Come", Artist: "Ornette Coleman", Price: 29.99},
	{ID: "6", Title: "A Love Supreme", Artist: "John Coltrane", Price: 34.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID responds with the album whose ID matches the given ID.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := findAlbumByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the new album.
	if err := validateAlbum(newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// findAlbumByID searches for an album by ID.
func findAlbumByID(id string) (Album, error) {
	for _, album := range albums {
		if album.ID == id {
			return album, nil
		}
	}
	return Album{}, http.ErrNoLocation
}

// validateAlbum checks if the album contains all necessary fields.
func validateAlbum(album Album) error {
	if album.ID == "" || album.Title == "" || album.Artist == "" || album.Price <= 0 {
		return fmt.Errorf("invalid album data")
	}
	return nil
}
