package user

import (
	_ "myproject/test"
	"testing"

	"github.com/gin-gonic/gin"
)

var userTestService *service
var ctx *gin.Context

func init() {
	userTestService = NewService()
	ctx = new(gin.Context)

}

func TestGetuserInfo(t *testing.T) {

	_, e := userTestService.GetUserInfo(ctx, "1")
	if e != nil {
		t.Errorf("GetuserInfo() error = %v", e)
	}

}
