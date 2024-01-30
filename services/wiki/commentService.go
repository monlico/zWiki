package wiki

import (
	"zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
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

type CommentListParams struct {
	Wid   uint `json:"wid" form:"wid" validate:"required"`
	Limit uint `json:"limit" form:"limit" validate:"required"`
	Page  uint `json:"page" form:"page" validate:"required"`
}

// 评论列表
func (c *CommentService) CommentList(params CommentListParams) ([]*wiki.CommentItem, int) {
	var (
		commentModel wiki.Comment
		userModel    wiki.User
		code         int = e.SUCCESS
	)

	commentList, err := commentModel.CommentList(params.Wid, params.Page, params.Limit)
	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
		return nil, code
	}

	var (
		allUserId   []uint
		allParentId []uint //获取所有commentid，获取他们的子集
		userMap     = make(map[uint]string)
		commentMap  = make(map[uint][]*wiki.CommentItem)
	)
	//统计所有uid
	for _, v := range commentList {
		allUserId = append(allUserId, v.Uid, v.RecoverUid)
		allParentId = append(allParentId, v.ID)
	}
	allUserDetail, err := userModel.GetUsersByIds(allUserId)

	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
		return nil, code
	}

	allChildrenDetail, err := commentModel.GetChildrenCommentByParentId(allParentId)

	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
		return nil, code
	}

	for _, v := range allUserDetail {
		userMap[v.ID] = v.Username
	}

	for _, v := range allChildrenDetail {
		v.Parent = userMap[v.RecoverUid]
		v.Name = userMap[v.Uid]
		commentMap[v.ParentId] = append(commentMap[v.ParentId], v)
	}

	//遍历评论，赋值字段
	for k, v := range commentList {
		commentList[k].Parent = userMap[v.RecoverUid]
		commentList[k].Name = userMap[v.Uid]
		commentList[k].Children = commentMap[v.ID]
	}
	return commentList, code
}

//删除评论
func (c *CommentService) DeleteComment(cid, uid, commentUid uint) int {
	var (
		commentModel wiki.Comment
	)
	var (
		code int = e.SUCCESS
	)
	//删除评论，校验能否能删除
	verifyDelete := c.VerifyCanDelete(commentUid, uid)
	if !verifyDelete {
		code = e.ERROR_PERMISSION_DELETE
		return code
	}
	//软删除
	err := commentModel.Delete(cid)

	if err != nil {
		logging.Error(err)
		code = e.ERROR_MYSQL
		return code
	}
	return code
	//返回成功或失败
}

func (c *CommentService) VerifyCanDelete(commentUid, uid uint) bool {
	if commentUid != uid { //只能自己删除自己的评论
		return false
	}
	return true
}
