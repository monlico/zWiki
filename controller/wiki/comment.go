package wiki

import (
	wikiModel "zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
	"zWiki/pkg/pvalidate"
	"zWiki/pkg/returnMsg"
	"zWiki/pkg/setting"
	"zWiki/pkg/util"
	"zWiki/services/wiki"

	"github.com/gin-gonic/gin"
)

//评论列表
// @route: /comment/index
func CommentListController(c *gin.Context) {

	var (
		comments []*wikiModel.CommentItem
		params   wiki.CommentListParams
		code     int = e.SUCCESS
	)

	if err := c.ShouldBindQuery(&params); err != nil {
		code = e.INVALID_PARAMS
		returnMsg.ReturnFailedMsg(code, "body param error:"+err.Error(), c)
		return
	}

	err := pvalidate.WikiValidator.Struct(params)
	if err != nil {
		// 输出验证错误信息
		var errMsg string
		errMsg = pvalidate.Translate(err)
		logging.Warn(errMsg)
		code = e.ERROR_VALIDATOR
		returnMsg.ReturnSuccessMsg(code, errMsg, c)
		return
	}

	if params.Limit == 0 {
		params.Limit = uint(setting.PageSize)
	}
	if params.Page == 0 {
		params.Page = 1
	}
	comments, code = (&wiki.CommentService{}).CommentList(params)
	returnMsg.ReturnSuccessMsg(code, comments, c)
	return
}

//添加评论
// @route: /comment/index
func AddCommentController(c *gin.Context) {

	var params wiki.CommentParams

	var (
		code int
	)

	// 将自定义验证器设置为默认验证器
	if err := c.ShouldBindJSON(&params); err != nil {
		code = e.INVALID_PARAMS
		returnMsg.ReturnFailedMsg(code, "body param verify error:"+err.Error(), c)
		return
	}

	err := pvalidate.WikiValidator.Struct(params)
	if err != nil {
		// 输出验证错误信息
		var errMsg string
		errMsg = pvalidate.Translate(err)
		logging.Warn(errMsg)
		code = e.ERROR_VALIDATOR
		returnMsg.ReturnSuccessMsg(code, errMsg, c)
		return
	}

	//获取token
	var token string = c.GetHeader("token")
	tmpUserDetail, err := util.ParseToken(token)

	code = (&wiki.CommentService{}).AddComment(params, tmpUserDetail.Id)

	returnMsg.ReturnSuccessMsg(code, "", c)
	return
}

//删除评论
// @route: /comment/index
func DeleteCommentController(c *gin.Context) {

	var params wiki.CommentParams

	var (
		code int
	)

	// 将自定义验证器设置为默认验证器
	if err := c.ShouldBindJSON(&params); err != nil {
		code = e.INVALID_PARAMS
		returnMsg.ReturnFailedMsg(code, "body param verify error:"+err.Error(), c)
		return
	}

	err := pvalidate.WikiValidator.Struct(params)
	if err != nil {
		// 输出验证错误信息
		var errMsg string
		errMsg = pvalidate.Translate(err)
		logging.Warn(errMsg)
		code = e.ERROR_VALIDATOR
		returnMsg.ReturnSuccessMsg(code, errMsg, c)
		return
	}

	//获取token
	var token string = c.GetHeader("token")
	tmpUserDetail, err := util.ParseToken(token)

	code = (&wiki.CommentService{}).AddComment(params, tmpUserDetail.Id)

	returnMsg.ReturnSuccessMsg(code, "", c)
	return
}
