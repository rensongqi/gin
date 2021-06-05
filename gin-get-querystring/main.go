package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发请求携带的query string 参数
		//name := c.Query("query") // 通过query获取请求中携带的quert string
		//name := c.DefaultQuery("query", "somebody")  // 取不到就用默认值
		//
		name := c.Query("query")
		age := c.Query("age")

		//name, ok := c.GetQuery("query")	// 取不到第二个参数就返回false
		//if !ok {
		//	name = "RSQRSQ"
		//	age = 18
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})

	r.Run(":8002")
}
