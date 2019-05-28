package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	// engine.Any("/", WebRoot)
	engine.POST("/svn_code", SvnCode)
	// 绑定端口，然后启动应用
	engine.Run("127.0.0.1:9206")
}

/*
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	context.String(403, "forbidden")
	// context.String(http.StatusOK, "hello, world")
}

/*
* 升级判断，并推送文件
 */
func SvnCode(context *gin.Context) {

	type svnc struct {
		name string
		age  int
	}
	var svncm svnc

	data, _ := context.Request.Body.Read(svncm)
	log.Println(data)
	// fmt.Println(data)
	context.String(http.StatusOK, "fine!")
}
