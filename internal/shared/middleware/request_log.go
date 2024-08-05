package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type ResponseBodyWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w ResponseBodyWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		var requestBodyBytes []byte
		if c.Request().Body != nil {
			requestBodyBytes, _ = io.ReadAll(c.Request().Body)
			c.Request().Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
		}

		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &ResponseBodyWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		err := next(c)

		duration := time.Since(start)

		logEvent := log.Info()
		msg := "success"
		if err != nil {
			msg = "error"
			logEvent = log.Error()
		}
		logEvent.
			Str("method", c.Request().Method).
			Str("path", c.Path()).
			Str("uri", c.Request().RequestURI).
			Str("ip", c.RealIP()).
			Str("user_agent", c.Request().UserAgent()).
			Int("status", c.Response().Status).
			Interface("params", c.QueryParams()).
			Interface("request", convertBodyToInterface(requestBodyBytes)).
			Interface("response", convertBodyToInterface(resBody.Bytes())).
			Dur("duration", duration).Msg(msg)

		return err
	}
}

func convertBodyToInterface(body []byte) any {
	var parsedResponseBody any
	if json.Unmarshal(body, &parsedResponseBody) != nil {
		parsedResponseBody = string(body)
	}
	return parsedResponseBody
}
