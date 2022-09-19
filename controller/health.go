package controller

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": "",
	})
}
