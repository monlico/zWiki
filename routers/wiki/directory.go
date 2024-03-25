package wiki

import (
	"zWiki/controller/wiki"

	"github.com/gin-gonic/gin"
)

func DirectoryRouter(r *gin.Engine) *gin.Engine {
	loginRouter := r.Group("/directory")
	loginRouter.GET("/index", wiki.DirectoryListController)
	loginRouter.POST("/index", wiki.DirectoryAddController)

	return r
}
