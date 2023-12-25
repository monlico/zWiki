package wiki

import (
	"net/http"
	"time"
	"zWiki/model/redis"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
	"zWiki/pkg/pvalidate"
	"zWiki/pkg/setting"
	"zWiki/pkg/util"
	"zWiki/services/wiki"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LoginGetGroupController(c *gin.Context) {

	var params wiki.LoginUserParams

	var (
		token string
		code  int
	)
	// 解析 JSON 请求体

	validate := validator.New()
	validate.RegisterValidation("chinese", pvalidate.ValidateChinese)

	// 将自定义验证器设置为默认验证器

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"body param error": err.Error()})
		return
	}

	err := validate.Struct(params)
	if err != nil {
		// 输出验证错误信息
		var errMsg string
		for _, er := range err.(validator.ValidationErrors) {
			errMsg += er.Error() + "/"
		}
		logging.Warn(errMsg)
		code = e.ERROR_VALIDATOR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": token,
		})
		return
	}

	var (
		username  string = params.UserName
		password  string = params.Password
		groupName string = params.GroupName
		platform  string = c.GetHeader("platform")
	)

	token, code = (&wiki.LoginService{}).Login(username, password, groupName, platform)

	if code == e.SUCCESS {
		tmpUserDetail, err := util.ParseToken(token)

		if err != nil {
			code = e.ERROR_LOGIN_SET_TOKEN_FAIL
		}
		var tokenKey string = setting.UserLoginTokenPre + token

		if err == nil {
			_, err := redis.Redis.Set(c, tokenKey, token, time.Duration(tmpUserDetail.ExpiresAt)).Result()
			if err != nil {
				code = e.ERROR_REDIS
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": token,
	})
}
