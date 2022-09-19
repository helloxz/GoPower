package controller

import (
	"os/exec"
	"regexp"

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
		}
	}

	//执行系统命令，仅支持Linux X64
	out, err := exec.Command("./bin/wol", mac, ip).Output()

	if err != nil {
		c.JSON(-1000, gin.H{
			"code": -1000,
			"msg":  string(out),
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
