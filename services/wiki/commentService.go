package wiki

import (
	"zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
)

type CommentService struct {
}

type CommentParams struct {
	ArticleId uint   `json:"article_id"`
	Text      string `json:"text"`
	ParentId  uint   `json:"parent_id"`
	RecoverId uint   `json:"recover_id"` //恢复的人id
}

// 添加评论
func (c *CommentService) AddComment(params CommentParams, uid uint) int {
	var commentModel wiki.Comment

	var (
		code int = e.SUCCESS
	)
	commentModel = wiki.Comment{
		ArticleId: params.ArticleId,
		Text:      params.Text,
		ParentId:  params.ParentId,
		RecoverId: params.RecoverId,
		Uid:       uid,
	}
	_, err := commentModel.Create()

	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
	}

	return code
}

// 评论列表
func (c *CommentService) CommentList(articleId uint) (*wiki.Comment, int) {
	var (
		commentModel wiki.Comment
		code         int = e.SUCCESS
	)

	commentList, err := commentModel.CommentList(articleId)
	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
	}

	return commentList, code
}
