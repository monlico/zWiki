package wiki

import (
	"zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
)

type RecordService struct {
}

type RecordParams struct {
	Wid   uint `json:"wid" form:"wid" validate:"required"`
	Limit int  `json:"limit" form:"limit" validate:"required"`
	Page  int  `json:"page" form:"page" validate:"required"`
}

// 评论列表
func (r *RecordService) RecordList(params RecordParams) ([]*wiki.Record, int) {
	var (
		commentModel wiki.Record
		code         int = e.SUCCESS
	)

	recordList, err := commentModel.RecordList(params.Wid, params.Page, params.Limit)
	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
		return nil, code
	}

	return recordList, code
}
