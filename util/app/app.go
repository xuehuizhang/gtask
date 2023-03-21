package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppG struct {
	C *gin.Context
}

func (a AppG) ResponseOk(data interface{}) {
	a.C.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func (a AppG) ResponseError(msg string) {
	a.C.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
		"data": "",
	})
}

func (a AppG) ResponseCode(code int, msg string, data interface{}) {
	a.C.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
