package wiki

import (
	models "zWiki/model/mysql"

	"gorm.io/gorm"
)

var db = models.Db

type User struct {
	gorm.Model

	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码

	//关联表
	Groups []*Group `gorm:"many2many:user_groups;"` //组
}

func (wu *User) GetUserByCondition(where map[string]interface{}) error {
	err := db.Model(&User{}).Where(where).Scan(wu).Error

	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}
	return nil
}
