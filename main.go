package main

import (
	"austinmarten1/supreme-octo-giggle/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/meme", handlers.GetMeme)

	r.Static("/memes", "./memes")
	r.Run(":8080")
}
