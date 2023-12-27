package pvalidate

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

// 自定义密码验证器
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// 密码至少包含一个数字
	hasNumber := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasNumber = true
			break
		}
	}

	// 密码至少包含一个小写字母
	hasLower := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
			break
		}
	}

	// 密码至少包含一个大写字母
	hasUpper := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
			break
		}
	}

	// 密码同时包含数字、大小写字母
	return hasNumber && hasLower && hasUpper
}
