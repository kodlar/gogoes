package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ali":       "yazar",
			"veli":      "bozar",
			"gülsuyunu": "çeker",
			"azar":      "azar",
		})
	})

	r.Run()
}
