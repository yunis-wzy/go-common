package middleware

import (
	"context"
	"go-common/logs"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func ServerRequestLog(ignoreURIs ...string) gin.HandlerFunc {
	ignoreMethodsSet := map[string]struct{}{}
	for _, key := range ignoreURIs {
		ignoreMethodsSet[key] = struct{}{}
	}
	return func(c *gin.Context) {
		URI := c.Request.RequestURI
		if _, ok := ignoreMethodsSet[string(URI)]; ok {
			c.Next()
			return
		}
		// 记录请求参数
		startTime := time.Now()
		c.Next()
		// 记录响应结果.
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		respData, _ := c.Writer.WriteString("")
		useTime := float64(time.Since(startTime).Microseconds()) / 1000
		logs.CtxInfo(context.TODO(), "[hertz_server_log]:method:{%s},URI:{%s},reqHeader:{%s},reqBody:{%s},resBody:{%s},latency(ms):{%s}",
			string(c.Request.Method),
			string(c.Request.RequestURI),
			c.Request.Header,
			string(reqBody),
			respData,
			useTime,
		)
	}
}