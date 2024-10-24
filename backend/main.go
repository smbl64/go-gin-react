package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	{
		api := r.Group("api")
		api.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"username": "admin",
				"email":    "foo@bar.nl",
			})

		})
	}

	r.Run()
}
