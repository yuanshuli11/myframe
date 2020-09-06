package user

import (
	"myproject/infrastructure/sys"
	"myproject/services/user"

	"github.com/gin-gonic/gin"
)

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("user_id")
		if userId == "" {
			sys.ReturnError(c, "ERR_PARAMS", "user_id")
			return
		}

		u, e := user.UserService.GetUserInfo(c, userId)
		if e != nil {
			sys.ReturnError(c, "ERR_GET_USER", userId)
			return
		}

		sys.ReturnSuccess(c, u)
		return
	}
}
