package main

import (
	"gin_instrumentation/handler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	routerGroup := r.Group("api/v1")
	routerGroup.Use(handler.MwPromethusHandler) // set before other routing rules

	routerGroup.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routerGroup.GET("metrics", gin.WrapH(promhttp.Handler()))
	r.Run()
}
