package user

import (
	"myproject/infrastructure/database"
	"myproject/infrastructure/mlog"
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id       int64  `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	CreateAt time.Time
	UpdateAt time.Time
}

func (u *user) TableName() string {
	return "users"
}

type dbRepository struct {
}

func (repo *dbRepository) GetUserById(c *gin.Context, userId string) (*User, error) {

	userData := new(user)
	r := database.DB.Select("*").Where("id = ?", userId).Find(&userData)

	if r.Error != nil {
		mlog.ErrorCtxf(c, "GetUserById error:%v", r.Error)
		return nil, r.Error
	}
	user := GetNewUserModel(userData.Id, userData.Name, userData.Phone, userData.Password)
	return user, nil

}
