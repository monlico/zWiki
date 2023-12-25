package pvalidate

import (
	"zWiki/pkg/logging"

	"github.com/go-playground/validator/v10"
)

var WikiValidator *validator.Validate

func init() {
	WikiValidator = validator.New()
	//注册自定义验证器

	err := WikiValidator.RegisterValidation("chinese", ValidateChinese)

	if err != nil {
		logging.Error(err)
	}
}

//输出错误信息
func PrintValidateErr(err error) string {

}
