package sys

type ErrorType struct {
	ErrNum  int
	LogInfo string
	ErrTpl  string
}

var ErrorTypes = map[string]ErrorType{
	/// 400xx 接口校验错误
	"ERR_PARAMS": {400001, "参数错误", "缺少参数:%s"},

	//401xx user相关错误
	"ERR_GET_USER": {400101, "获取用户失败", "获取用户信息失败 user_id:%s"},
}
