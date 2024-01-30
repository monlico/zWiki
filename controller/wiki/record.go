package wiki

import (
	wikiModel "zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
	"zWiki/pkg/pvalidate"
	"zWiki/pkg/returnMsg"
	"zWiki/services/wiki"

	"github.com/gin-gonic/gin"
)

//wiki历史记录列表
// @route: /record/index
func RecordListController(c *gin.Context) {

	var (
		records []*wikiModel.Record
		params  wiki.RecordParams
		code    int = e.SUCCESS
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

	records, code = (&wiki.RecordService{}).RecordList(params)
	returnMsg.ReturnSuccessMsg(code, records, c)
	return
}
