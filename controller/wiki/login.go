package wiki

import (
	"net/http"
	"time"
	"zWiki/model/redis"
	"zWiki/pkg/e"
	"zWiki/pkg/setting"
	"zWiki/pkg/util"
	"zWiki/services/wiki"

	"github.com/gin-gonic/gin"
)

func LoginGetGroupController(c *gin.Context) {

	var params wiki.LoginUserParams

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var (
		username  string = params.UserName
		password  string = params.Password
		groupName string = params.GroupName
		platform  string = c.GetHeader("platform")
	)

	var (
		token string
		code  int
	)

	token, code = (&wiki.LoginService{}).Login(username, password, groupName, platform)

	tmpUserDetail, err := util.ParseToken(token)

	if err != nil {
		code = e.ERROR_LOGIN_SET_TOKEN_FAIL
	}
	var tokenKey string = setting.UserLoginTokenPre + token

	if code == e.SUCCESS {
		redis.Redis.Set(c, tokenKey, token, time.Duration(tmpUserDetail.ExpiresAt))
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": token,
	})
}
