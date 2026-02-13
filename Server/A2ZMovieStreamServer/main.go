package main

import (
	"fmt"

	controller "github.com/Pranjul20/A2ZMovieStream/Server/A2ZMovieStreamServer/Controllers"
	database "github.com/Pranjul20/A2ZMovieStream/Server/A2ZMovieStreamServer/Database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {

	router := gin.Default()

	var client *mongo.Client = database.Connect()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, A2ZMovieStream!")
	})

	router.GET("/movies", controller.GetMovies(client))
	router.GET("/movie/:imdb_id", controller.GetMovie(client))

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
