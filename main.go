package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "你好!这是gin的测试")
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	if err := r.Run(":8600").Error(); err != "" {
		panic("发生错误")
	}

}
