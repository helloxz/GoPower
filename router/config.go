package router

import (
	"http_shutdown/controller"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

/*获取配置*/
var conf *ini.File

/*初始化conf配置*/
func Init_conf([]string) {
	//获取命令行参数
	args := os.Args
	//获取切片长度
	args_cap := len(args)
	var conf_path string
	if args_cap == 3 && args[1] == "-c" {
		conf_path = args[2]
	} else {
		conf_path = "config.ini"
	}

	//读取配置
	conf, _ = ini.Load(conf_path)
}

/*认证中间件*/
func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取key参数
		key := c.PostForm("key")
		//获取配置中的key
		//读取配置文件
		get_key := conf.Section("infos").Key("key").String()

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
	port := conf.Section("servers").Key("port").String()
	//gin运行模式
	RunMode := conf.Section("servers").Key("RunMode").String()
	gin.SetMode(RunMode)

	//日志记录到文件
	f, _ := os.Create("gopower.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//日志同时输出到控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	//使用跨域
	r.Use(cors())
	//指定一个路由，并指定用哪个方法来处理
	r.POST("/api/shutdown", auth(), controller.Shutdown)
	r.POST("/api/health", auth(), controller.Health)
	r.POST("/api/startup", auth(), controller.StartUP)
	r.POST("/load_settings",controller.LoadSettings)
	r.POST("/save_settings",controller.SaveSettings)

	//运行webui
	webui_status := conf.Section("servers").Key("webui").String()
	if webui_status == "on" {
		r.StaticFS("/", http.Dir("./front/dist"))
	} else if webui_status != "on" {
		r.LoadHTMLFiles("./webui/403.html")
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "403.html", gin.H{})
		})
	}

	//运行gin，并从配置文件指定端口
	r.Run(port)
}
