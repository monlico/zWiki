package wiki

import (
	"zWiki/controller/wiki"

	"github.com/gin-gonic/gin"
)

func RecordRouter(r *gin.Engine) *gin.Engine {
	loginRouter := r.Group("/record")
	loginRouter.GET("/index", wiki.RecordListController)

	return r
}
