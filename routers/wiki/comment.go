package wiki

import (
	"zWiki/controller/wiki"

	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.Engine) *gin.Engine {
	loginRouter := r.Group("/comment")
	loginRouter.GET("/index", wiki.CommentListController)
	loginRouter.POST("/index", wiki.AddCommentController)

	return r
}
