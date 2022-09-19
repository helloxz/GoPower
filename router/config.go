package router

import (
	"fmt"
	"http_shutdown/controller"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

/*认证中间件*/
func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取key参数
		key := c.PostForm("key")
		//获取配置中的key
		//读取配置文件
		cfg, _ := ini.Load("config.ini")
		get_key := cfg.Section("infos").Key("key").String()

		if key != get_key {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "认证失败",
				"data": "",
			})
			c.Abort()
		} else {
			c.Next()
		}

	}
}

/*跨域中间件*/
func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		// 允许放行OPTIONS请求
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

// 所有路由，全部配置到这里
func Start() {
	//读取配置文件
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	port := cfg.Section("servers").Key("port").String()
	//gin运行模式
	RunMode := cfg.Section("servers").Key("RunMode").String()
	gin.SetMode(RunMode)

	r := gin.Default()

	//使用跨域
	r.Use(cors())
	//指定一个路由，并指定用哪个方法来处理
	r.POST("/api/shutdown", auth(), controller.Shutdown)
	r.POST("/api/health", auth(), controller.Health)
	r.POST("/api/startup", auth(), controller.StartUP)

	//运行gin，并从配置文件指定端口
	r.Run(port)
}
