package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	INVALID_PARAMS:                 "请求参数错误",
	ERROR_TOKEN_NO_EXIST:           "请先登录！",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token已过期，请重新登录",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token解析失败",

	//登录，用户相关错误码
	ERROR_LOGIN_REGISTITION_GROUP_LIMIT: "每位用户最多加入三个小组哟",
	ERROR_LOGIN_USER_UNEXISTS:           "用户不存在",
	ERROR_LOGIN_SET_TOKEN_FAIL:          "设置token失败",
	ERROR_PASSWORD:                      "密码错误",


	//数据验证错误
	ERROR_VALIDATOR:      "数据验证错误！请注意参数格式",
	ERROR_DATA_TRANSFORM: "数据转换出现错误",
	//MySQL 相关错误码
	ERROR_MYSQL: "数据库出现错误！请联系管理员修复！",
	ERROR_REDIS: "redis数据库出现错误！请联系管理员修复",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
