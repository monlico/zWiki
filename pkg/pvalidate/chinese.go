package pvalidate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type CustomChineseValidator struct {
	Validator *validator.Validate
}

// Validate 验证方法
func (cv *CustomChineseValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func ValidateChinese(fl validator.FieldLevel) bool {
	// 中文正则表达式
	chineseRegex := regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	return chineseRegex.MatchString(fl.Field().String())
}
