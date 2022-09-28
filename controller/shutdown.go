package controller

import (
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Shutdown(c *gin.Context) {
	//根据不同的操作系统执行不同
	sysType := runtime.GOOS
	var out []byte
	var err error
	if sysType == "windows" {
		out, err = exec.Command("shutdown", "/s", "/t", "60").Output()
	} else if sysType == "linux" {
		out, err = exec.Command("shutdown", "-h", "1").Output()
	} else {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  "当前系统尚不支持！",
			"data": "",
		})
		//终止执行
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  "关机命令执行失败！",
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
