package main

import (
	_ "07_package/main/pack/pack1"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

// 程序的执行开始于
// 导入包，
// 初始化 main 包
// 然后调用 main 函数。
func main() {

	// . "07_package/main/pack/pack1" 可以直接使用其全局变量和方法

	//  _ "07_package/main/pack/pack1" 只会执行初始化变量和初始化方法
	/*
		fmt.Printf("pack1.Pack1Float: %v\n", Pack1Float)
		fmt.Printf("pack1.Pack1Float: %v\n", Pack1Float)
		fmt.Printf("pack1.ReturnStr(): %v\n", ReturnStr())
	*/

	// r := gin.Default()
	// r.GET("/ping",func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK,gin.H{
	// 		"message":"pong",
	// 	})
	// })



	// r.Run()
}
