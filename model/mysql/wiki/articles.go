package wiki

import "gorm.io/gorm"

type Articles struct {
	gorm.Model

	Title string `json:"title"`
	Uid   uint   `json:"uid"`
	DtId  uint   `json:"dt_id"` //目录id
}

type ArticlesPart1 struct {
	Id    uint   `json:"id"`
	DtId  uint   `json:"dt_id"`
	Title string `json:"title"`
}

//评论列表
func (a *Articles) GetArticlesByDirectoryId(directoryId []uint) (map[uint][]ArticlesPart1, error) {
	var (
		arts []ArticlesPart1
		res  = make(map[uint][]ArticlesPart1)
	)

	err := db.Model(a).Select("id", "title", "dt_id").Debug().Where(map[string]interface{}{
		"dt_id": directoryId,
	}).Scan(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, art := range arts {
		res[art.DtId] = append(res[art.DtId], art)
	}

	return res, nil
}
