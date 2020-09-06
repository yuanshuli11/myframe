package middleware

import (
	"myproject/infrastructure/mlog"
	"time"

	"github.com/gin-gonic/gin"
)

//RequestInfo 生成access日志中间件
func RequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {

		//获取traceId 作为请求的唯一ID
		uniqid := mlog.GetUniqid(c)
		c.Set("uniqid", uniqid)
		beforeTime := time.Now().UnixNano()
		c.Request.ParseForm()
		c.Next()
		time.Sleep(1 * time.Second)
		afterTime := time.Now().UnixNano()
		//这里以毫秒为单位 低于1ms的可能因为四舍五入而显示为0
		cost := int((afterTime - beforeTime) / 1000000)
		mlog.Access(c, cost)
	}
}
