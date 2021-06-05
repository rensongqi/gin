package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	// Login Get
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// Login POST
	r.POST("login", func(c *gin.Context) {
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("xxx", "********")

		username , ok := c.GetPostForm("username")
		if !ok {
			username = "ssssssbbbbbb"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			// 获取form表单提交的数据
			"Name": username,
			"Password": password,
		})
	})
	r.Run(":8004")
}
