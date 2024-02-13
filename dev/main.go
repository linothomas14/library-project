package main

import (
	"fmt"
	"library-project/config"
	"library-project/controller"
	"library-project/repository"
	"library-project/service"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var ()

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}

func init() {
	if err := config.LoadConfig("."); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var db *mongo.Client
	var err error

	db, err = config.NewMongoDBClient()
	if err != nil {
		log.Printf("Error occured on DB : %v", err)
		return
	}
	log.Printf("Success Connect DB")

	defer config.CloseDatabaseConnection(db)

	bookRepository := repository.NewBookRepository(db)

	bookService := service.NewBookService(bookRepository)

	bookController := controller.NewBookController(bookService)

	r := gin.Default()

	bookRoutes := r.Group("api/books")
	{
		bookRoutes.GET("/", bookController.Fetch)
		bookRoutes.GET("/:id", bookController.FetchByID)
		bookRoutes.POST("/", bookController.Create)
	}

	r.GET("api/ping", PingHandler)
	r.Run(fmt.Sprintf("localhost:%d", config.Configuration.Server.Port))
}
