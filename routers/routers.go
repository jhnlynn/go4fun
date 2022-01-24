package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
	"net/http"
	"os"

	//"go-play/middleware"
)

func Routers() *gin.Engine {

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		cors.Default()
		c.Next()
		gin.Logger()
		gin.Recovery()
		//middleware.CORSMiddleware()
	})


	router.GET("/", func(c *gin.Context) {
		log.Println("process...")
		c.JSON(http.StatusOK, gin.H {
			"s": os.Getenv("en"),
		})
	})

	return router
}
