package wiki

import (
	"zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"

	"gorm.io/gorm"
)

type CommentService struct {
}

type CommentParams struct {
	Wid       uint   `json:"wid" validate:"required"`
	Text      string `json:"text" validate:"required"`
	Quote     string `json:"quote"` //引用
	ParentId  uint   `json:"parent_id"`
	RecoverId uint   `json:"recover_id" validate:"required"` //恢复的人id
}

// 添加评论
func (c *CommentService) AddComment(params CommentParams, uid uint) int {
	var commentModel wiki.Comment

	var (
		code int = e.SUCCESS
	)
	commentModel = wiki.Comment{
		Wid:        params.Wid,
		Text:       params.Text,
		Quote:      params.Quote,
		ParentId:   params.ParentId,
		RecoverUid: params.RecoverId,
		Uid:        uid,
	}
	_, err := commentModel.Create()

	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
	}

	return code
}

type CommentItem struct {
	gorm.Model
	ParentId   uint          `json:"parent_id"`
	Wid        uint          `json:"wid"`         //文章id
	RecoverUid uint          `json:"recover_uid"` //恢复的人id
	Endorse    uint          `json:"endorse"`
	Uid        uint          `json:"uid"`
	Text       string        `json:"text"`
	Name       string        `json:"name"`
	Quote      string        `json:"quote"`
	Parent     string        `json:"parent"`
	Children   []CommentItem `json:"children"`
}

// 评论列表
func (c *CommentService) CommentList(articleId uint) ([]CommentItem, int) {
	var (
		commentModel wiki.Comment
		userModel    wiki.User
		code         int = e.SUCCESS
	)

	commentList, err := commentModel.CommentList(articleId)
	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
		return nil, code
	}

	var (
		allUserId []uint
		userMap   = make(map[uint]string)
	)
	//统计所有uid
	for _, v := range commentList {
		allUserId = append(allUserId, v.Uid, v.RecoverUid)
	}
	allUserDetail, err := userModel.GetUsersByIds(allUserId)

	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
		return nil, code
	}

	for _, v := range allUserDetail {
		userMap[v.ID] = v.Username
	}

	var (
		foreachComment func(comments []*wiki.Comment) []CommentItem
	)

	//遍历评论，赋值字段
	foreachComment = func(comments []*wiki.Comment) []CommentItem {
		var (
			res     []CommentItem
			resItem CommentItem
		)
		for _, v := range comments {
			resItem = CommentItem{
				Model:      v.Model,
				ParentId:   v.ParentId,
				Wid:        v.Wid,
				RecoverUid: v.RecoverUid,
				Quote:      v.Quote,
				Endorse:    v.Endorse,
				Uid:        v.Uid,
				Text:       v.Text,
			}
			resItem.Name = userMap[v.Uid]
			resItem.Parent = userMap[v.RecoverUid]
			if len(v.Children) > 0 {
				resItem.Children = foreachComment(v.Children)
			}
			res = append(res, resItem)
		}
		return res
	}

	res := foreachComment(commentList)

	return res, code
}
