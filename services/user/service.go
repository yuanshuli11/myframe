package user

import "github.com/gin-gonic/gin"

type service struct {
	Db    repository
	Cache repository
}

var (
	UserService *service
)

func init() {
	UserService = NewService()
}

func NewService() *service {
	service := new(service)

	service.Db = new(dbRepository)
	service.Cache = new(cacheRepository)
	return service
}
func (s *service) GetUserInfo(c *gin.Context, userId string) (*User, error) {

	//先从缓存中获取数据
	cache, _ := s.Cache.GetUserById(c, userId)
	//若缓存非空，则直接返回
	if cache != nil {
		return cache, nil
	}

	//从DB中获取数据
	user, e := s.Db.GetUserById(c, userId)
	if e != nil {
		return nil, e
	}
	return user, nil

}
