package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	{
		api := r.Group("api")
		api.Use(cors.Default())
		api.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"username": "admin",
				"email":    "foo@bar.nl",
			})

		})

		// TODO: This in an  allows-all
		// See docs https://github.com/gin-contrib/cors
	}

	r.Run()
}
