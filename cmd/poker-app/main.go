package main

import (
	"flag"
	"fmt"
	"net/http"
	"poker/config"
	"poker/internal/api/di"

	"github.com/gin-gonic/gin"
)

func init() {
	env := flag.String("env", "dev", "Application enviroment (dev or pro)")
	flag.Parse()
	fmt.Println("Enviroment:", *env)
	
	// 初始化
	config.InitConfig(*env)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appRoute, err := di.InitializeApp()
	if err != nil {
		fmt.Printf(`Initialize route error: %v`, err)
	}

	appRoute.SetUp(r)

	r.Run()
}