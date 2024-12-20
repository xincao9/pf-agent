package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pf-agent/internal/logger"
	"reflect"
)

func RenderJSON(c *gin.Context, code int, message string) {
	if code != http.StatusOK {
		logger.L.Error(fmt.Sprintf("code =%d, message = %v", code, message))
	}
	c.JSON(code,
		gin.H{
			"code":    code,
			"message": message,
		})
}

func RenderJSONDetail(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code,
		gin.H{
			"code":    code,
			"message": message,
			"data":    data,
		})
}

func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}
