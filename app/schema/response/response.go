package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type success struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Content interface{} `json:"content"`
}

type failed struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message interface{} `json:"msg"`
}


func Success(c *gin.Context, code int, content interface{}) {
	c.JSON(http.StatusOK, success{
		Code:    code,
		Success: true,
		Content: content,
	})
}

func Failed(c *gin.Context, code int, msg interface{}) {
	c.JSON(http.StatusOK, failed{
		Code:    code,
		Success: false,
		Message: msg,
	})
}
