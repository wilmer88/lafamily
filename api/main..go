package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafamily/api/controllers"
	// "gorm-test/controllers"
	"net/http"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

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