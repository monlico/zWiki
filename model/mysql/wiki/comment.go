package wiki

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	ParentId   uint   `json:"parent_id"`
	Wid        uint   `json:"wid"`         //文章id
	RecoverUid uint   `json:"recover_uid"` //恢复的人id
	Endorse    uint   `json:"endorse"`
	Quote      string `json:"quote"` //引用的内容
	Uid        uint   `json:"uid"`
	Text       string `json:"text"`
}
type CommentJoin struct {
	Comment
	Children []*Comment `gorm:"foreignKey:parent_id;references:ID" json:"Children"`
}

type CommentItem struct {
	Comment
	Name     string         `json:"name"`
	Parent   string         `json:"parent"`                                              //被回复的人的用户名
	Children []*CommentItem `gorm:"foreignKey:parent_id;references:ID"  json:"children"` //回复人的用户名
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
func (c *Comment) CommentList(articleId, page, limit uint) ([]*CommentItem, error) {
	var res []*CommentItem
	err := db.Model(c).Debug().
		Where(map[string]interface{}{
			"wid":       articleId,
			"parent_id": 0,
		}).Offset(int((page - 1) * limit)).
		Limit(int(limit)).
		Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return res, nil
}

//获得父级评论列表
func (c *Comment) GetParentCommentByWid(articleId uint) ([]*Comment, error) {
	var res []*Comment
	err := db.Model(c).Debug().
		Where(map[string]interface{}{
			"wid":       articleId,
			"parent_id": 0,
		}).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return res, nil
}

//获得子级评论列表
func (c *Comment) GetChildrenCommentByParentId(parentIds []uint) ([]*CommentItem, error) {
	var res []*CommentItem
	err := db.Model(c).Debug().
		Where(map[string]interface{}{
			"parent_id": parentIds,
		}).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return res, nil
}

//默认软删除
func (c *Comment) Delete(cid uint) error {
	err := db.Delete(c, cid).Error

	if err != nil {
		return err
	}
	return nil
}
