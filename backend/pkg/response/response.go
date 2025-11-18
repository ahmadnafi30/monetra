package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"isSuccess"`
}

func Success(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Data:    data,
	})
}

func ErrorCtx(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Data:    err.Error(),
	})
}

func Error(code int, message string, err error) error {
	if err != nil {
		return fmt.Errorf("[%d] %s: %v", code, message, err)
	}
	return fmt.Errorf("[%d] %s", code, message)
}
