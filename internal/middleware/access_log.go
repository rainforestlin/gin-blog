package middleware

import (
	"bytes"
	"time"

	"github.com/julianlee107/blogWithGin/global"

	"github.com/julianlee107/blogWithGin/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 通过AccessLogWriter中的body取到值
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"response": bodyWriter.body.String(),
			"request":  c.Request.PostForm.Encode(),
		}

		s := "access log: method: %s, status_code:%d, begin_time: %d, end_time: %d"

		global.Logger.WithFields(fields).Infof(c, s, c.Request.Method, bodyWriter.Status(), beginTime, endTime)

	}
}
