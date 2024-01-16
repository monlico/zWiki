package wiki

import (
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
	"zWiki/pkg/pvalidate"
	"zWiki/pkg/returnMsg"
	"zWiki/pkg/util"
	"zWiki/services/wiki"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//评论列表
// @route: /comment/index
func CommentListController(c *gin.Context) {

	var (
		comments interface{}
		code     int = e.SUCCESS
	)

	articleStrId, articleExist := c.GetQuery("wid")

	if !articleExist {
		code = e.ERROR_VALIDATOR
		returnMsg.ReturnFailedMsg(code, comments, c)
		return
	}

	articleIntId, err := com.StrTo(articleStrId).Int()

	if err != nil {
		code = e.ERROR_DATA_TRANSFORM
		returnMsg.ReturnFailedMsg(code, comments, c)
		return
	}

	comments, code = (&wiki.CommentService{}).CommentList(uint(articleIntId))

	if code != e.SUCCESS {
		returnMsg.ReturnFailedMsg(code, comments, c)
		return
	}

	returnMsg.ReturnSuccessMsg(code, comments, c)
	return
}

//添加评论
// @route: /login/getGroup
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
