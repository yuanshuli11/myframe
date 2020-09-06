package user

import "github.com/gin-gonic/gin"

type repository interface {
	//根据userId 获取用户信息
	GetUserById(c *gin.Context, userId string) (*User, error)
}
