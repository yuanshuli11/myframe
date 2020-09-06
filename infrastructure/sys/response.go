package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func ReturnSuccess(c *gin.Context, data interface{}) {
	resp := JSONResponse{0, "success", data}
	c.JSON(0, resp)
}

func ReturnError(c *gin.Context, errorType string, v ...interface{}) {
	newErr, ok := ErrorTypes[errorType]
	format := "未知错误"
	if ok {
		format = newErr.ErrTpl
	}
	errstr := fmt.Sprintf(format, v...)
	resp := JSONResponse{ErrorTypes[errorType].ErrNum, errstr, nil}
	c.JSON(0, resp)
}
