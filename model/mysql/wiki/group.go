package wiki

import "gorm.io/gorm"

type Group struct {
	gorm.Model

	GroupName string
}

func (g *Group) GetGroupByCondition(where map[string]interface{}) error {
	err := db.Model(&User{}).Where(where).Scan(g).Error

	if err != nil {
		return err
	}
	return nil
}
