package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/steven1213/go-douyin/constant"
	"net/http"
)

func Resp(context *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	context.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// RespSuccess 成功响应
func RespSuccess(context *gin.Context, data interface{}) {
	Resp(context, http.StatusOK, constant.ApiOK, "success", data)
}

// RespError 失败响应
func RespError(context *gin.Context, code int, msg string) {
	Resp(context, http.StatusOK, code, msg, nil)
}

// RespPage 分页响应
func RespPage(context *gin.Context, data interface{}) {
	Resp(context, http.StatusOK, constant.ApiOK, "success", data)
}
