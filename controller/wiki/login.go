package wiki

import (
	"time"
	"zWiki/model/redis"
	"zWiki/pkg/commonStruct"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
	"zWiki/pkg/pvalidate"
	"zWiki/pkg/returnMsg"
	"zWiki/pkg/setting"
	"zWiki/pkg/util"
	"zWiki/services/wiki"

	"github.com/gin-gonic/gin"
)

//获取组名
// @route: /login/getGroup
func LoginGetGroupController(c *gin.Context) {

	var params wiki.LoginUserDetailParams

	var (
		groups []*commonStruct.CommonKeyValueStr
		code   int
	)

	// 将自定义验证器设置为默认验证器
	if err := c.ShouldBindJSON(&params); err != nil {
		code = e.INVALID_PARAMS
		returnMsg.ReturnFailedMsg(code, "body param error:"+err.Error(), c)
		return
	}

	err := pvalidate.WikiValidator.Struct(params)
	if err != nil {
		// 输出验证错误信息
		var errMsg string
		errMsg = pvalidate.Translate(err)
		logging.Warn(errMsg)
		code = e.ERROR_VALIDATOR
		returnMsg.ReturnSuccessMsg(code, errMsg, c)
		return
	}

	var (
		username string = params.UserName
		password string = params.Password
	)

	groups, code = (&wiki.LoginService{}).GetGroup(username, password)

	returnMsg.ReturnSuccessMsg(code, groups, c)
	return
}

//登录
// @route: /login/index
func LoginController(c *gin.Context) {

	var params wiki.LoginUserParams

	var (
		token string
		code  int
	)

	// 将自定义验证器设置为默认验证器
	if err := c.ShouldBindJSON(&params); err != nil {
		code = e.INVALID_PARAMS
		returnMsg.ReturnFailedMsg(code, "body param error:"+err.Error(), c)
		return
	}

	err := pvalidate.WikiValidator.Struct(params)
	if err != nil {
		// 输出验证错误信息
		var errMsg string
		errMsg = pvalidate.Translate(err)
		logging.Warn(errMsg)
		code = e.ERROR_VALIDATOR
		returnMsg.ReturnSuccessMsg(code, errMsg, c)
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
	returnMsg.ReturnSuccessMsg(code, token, c)
	return
}
