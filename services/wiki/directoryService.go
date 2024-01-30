package wiki

import (
	"zWiki/model/mysql/wiki"
	"zWiki/pkg/e"
	"zWiki/pkg/logging"
)

type DirectoryService struct {
}

//目录列表
func (directory *DirectoryService) GetDirectory(gid uint) ([]wiki.DirectoryItem, int) {
	var (
		directoryModel wiki.Directory
		articleModel   wiki.Articles
	)

	var (
		code            int = e.SUCCESS
		allDirectoryIds []uint
	)

	//拿到所欲的目录
	directories, err := directoryModel.GetAllDirectoryByGid(gid)

	if err != nil {
		code = e.ERROR_MYSQL
		return nil, code
	}

	for _, v := range directories {
		allDirectoryIds = append(allDirectoryIds, v.ID)
	}

	articlesDetail, err := articleModel.GetArticlesByDirectoryId(allDirectoryIds)

	if err != nil {
		code = e.ERROR_MYSQL
		return nil, code
	}

	for k, v := range directories {
		directories[k].Wikis = articlesDetail[v.ID]
	}

	return directories, code
}

type AddDirectoryParam struct {
	ParentId uint   `json:"parent_id" validate:"required"`
	Name     string `json:"name"  validate:"required"`
	Gid      uint   `json:"gid"  validate:"required"`
	Uid      uint   `json:"uid"`
}

//新增目录
func (directory *DirectoryService) AddDirectory(params AddDirectoryParam) (int, error) {
	var (
		directoryModel wiki.Directory
	)

	var (
		code = e.SUCCESS
	)

	directoryModel = wiki.Directory{
		ParentId: params.ParentId,
		Name:     params.Name,
		Gid:      params.Gid,
		Uid:      params.Uid,
	}

	err := directoryModel.AddDirectory()

	if err != nil {
		code = e.ERROR_MYSQL
		logging.Error(err)
		return code, err
	}

	return code, nil
}
