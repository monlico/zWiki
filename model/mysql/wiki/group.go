package wiki

import "gorm.io/gorm"

type Group struct {
	gorm.Model

	GroupName string `json:"group_name"`

	//关联用户
	Users []*User `gorm:"many2many:user_groups;"`
}

func (g *Group) GetGroupByCondition(where map[string]interface{}) error {
	err := db.Model(&User{}).Where(where).Scan(g).Error

	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}
	return nil
}

//创建组
func (g *Group) Create(groupName string, userId uint) (uint, error) {
	var (
		NewUser User
		groupId uint
	)
	err := NewUser.GetUserByCondition(map[string]interface{}{
		"id": userId,
	})

	if err != nil {
		return groupId, err
	}
	g.GroupName = groupName
	g.Users = []*User{&NewUser}

	err = db.Model(&Group{}).Create(g).Error
	if err != nil {
		return groupId, err
	}
	groupId = g.ID
	return groupId, nil
}
