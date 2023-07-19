package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcoantonio63/crud-api/app/config"
	route "github.com/marcoantonio63/crud-api/app/routes"
)

func main() {
	r := gin.Default()

	config.InitDatabase()

	route.UserRoutes(r.Group("/user"))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name":    "Api nome",
			"version": 1.0,
		})
	})

	r.Run(":3000")
}
