package middleware

import (
	"myproject/infrastructure/sys"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func VerifySign() gin.HandlerFunc {
	return func(c *gin.Context) {
		appId := c.Request.Header.Get("app_id")
		if appId == "" {
			sys.ReturnError(c, "ERR_LOGIN_NO_APPID")
			c.Abort()
			return
		}
		//		err := token.ValidateRequest(c, appId, "bbdaa965bdbea8551553aa28afb6025b")
		//
		//		if err != nil {
		//			sys.ReturnErrorMsg(c, "ERR_LOGIN_ERROR_SIGN", err.Error())
		//			c.Abort()
		//			return
		//		}
		logrus.Info("鉴权成功")
		//c.Next()
	}
}
