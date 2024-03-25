package wiki

import "gorm.io/gorm"

type Record struct {
	gorm.Model

	Wid      uint   `json:"wid"`
	Uid      uint   `json:"uid"`
	Username string `json:"username"`
	Type     int    `json:"type"`
}

// 创建记录
func (r *Record) Create() (uint, error) {
	err := db.Model(r).Create(r).Error
	if err != nil {
		return 0, err
	}
	return r.ID, nil
}

//记录列表
func (r *Record) RecordList(wid uint, page, limit int) ([]*Record, error) {
	var (
		res []*Record
		err error
	)
	db := db.Model(r).Debug().
		Where(map[string]interface{}{
			"wid": wid,
		})
	if limit == -1 {
		err = db.Scan(&res).Error
	} else {
		err = db.Offset((page - 1) * limit).Limit(limit).
			Scan(&res).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return res, nil
}
