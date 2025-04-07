package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	id := c.Param("id")
	for _, movie := range movies {
		if movie.ID == id {
			c.JSON(http.StatusOK, movie)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
}

func deleteMovie(c *gin.Context) {
	id := c.Param("id")
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			c.JSON(http.StatusOK, movie)
			return
		}
	}
}

func updateMovie(c *gin.Context) {
	id := c.Param("id")
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			var updatedMovie Movie
			if err := c.BindJSON(&updatedMovie); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
				return
			}
			updatedMovie.ID = id
			movies = append(movies, updatedMovie)
			c.JSON(http.StatusOK, updatedMovie)
			return
		}
	}
}

func addMovie(c *gin.Context) {
	var movie Movie
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	c.JSON(http.StatusCreated, movie)
}

func main() {
	router := gin.Default()

	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "Movie 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "123457", Title: "Movie 2", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	//endpoints
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovie)
	router.POST("/movies", addMovie)
	router.PUT("/movies/:id", updateMovie)
	router.DELETE("/movies/:id", deleteMovie)

	fmt.Printf("Server started at http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
