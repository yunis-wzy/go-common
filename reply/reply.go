package reply

import (
	"go-common/env"
	"net/http"

	"go-common/status"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int32       `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	LogID interface{} `json:"log_id"`
}

func RawReply(c *gin.Context, httpCode int, errCode int32, errMsg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code:  errCode,
		Msg:   errMsg,
		Data:  data,
		LogID: c.Value("K_LOGID"),
	})
}

func SuccessReply(c *gin.Context, data interface{}) {
	RawReply(c, http.StatusOK, 0, "", data)
}

func StatusReply(c *gin.Context, httpCode int, st *status.Status) {
	msg := st.Error()
	RawReply(c, httpCode, st.Code(), msg, nil)
}

func ErrorReply(c *gin.Context, httpCode int, err error) {
	RawReply(c, httpCode, status.CommonInternalError.Code(), err.Error(), nil)
}

func ErrorParameterReply(c *gin.Context, httpCode int, msg string) {
	if env.Env() == env.ENVBOE {
		RawReply(c, httpCode, status.CommonInternalError.Code(), msg, nil)
	} else {
		RawReply(c, httpCode, status.CommonInvalidParams.Code(), status.CommonInvalidParams.Error(), nil)
	}
}

func ErrorPenetrateReply(c *gin.Context, httpCode int, code int32, msg string, data interface{}) {
	RawReply(c, httpCode, code, msg, data)
}
