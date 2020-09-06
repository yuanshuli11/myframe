package user

import "github.com/gin-gonic/gin"

type cacheRepository struct {
}

func (cache *cacheRepository) GetUserById(c *gin.Context, userId string) (*User, error) {

	return nil, nil
}
