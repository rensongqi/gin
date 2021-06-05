package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main()  {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 处理multipart forms提交文件时默认的内存限制是32MiB
	// 可通过下边这个方式去修改
	r.MaxMultipartMemory = 8
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取单个文件文件
		f, err := c.FormFile("f1")  // 从请求中获取携带的参数  一致
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		//将读取到的文件保存在服务端的本地
	})
	r.Run(":8006")
}
