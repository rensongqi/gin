package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	name, ok := c.Get("name")  // 跨中间件取值
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件,统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	c.Next()  // 调用后续的处理函数

	cost := time.Since(start)
	fmt.Printf("cost :%#v\n", cost)
	fmt.Println("m1 out ...")

	//c.Abort()  // 阻止调用后续函数
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "RSQ") // 跨中间件取值
	c.Abort() // 阻止调用后续函数
	fmt.Println("m2 out ...")
}

func authMiddleware(c *gin.Context) {
	/*
	是否登录进行判断
	if 是登录用户，权限验证
	c.Next()
	else
	c.Abort()
	*/
}

func authMiddleware2(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			/*
				存放具体的逻辑
				是否登录进行判断
				if 是登录用户，权限验证
				c.Next()
				else
				c.Abort()
			*/
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()  // 默认使用了Logger跟Recovery中间件
	//不使用任何中间件
	//r := gin.New()

	r.Use(m1, m2, authMiddleware2(true) ) //  全局注册中间件函数，优先级最高
	//r.GET("/index", m1, m2, indexHandler)
	r.GET("/index", indexHandler)

	// 路由组使用中间件法一：
	xxGroup := r.Group("/xx", authMiddleware2(true))
	xxGroup.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "xxGroup",
		})
	})

	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware2(true))
	xx2Group.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "xx2Group",
		})
	})
	r.Run(":8010")
}
