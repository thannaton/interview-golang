package mdw

import (
	"io"

	"bytes"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	logUtils "github.com/thannaton/interview-golang/internal/pkg/logs"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read Request Header
		reqHeader := c.Request.Header

		// Read Request URL
		reqURL := c.Request.URL

		// Read Method
		reqMethod := c.Request.Method

		// Read request body
		reqBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Error(err)
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		restApiRequestLogger(reqHeader, reqURL.String(), reqMethod, reqBody)

		// Wrap the response writer
		respBodyWriter := &bodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = respBodyWriter

		// Process request
		c.Next()
		
		// Read Response Header
		respHeader := c.Writer.Header()

		// Read Response Status
		respStatus := c.Writer.Status()

		// Read Response Body
		respBody := respBodyWriter.body.Bytes()

		if c.Writer.Status() >= 200 && c.Writer.Status() < 300 {
			restApiResponseLogger(respHeader, reqURL.String(), reqMethod, respStatus, respBody)
		} else {
			restApiResponseErrorLogger(respHeader, reqURL.String(), reqMethod, respStatus, respBody)
		}
	}
}

func restApiRequestLogger(reqHeader map[string][]string, reqURL string, reqMethod string, reqBody []byte) {
	format := "[%v][%v: %v][%v: %v][%v: %v][%v: %v]"
	logUtils.Info.Printf(format, color.GreenString("Rest API Request"), color.GreenString("Method"), reqMethod, color.GreenString("Header"), reqHeader, color.GreenString("URL"), reqURL, color.GreenString("Body"), string(reqBody))
}

func restApiResponseLogger(respHeader map[string][]string, respURL string, respMethod string, httpStatus int, respBody []byte) {
	format := "[%v][%v: %v][%v: %v][%v: %v][%v: %v][%v: %v]"
	logUtils.Info.Printf(format, color.GreenString("Rest API Response"), color.GreenString("Method"), respMethod, color.GreenString("HttpStatus"), httpStatus, color.GreenString("Header"), respHeader, color.GreenString("URL"), respURL, color.GreenString("Body"), string(respBody))
}

func restApiResponseErrorLogger(respHeader map[string][]string, respURL string, respMethod string, httpStatus int, respBody []byte) {
	format := "[%v][%v: %v][%v: %v][%v: %v][%v: %v][%v: %v]"
	logUtils.Error.Printf(format, color.GreenString("Rest API Response"), color.GreenString("Method"), respMethod, color.GreenString("HttpStatus"), httpStatus, color.GreenString("Header"), respHeader, color.GreenString("URL"), respURL, color.GreenString("Body"), string(respBody))
}
