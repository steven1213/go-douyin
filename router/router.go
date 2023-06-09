package router

import "github.com/gin-gonic/gin"

func InitRouter(model string) *gin.Engine {
	r := gin.Default()
	if model == "debug" {
		r.Use(gin.Logger())
	}

	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	return r
}
