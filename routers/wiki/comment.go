package wiki

import (
	"zWiki/controller/wiki"
	"zWiki/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.Engine) *gin.Engine {
	loginRouter := r.Group("/comment")
	loginRouter.Use(middleware.CherryTokenValidMiddleware())
	loginRouter.GET("/index", wiki.CommentListController)
	loginRouter.POST("/index", wiki.AddCommentController)

	return r
}
