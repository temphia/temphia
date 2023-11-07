package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

type ginHands struct {
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	ClientIP   string
	MsgStr     string
}

func (m *Middleware) Log(c *gin.Context) {
	t := time.Now()
	// before request
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	pp.Println("@Req-host", c.Request.Host)
	pp.Println("@Url-host", c.Request.URL.Host)

	// time.Sleep(2 * time.Second)

	c.Next()
	// after request
	// latency := time.Since(t)
	// clientIP := c.ClientIP()
	// method := c.Request.Method
	// statusCode := c.Writer.Status()
	if raw != "" {
		path = path + "?" + raw
	}
	msg := c.Errors.String()
	if msg == "" {
		msg = "Request"
	}

	data := &ginHands{
		Path:       path,
		Latency:    time.Since(t),
		Method:     c.Request.Method,
		StatusCode: c.Writer.Status(),
		ClientIP:   c.ClientIP(),
		MsgStr:     msg,
	}

	switch {
	case data.StatusCode >= 400 && data.StatusCode < 500:
		{

			m.Logger.Warn().Str("ser_name", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("client_ip", data.ClientIP).Msg(data.MsgStr)
		}
	case data.StatusCode >= 500:
		{
			m.Logger.Error().Str("ser_name", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("client_ip", data.ClientIP).Msg(data.MsgStr)
		}
	default:
		m.Logger.Info().Str("ser_name", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("client_ip", data.ClientIP).Msg(data.MsgStr)
	}

}
