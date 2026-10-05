/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 10:34 PM
 * Description:
 **/

package xlog

import (
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"github.com/weitrue/Seckill/pkg/logger/meta"
	"github.com/weitrue/Seckill/pkg/logger/xzap"

	"github.com/gin-gonic/gin"
)

func GinInterceptor(zapLogger *xzap.ZapLogger, msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		requestBody, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)
		bodyLogWriter := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		start := time.Now()

		c.Next()

		responseBody := bodyLogWriter.body.Bytes()
		// log := WithContext(c.Request.Context())
		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors {
				zapLogger.Error("msg", e)
			}
		} else {
			zapLogger.Info(msg,
				meta.NewField("status", c.Writer.Status()),
				meta.NewField("method", c.Request.Method),
				meta.NewField("function", c.HandlerName()),
				meta.NewField("path", path),
				meta.NewField("query", query),
				meta.NewField("ip", c.ClientIP()),
				meta.NewField("user-agent", c.Request.UserAgent()),
				meta.NewField("token", c.Request.Header.Get("session_id")),
				meta.NewField("content-type", c.Request.Header.Get("Content-Type")),
				meta.NewField("latency", float64(time.Now().Sub(start).Nanoseconds()/1000000.0)),
				meta.NewField("request", string(requestBody)),
				meta.NewField("response", string(responseBody)),
			)
		}
	}
}

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
