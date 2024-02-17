package controller

import (
	"github.com/gin-gonic/gin"
)

func Controller(port string) {
	r := gin.Default()

	//db := database.MySQLInit()

	r.Run(port)
}
