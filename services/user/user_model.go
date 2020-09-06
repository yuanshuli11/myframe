package user

type User struct {
	Id       int64  `json:""`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func GetNewUserModel(userId int64, name, phone, password string) *User {
	user := new(User)
	user.Id = userId
	user.Name = name
	user.Password = password
	user.Phone = phone
	return user

}
