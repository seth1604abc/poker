package main

import (
	"flag"
	"fmt"
	"net/http"
	"poker/config"
	"poker/internal/model"

	"github.com/gin-gonic/gin"
)

func init() {
	env := flag.String("env", "dev", "Application enviroment (dev or pro)")
	flag.Parse()
	fmt.Println("Enviroment:", *env)
	
	// 初始化
	config.InitConfig(*env)
	model.InitMysql()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}