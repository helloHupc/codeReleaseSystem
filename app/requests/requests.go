package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func doValidate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).ValidateStruct()
}

type ValidatorFunc func(reqData interface{}, ctx *gin.Context) map[string][]string

func ValidateRequest(c *gin.Context, reqData interface{}, validateFn ValidatorFunc) (isValid bool, errors map[string][]string) {
	if err := c.ShouldBindJSON(reqData); err != nil {
		return false, map[string][]string{"param_error": {err.Error()}}
	}
	errors = validateFn(reqData, c)
	return len(errors) == 0, errors
}
