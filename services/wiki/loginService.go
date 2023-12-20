package wiki

import (
	"zWiki/model/mysql/wiki"

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
func (l *LoginService) GetGroup(params *LoginUserDetailParams) (string, error) {
	var ( //Model
		userModel wiki.User
	)

	err := userModel.GetUserByCondition(map[string]interface{}{
		"username": params.UserName,
	})

	if err != nil {
		return "", err
	}

	return userModel.Username, nil
}

func (l *LoginService) LogInGroup(groupName string) (string, int) {
	var (
		groupModel wiki.Group
		code       int
	)
	err := groupModel.GetGroupByCondition(map[string]interface{}{
		"group_name": groupName,
	})

	if err == gorm.ErrRecordNotFound { //找不到数据
		//判断加入了几个组
		groupModel.create() //创建组
	} else if err != nil {
		return "", err
	}
	return groupModel.GroupName, nil
}
