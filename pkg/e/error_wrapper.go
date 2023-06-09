package e

import (
	"github.com/gin-gonic/gin"
	"github.com/steven1213/go-douyin/constant"
	"net/http"
)

type WrapperHandler func(c *gin.Context) (interface{}, error)

func ErrorWrapper(handler WrapperHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": constant.InternalServerError,
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
