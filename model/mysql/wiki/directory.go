package wiki

import "gorm.io/gorm"

type Directory struct {
	gorm.Model

	ParentId uint   `json:"parent_id"`
	Name     string `json:"name"`
	Gid      uint   `json:"gid"`
	Uid      uint   `json:"uid"`
}

type DirectoryItem struct {
	Directory
	Wikis []ArticlesPart1 `gorm:"-" json:"wikis"`
}

//目录列表，只用返回文章的title就够了
func (d *Directory) GetAllDirectoryByGid(gid uint) ([]DirectoryItem, error) {
	var res []DirectoryItem
	err := db.Model(d).Debug().Where(map[string]interface{}{
		"gid": gid,
	}).Find(&res).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return res, nil
}

//添加目录
func (d *Directory) AddDirectory() error {
	err := db.Debug().Create(d).Error
	if err != nil {
		return err
	}
	return nil
}
