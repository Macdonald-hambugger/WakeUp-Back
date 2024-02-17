package controller

import (
	"WakeUp-Back/database"
	"WakeUp-Back/service"
	"github.com/gin-gonic/gin"
)

func Controller(port string) {
	r := gin.Default()

	db := database.MySQLInit()

	r.POST("/signup/:id", func(c *gin.Context) {
		service.Login(db, c)
	})
	r.Run(port)
}
