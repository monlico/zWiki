package wiki

import "gorm.io/gorm"

type Articles struct {
	gorm.Model

	Uid uint `json:"uid"`
}
