package wiki

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	ParentId  uint      `json:"parent_id"`
	ArticleId uint      `json:"article_id"` //文章id
	RecoverId uint      `json:"recover_id"` //恢复的人id
	Endorse   uint      `json:"endorse"`
	Uid       uint      `json:"uid"`
	Text      string    `json:"text"`
	Comments  []Comment `gorm:"foreignkey:parent_id;references:ID" json:"comments"`
}

// 创建评论
func (c *Comment) Create() (uint, error) {
	err := db.Model(c).Create(c).Error
	if err != nil {
		return 0, err
	}
	return c.ID, nil
}

// 评论列表

func (c *Comment) CommentList(articleId uint) (*Comment, error) {
	err := db.Model(c).Preload("Comments").Where(map[string]interface{}{
		"article_id": articleId,
	}).Scan(c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return c, nil
}
