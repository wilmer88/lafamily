package main

import (
	// "os"
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafamily/api/controllers"
	// "gorm-test/controllers"
	"net/http"
	"github.com/gin-contrib/cors"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {

	// port := os.Getenv("Port")
	// if port ==""{
	// 	port = "8080"
	// }

	r := setupRouter()
	_ = r.Run(":8080")
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"https://lafamily.herokuapp.com/"}
	config.AllowOrigins = []string{"https://localhost:8080/"}
	r.Use(cors.New(config))
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	userRepo := controllers.New()
	r.POST("/lafamily", userRepo.CreateUser)
	r.GET("/lafamily", userRepo.GetUsers)
	r.GET("/lafamily/:id", userRepo.GetUser)
	r.PUT("/lafamily/:id", userRepo.UpdateUser)
	r.DELETE("/lafamily/:id", userRepo.DeleteUser)

	return r
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}