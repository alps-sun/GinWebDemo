package main

/**
* @Author Mario
* @Date 2022-04-09 15:12
* @Version ：
* @Description :
 */

import "github.com/gin-gonic/gin"

func gintest() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run("127.0.0.1:8888")
}

func main() {
	gintest()
}
