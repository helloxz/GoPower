package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	//获取当前时间
	currentTime := time.Now()
	//运行时间
	difftime := currentTime.Sub(now)
	//运行时间保留2位数
	runtime := strconv.FormatFloat(difftime.Hours(), 'f', 2, 64)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": runtime,
	})
}

//全局变量，由主函数调用并传递
var now time.Time

func GetStartTime(x time.Time) {
	now = x
}
