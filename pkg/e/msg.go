package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	//登录，用户相关错误码
	ERROR_LOGIN_REGISTITION_GROUP_LIMIT: "每位用户最多加入三个小组哟",
	ERROR_LOGIN_USER_UNEXISTS:           "用户不存在",

	//MySQL 相关错误码
	ERROR_MYSQL: "数据库出现错误！请联系管理员修复！",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
