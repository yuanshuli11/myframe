package app

import (
	"myproject/app/controller/user"
	"myproject/app/middleware"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()

	engine.Use(middleware.RequestInfo())
}

func GetRouters() *gin.Engine {

	authRouter := engine.Group("/user", middleware.VerifySign())
	{

		authRouter.GET("/get", user.GetUserInfo())

	}
	return engine
}
