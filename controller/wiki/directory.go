package wiki

import (
	wikiModel "zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
	"zWiki/pkg/pvalidate"
	"zWiki/pkg/returnMsg"
	"zWiki/pkg/util"
	"zWiki/services/wiki"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//目录
// @route: /directory/index
func DirectoryListController(c *gin.Context) {
	var (
		directory []wikiModel.DirectoryItem
		code      int = e.SUCCESS
	)

	articleStrId, articleExist := c.GetQuery("gid")

	if !articleExist {
		code = e.ERROR_VALIDATOR
		returnMsg.ReturnFailedMsg(code, directory, c)
		return
	}

	articleIntId, err := com.StrTo(articleStrId).Int()

	if err != nil {
		code = e.ERROR_DATA_TRANSFORM
		returnMsg.ReturnFailedMsg(code, directory, c)
		return
	}

	directory, code = (&wiki.DirectoryService{}).GetDirectory(uint(articleIntId))

	if code != e.SUCCESS {
		returnMsg.ReturnFailedMsg(code, directory, c)
		return
	}

	returnMsg.ReturnSuccessMsg(code, directory, c)
	return
}

//添加目录
// @route: /directory/index
func DirectoryAddController(c *gin.Context) {
	var (
		params wiki.AddDirectoryParam
		code   int = e.SUCCESS
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

	params.Uid = tmpUserDetail.Id

	code, err = (&wiki.DirectoryService{}).AddDirectory(params)

	if code != e.SUCCESS {
		returnMsg.ReturnFailedMsg(code, err, c)
		return
	}

	returnMsg.ReturnSuccessMsg(code, "success", c)
	return
}
