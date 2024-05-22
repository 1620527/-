package controllers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Count   int         `json:"count"`
}

func ReturnSuccess(c *gin.Context, code int, message interface{}, data interface{}, count int) {
	json := JsonStruct{Code: code, Message: message, Data: data, Count: count}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, message interface{}) {
	json := JsonStruct{Code: code, Message: message, Data: nil, Count: 0}
	c.JSON(200, json)
}

func EncryMd5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
