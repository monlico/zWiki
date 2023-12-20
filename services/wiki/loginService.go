package wiki

import (
	"zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"

	"gorm.io/gorm"
)

type LoginService struct {
}

type LoginUserDetailParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type LoginUserParams struct {
	LoginUserDetailParams
	GroupName string `json:"group_name"`
}

/**
 * GetGrou获取组名
 * @return 组名,error
 */
func (l *LoginService) GetGroup(params *LoginUserDetailParams) (string, int) {
	var ( //Model
		userModel wiki.User
	)
	var (
		code   = e.SUCCESS
		groups string
	)

	err := userModel.GetUserByCondition(map[string]interface{}{
		"username": params.UserName,
	})

	if err != nil {
		code = e.ERROR_MYSQL
	}

	for _, group := range userModel.Groups {
		groups += group.GroupName
	}

	return groups, code
}

//注册用户
func (l *LoginService) LogInUser(username, password string) int {
	//var (
	//	UserModel wiki.User
	//)

	var code = e.SUCCESS

	return code
}

//注册组
func (l *LoginService) LogInGroup(groupName string, userId int) int {
	var (
		groupModel wiki.Group
		code       int = e.SUCCESS
	)

	err := groupModel.GetGroupByCondition(map[string]interface{}{
		"group_name": groupName,
	})

	if err == gorm.ErrRecordNotFound { //找不到数据
		//判断加入了几个组
		isCanCreateNewGroup := l.VerifyCanUserCreate(userId)
		if !isCanCreateNewGroup { //超出权限，不能注册组
			code = e.ERROR_LOGIN_REGISTITION_GROUP_LIMIT
		} else {
			createErr := groupModel.Create(groupName, userId) //创建组
			if createErr != nil {
				code = e.ERROR_MYSQL
			}
		}
	} else if err != nil {
		logging.Error(err)
		code = e.ERROR_MYSQL
	}
	return code
}

/**
*判断用户是否能创建组
 */
func (l *LoginService) VerifyCanUserCreate(userId int) bool {
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
