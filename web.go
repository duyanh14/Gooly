package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//db := services.Database{}.Conn()

	//////////////////////////////////////////

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/search", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}