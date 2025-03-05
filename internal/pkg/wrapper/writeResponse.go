package wrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logUtils "github.com/thannaton/interview-golang/internal/pkg/logs"
)

func WriteResponse[T any](c *gin.Context, data T) {
	if c.Errors != nil {
		logUtils.Error.Println(c.Errors.String())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": c.Errors.String(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
