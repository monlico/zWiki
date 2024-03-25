package wiki

import (
	"errors"
	"zWiki/model/mysql/wiki"
	"zWiki/model/redis"
	"zWiki/pkg/commonStruct"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
	"zWiki/pkg/setting"
	"zWiki/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type LoginService struct {
}

type LoginUserDetailParams struct {
	UserName string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}
type LoginUserParams struct {
	LoginUserDetailParams
	GroupName string `json:"group_name" validate:"chinese"  label:"组名"`
}

//登录接口
func (l *LoginService) Login(username, password, groupName, platform string) (string, int) {
	//判断user是否存在，如果不存在注册
	var (
		userModel  wiki.User
		groupModel wiki.Group
	)

	var (
		code   int = e.SUCCESS
		userId uint
		token  string
	)

	err := userModel.GetUserByCondition(map[string]interface{}{
		"username": username,
	})
	if err != nil {
		logging.Error(err)
		code = e.ERROR_MYSQL
	}
	if userModel.ID == 0 {
		userId, code = l.LogInUser(username, password)
	} else {
		userId = userModel.ID
	}

	err = groupModel.GetGroupByCondition(map[string]interface{}{
		"group_name": groupName,
	})

	if err != nil {
		logging.Error(err)
		code = e.ERROR_MYSQL
	}

	if groupModel.ID == 0 {
		_, code = l.LogInGroup(groupName, userId)
	}

	//登录，设置token,前面都成功才设置token,如果注册成功，或者密码正确？
	if code == e.SUCCESS {
		if userModel.Password != password {
			code = e.ERROR_PASSWORD
			return token, code
		}
		//密码正确，先删除user_id想同的token
		token, err = util.GenerateToken(username, platform, userId)
		if err != nil {
			logging.Error(err)
			code = e.ERROR_LOGIN_SET_TOKEN_FAIL
		}
	}

	return token, code
}

/**
 * GetGrou获取组名
 * @return 组名,error
 */
func (l *LoginService) GetGroup(username, password string) ([]*commonStruct.CommonKeyValueStr, int) {
	var ( //Model
		userModel wiki.User
	)
	var (
		code       int = e.SUCCESS
		returnData []*commonStruct.CommonKeyValueStr
	)

	err := userModel.GetUserByCondition(map[string]interface{}{
		"username": username,
	})

	if err != nil {
		logging.Error(err)
		code = e.ERROR_MYSQL
	}

	if userModel.Password != password {
		return returnData, e.ERROR_PASSWORD
	}

	for _, group := range userModel.Groups {
		returnData = append(returnData, &commonStruct.CommonKeyValueStr{
			Value: group.ID,
			Label: group.GroupName,
		})
	}

	return returnData, code
}

//注册用户
func (l *LoginService) LogInUser(username, password string) (uint, int) {
	var (
		UserModel wiki.User
		userId    uint
	)

	var code = e.SUCCESS

	//加密密码
	createId, createErr := UserModel.Create(username, password) //创建组
	if createErr != nil {
		logging.Error(createErr)
		code = e.ERROR_MYSQL
	}
	userId = createId

	return userId, code
}

//注册组
func (l *LoginService) LogInGroup(groupName string, userId uint) (uint, int) {
	var (
		groupModel wiki.Group
	)
	var (
		code    int  = e.SUCCESS
		groupId uint = 0
	)

	//判断是否能创建组
	isCanCreateNewGroup := l.VerifyCanUserCreate(userId)
	if !isCanCreateNewGroup { //超出权限，不能注册组
		code = e.ERROR_LOGIN_REGISTITION_GROUP_LIMIT
	} else {
		createId, createErr := groupModel.Create(groupName, userId) //创建组
		if createErr != nil {
			logging.Error(createErr)
			code = e.ERROR_MYSQL
		}
		groupId = createId
	}

	return groupId, code
}

/**
*判断用户是否能创建组
 */
func (l *LoginService) VerifyCanUserCreate(userId uint) bool {
	var ( //Model
		userModel wiki.User
	)

	err := userModel.GetUserByCondition(map[string]interface{}{
		"id": userId,
	})

	if err != nil {
		logging.Error(err)
	}

	if len(userModel.Groups) > 3 {
		return false
	}

	return true
}

/*
*登录时检查是否该用户已登录，删除相同userID的token
 */
func (l *LoginService) DeleteCommonUSerIdToken(c *gin.Context, userId uint) {
	var tmpTokenPre string = setting.UserLoginTokenPre + com.ToStr(userId) + "*"

	keys, err := redis.Redis.Keys(c, tmpTokenPre).Result() //只会有一个

	if err != nil {
		logging.Error(err)
	}

	if len(keys) > 0 {
		//按理来说只会有一个
		for _, v := range keys {
			_, err = redis.Redis.Del(c, v).Result()
			if err != nil {
				logging.Error(err)
			}
		}
	}
	if err != nil {
		logging.Error(errors.New("删除异地登录token失败"))
	}
}
