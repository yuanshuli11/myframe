package token

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type SignRequest struct {
	accessKeyId   string
	nonce         string
	timestamp     string
	signedHeaders string
	principal     string
	signature     string

	parameters []string
}

func (r *SignRequest) addParameter(key, value string) bool {
	if key == "" || value == "" {
		return false
	}
	r.parameters = append(r.parameters, key+"="+value)
	return true
}
func (r *SignRequest) addDataString(req *http.Request) bool {
	//qs := req.URL.RawQuery
	rawQueryList := make([]string, 0)
	//	if qs != "" {
	//		rawQueries := strings.Split(qs, "&")
	//		for _, query := range rawQueries {
	//			if query != "" {
	//				rawQueryList = append(rawQueryList, query)
	//			}
	//		}
	//	}
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			if len(v) > 0 {
				rawQueryList = append(rawQueryList, k+"="+v[0])

			}
		}

	}
	if len(rawQueryList) > 0 {
		sort.Strings(rawQueryList)
		r.addParameter("data", strings.Join(rawQueryList, "&"))
	}
	return true
}
func ValidateRequest(c *gin.Context, appId, appKey string) error {
	timestamp, err := strconv.ParseInt(c.Request.Header.Get("timestamp"), 10, 64)
	if err != nil {
		return errors.New("签名时间戳非法")
	}
	nowTime := time.Now().Unix()

	if math.Abs(float64(nowTime-timestamp)) > 30000 {
		return errors.New("签名已过期")
	}
	serverSign, err := getSign(c, appId, appKey, timestamp)
	if err != nil {
		return err
	}

	clientSign := c.Request.Header.Get("sign")
	if clientSign == serverSign {
		return nil
	}
	return errors.New("签名计算失败" + serverSign)
}
func getSign(c *gin.Context, appId, appKey string, timestamp int64) (string, error) {
	signReq := &SignRequest{}
	signReq.addParameter("host", c.Request.Host)

	signReq.addParameter("method", c.Request.Method)
	if c.Request.URL.Path == "" {
		signReq.addParameter("path", "/")
	} else {
		signReq.addParameter("path", c.Request.URL.Path)
	}
	signReq.addParameter("app_id", appId)
	signReq.addParameter("app_secret", appKey)
	signReq.addParameter("timestamp", strconv.FormatInt(timestamp, 10))
	signReq.addDataString(c.Request)

	sort.Strings(signReq.parameters)
	h := md5.New()
	str := strings.Join(signReq.parameters, "&")
	h.Write([]byte(str))
	if c.Request.Header.Get("show_sign") == "true" {
		fmt.Fprintf(c.Writer, "%s", str)
	}
	return hex.EncodeToString(h.Sum(nil)), nil

}
