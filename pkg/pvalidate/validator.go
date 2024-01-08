package pvalidate

import (
	"reflect"
	"zWiki/pkg/logging"

	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var WikiValidator *validator.Validate

// Validate/v10 全局验证器
var Trans ut.Translator

func init() {
	WikiValidator = validator.New()
	zhLocale := zh_Hans_CN.New()
	uni := ut.New(zhLocale, zhLocale)

	Trans, _ = uni.GetTranslator("zh_Hans_CN")
	//绑定自定义验证规则
	BindCustomizedValidate()

	//注册自定义验证器
	err := zhTrans.RegisterDefaultTranslations(WikiValidator, Trans)
	if err != nil {
		logging.Error("err:", err)
	}

	// 注册自定义验证器，获取json的字段名
	WikiValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		if name == "" {
			name = fld.Name
		}
		return name
	})

}
func registrationFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(fe.Tag(), fe.Field())
	return t
}

// 检验并返回检验错误信息
func Translate(err error) (errMsg string) {
	errs := err.(validator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(Trans)
	}
	return
}

//翻译设置，覆盖也在这里
func registerTranslations(ut ut.Translator) error {
	if err := ut.Add("chinese", "{0} 必须只包含中文", false); err != nil {
		return err
	}
	if err := ut.Add("password", "{0} 必须包含大小写字母，数字", false); err != nil {
		return err
	}
	return nil
}

//绑定自定义验证规则
func BindCustomizedValidate() {
	var err error
	err = WikiValidator.RegisterValidation("chinese", ValidateChinese) //注册中文验证
	if err != nil {
		logging.Error(err)
	}
	err = WikiValidator.RegisterTranslation("chinese", Trans, registerTranslations, registrationFunc) //注册中文验证报错
	if err != nil {
		logging.Error(err)
	}

	err = WikiValidator.RegisterValidation("password", ValidatePassword) //注册密码验证
	if err != nil {
		logging.Error(err)
	}
	err = WikiValidator.RegisterTranslation("password", Trans, registerTranslations, registrationFunc) //注册密码验证报错
	if err != nil {
		logging.Error(err)
	}

}
