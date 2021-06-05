package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 1 使用map序列化json
	r.GET("/map", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name": "小王子",
		//	"message": "Hello, golang",
		//	"age": 18,
		//}
		data := gin.H{"age":18,"message":"Hello, golang","name":"小王子"}
		c.JSON(http.StatusOK, data)
	})

	// 2 使用struct返回序列化json
	type msg struct{
		Name string `json:"name"`
		Message string
		Age int
	}
	r.GET("/struct", func(c *gin.Context) {
		data := msg{
			Name: "RSQ",
			Message: "Hello, golang",
			Age: 28,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8001")
}
