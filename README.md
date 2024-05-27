# Simple Go Album API

This is a simple web server written in Go using the Gin web framework. The server provides a basic API to manage a list of music albums.

## Features

- **Retrieve Albums**: Fetch a list of all albums.
- **Add Albums**: Add a new album to the list via a POST request.

## Requirements

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/simple-go-album-api.git
   cd simple-go-album-api
   ```

2. Initialize Go modules:
   ```sh
   go mod init example.com/simple-go-album-api
   go mod tidy
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints

### GET /albums

Retrieves the list of all albums.

- **URL**: `/albums`
- **Method**: `GET`
- **Response**: JSON array of album objects

### POST /albums

Adds a new album to the list.

- **URL**: `/albums`
- **Method**: `POST`
- **Request Body**: JSON object representing the album
- **Response**: JSON object of the newly added album

## Code Explanation

### album Struct

Represents data about a record album.
```go
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
```

### Sample Data

Initial list of albums.
```go
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```

### Handlers

#### getAlbums

Responds with the list of all albums as JSON.
```go
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
```

#### postAlbums

Adds a new album from JSON received in the request body.
```go
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
```

### main Function

Sets up the Gin router and routes, and starts the server.
```go
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
```