package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"login/app/schema/response"
	"net/http"
)

const (
	ParamFormatError = "参数格式错误"
	ParamTypeError   = "参数格式错误"
)

func Validate(ctx *gin.Context,p interface{})  error{
	if err := ctx.ShouldBind(p); err != nil {
		response.Failed(ctx, http.StatusBadRequest, ParamFormatError)
		return err
	}

	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		response.Failed(ctx, http.StatusBadRequest, ParamTypeError)
		return err
	}
	return nil
}