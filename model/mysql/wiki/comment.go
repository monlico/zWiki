package wiki

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	ParentId   uint   `json:"parent_id"`
	Wid        uint   `json:"wid"`         //文章id
	RecoverUid uint   `json:"recover_uid"` //恢复的人id
	Endorse    uint   `json:"endorse"`
	Uid        uint   `json:"uid"`
	Text       string `json:"text"`
	//子集
	Children []*Comment `gorm:"foreignKey:parent_id;references:ID" json:"Children"`
}

func (c *Comment) Table() string {
	return "wiki_comment"
}

// 创建评论
func (c *Comment) Create() (uint, error) {
	err := db.Model(c).Create(c).Error
	if err != nil {
		return 0, err
	}
	return c.ID, nil
}

//评论列表
func (c *Comment) CommentList(articleId uint) ([]*Comment, error) {
	var res []*Comment

	err := db.Model(c).Debug().Preload("Children").
		Where(map[string]interface{}{
			"wid":       articleId,
			"parent_id": 0,
		}).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return res, nil
}
