package main

import (
	"github.com/gin-gonic/gin"//引入Gin框架
	"net/http"
	"peixun/middleware"//中间件
	"strconv"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,"123")
	})
	router.POST("/", func(c *gin.Context) {
		num1,error:=strconv.Atoi(c.PostForm("Num1"))
		if error!=nil{}
		num2,error:=strconv.Atoi(c.PostForm("Num2"))
		if error!=nil{}
		ans:=num1+num2
		c.JSON(http.StatusOK,gin.H{"ans":ans})
	})

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
	// router.Run(":3000") for a hard coded port
}