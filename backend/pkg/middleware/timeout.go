package middleware

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ahmadnafi30/monetra/backend/pkg/response"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func (m *middleware) Timeout() gin.HandlerFunc {
	timeLimit, _ := strconv.Atoi(os.Getenv("TIME_OUT_LIMIT"))

	return timeout.New(
		timeout.WithTimeout(time.Duration(timeLimit)*time.Second),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	response.ErrorCtx(c, http.StatusRequestTimeout, "Request Time Out", errors.New(""))
}
