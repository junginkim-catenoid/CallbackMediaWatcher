package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	println("Hello CallbackMediaWatcher app")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//time.Sleep(time.Second * 15)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3002") //port
}
