package main

import "github.com/gin-gonic/gin"

func setupServer() {
	r := gin.Default()

	r.GET("/ping", pongGetHandler)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}

func pongGetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
