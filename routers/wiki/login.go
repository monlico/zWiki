package wiki

import (
	"zWiki/controller/wiki"
	"zWiki/controller/wikiWebSocket"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.Engine) *gin.Engine {
	loginRouter := r.Group("/login")
	loginRouter.GET("/getGroup", wiki.LoginGetGroupController)
	loginRouter.POST("/index", wiki.LoginController)

	loginRouter.GET("/tmpIndex", wikiWebSocket.HandleLoginConnections)

	return r
}
