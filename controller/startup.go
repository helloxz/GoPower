package controller

import (
	"os/exec"
	"regexp"
	"runtime"

	"github.com/gin-gonic/gin"
)

func StartUP(c *gin.Context) {
	//获取MAC地址
	mac := string(c.PostForm("mac"))
	//获取IP
	ip := string(c.PostForm("ip"))

	//判断MAC地址是否正确
	var validMAC = regexp.MustCompile(`^[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}:[0-9a-zA-Z]{2}$`)
	v_re := validMAC.MatchString(mac)

	//如果MAC地址不正确
	if !v_re {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  "MAC地址格式不正确！",
			"data": "",
		})
		c.Abort()
		return
	}
	//如果IP不为空，还要验证IP是否正确
	if ip != "" {
		var validIP = regexp.MustCompile(`^(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}$`)
		v_re := validIP.MatchString(ip)

		if !v_re {
			c.JSON(-1000, gin.H{
				"code": -1000,
				"msg":  "IP地址格式不正确！",
				"data": "",
			})
			c.Abort()
			return
		}
	}

	//发起wol唤醒指令
	sysType := runtime.GOOS
	var bin string
	if sysType == "linux" {
		bin = "./bin/wol"
	} else if sysType == "windows" {
		bin = "./bin/wol.exe"
	} else {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  "当前系统不支持wol命令！",
			"data": "",
		})
		return
	}
	_, err := exec.Command(bin, mac, ip).Output()

	if err != nil {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  "命令执行失败！",
			"data": "",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": "",
		})
	}
}
