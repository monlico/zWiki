package wiki

import (
	models "zWiki/model/mysql"

	"gorm.io/gorm"
)

var db = models.Db

type User struct {
	gorm.Model

	Username string // 用户名
	Password string // 密码
}

func (wu *User) GetUserByCondition(where map[string]interface{}) error {
	err := db.Model(&User{}).Where(where).Scan(wu).Error

	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}
	return nil
}
