package response

import (
	"net/http"
	"oa/internal/model/common"
	"oa/internal/util/err"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`           // 响应码
	Msg  string `json:"msg"`            // 响应消息
	Data any    `json:"data,omitempty"` // 响应数据
	Err  error  `json:"-"`              // 错误信息，序列化忽略
}

func Execute(fn func(ctx *gin.Context) Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := fn(ctx)
		if res.Err != nil {
			ctx.JSON(http.StatusInternalServerError, common.Fail[string](
				err.HandleError(
					&common.ServerErr{
						Msg: res.Msg,
						Err: res.Err,
					},
				),
			))
			return
		}

		// 否则说明无异常
		ctx.JSON(http.StatusOK, common.Success(res.Data, res.Msg))
	}
}
