package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidateError struct {
	Key     string
	Message string
}

type ValidateErrors []*ValidateError

func (v *ValidateError) Error() string {
	return v.Message
}

func (v ValidateErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidateErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// BindAndValid 对入参校验ShouldBind进行二次封装,如果发生错误，在通过中间件Translations中设置的Translator对错误消息体进行具体的翻译
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidateErrors) {
	var errs ValidateErrors
	err := c.ShouldBind(v)

	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidateError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
