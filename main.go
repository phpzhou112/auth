package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "你好!这是gin的测试212")
	})
	//heapOut,_ := os.Create("heap.pprof")
	//pprof.StartCPUProfile(heapOut)  // cpu
	//pprof.WriteHeapProfile(heapOut) // 内存情况

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	if err := r.Run(":1234").Error(); err != "" {
		panic("发生错误")
	}

}
