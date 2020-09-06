package mlog

import (
	"fmt"
	"net"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logClient *logrus.Logger
var logLevels = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
}

func init() {
	logClient = logrus.New()
	logClient.Formatter = &logrus.JSONFormatter{}
}

func GetLogrus() *logrus.Logger {
	return logClient

}
func ConfigLocalFilesystemLogger(logPath, loglevel string) {
	logFileName := "logs"
	baseLogPaht := path.Join(logPath, logFileName)
	level, ok := logLevels[loglevel]
	if ok {
		logClient.SetLevel(level)
	} else {
		logClient.SetLevel(logrus.WarnLevel)
	}

	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),   // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*1200),  // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error", err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)
}

//打印访问日志
func Access(c *gin.Context, cost int) {

	logClient.WithFields(logrus.Fields{
		"type":       "access",
		"request_id": GetUniqid(c),
		"user_ip":    GetClientIp(c),
		"cost":       strconv.Itoa(cost),
		"host":       c.Request.URL.Host,
		"uri":        c.Request.RequestURI,
		"method":     c.Request.Method,
		"http_code":  c.Writer.Status(),
	}).Info("")

}

func DebugCtxf(c *gin.Context, format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)

	logClient.WithFields(logrus.Fields{
		"type":       "app",
		"request_id": GetUniqid(c),
		"user_ip":    GetClientIp(c),
	}).Debug(str)

}

func InfoCtxf(c *gin.Context, format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)

	logClient.WithFields(logrus.Fields{
		"type":       "app",
		"request_id": GetUniqid(c),
		"user_ip":    GetClientIp(c),
	}).Info(str)

}

func WarnCtxf(c *gin.Context, format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)

	logClient.WithFields(logrus.Fields{
		"type":       "app",
		"request_id": GetUniqid(c),
		"user_ip":    GetClientIp(c),
	}).Warn(str)
}

func ErrorCtxf(c *gin.Context, format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)

	logClient.WithFields(logrus.Fields{
		"type":       "app",
		"request_id": GetUniqid(c),
		"user_ip":    GetClientIp(c),
	}).Error(str)
}

func GetClientIp(c *gin.Context) (ip string) {

	if c.Request == nil {
		return ""
	}
	ipCache, exists := c.Get("HTTP_USER_IP")
	if exists {
		return ipCache.(string)
	}
	remoteAddr := c.Request.RemoteAddr
	if ip := c.Request.Header.Get("HTTP-X-FORWARDED-FOR"); ip != "" {
		remoteAddr = ip
	} else if ip = c.Request.Header.Get("X-FORWARDED-FOR"); ip != "" {
		remoteAddr = ip
	} else if ip = c.Request.Header.Get("HTTP-CLIENT-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = c.Request.Header.Get("CLIENT-IP"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	index := strings.Index(remoteAddr, ",")
	if index != -1 {
		remoteAddr = remoteAddr[0:index]
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	clientIp := net.ParseIP(remoteAddr)
	if clientIp != nil {
		ip = clientIp.String()
	} else {
		ip = ""
	}

	c.Set("HTTP_USER_IP", ip)
	return ip
}
func GetUniqid(c *gin.Context) (uniqid string) {
	id, exists := c.Get("uniqid")
	if !exists {
		uniqid = ""
	} else {
		uniqid = id.(string)
	}
	return uniqid
}
