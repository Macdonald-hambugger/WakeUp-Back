package controller

import (
	"WakeUp-Back/database"
	"WakeUp-Back/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Controller(port string) {
	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			MaxAge:       24 * time.Hour,
		}))
	db := database.MySQLInit()

	r.POST("/signup/:id", func(c *gin.Context) {
		service.Login(db, c)
	})

	r.GET("/loadfriends", func(c *gin.Context) {
		service.LoadFriends(c)
	})

	r.Run(port)
}
