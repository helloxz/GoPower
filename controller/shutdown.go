package controller

import (
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Shutdown(c *gin.Context) {
	out, err := exec.Command("shutdown", "/s").Output()
	if err != nil {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  "关机失败！" + string(out),
			"data": "",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  string(out),
			"data": "",
		})
	}

}
