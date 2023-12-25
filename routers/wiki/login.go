package wiki

import (
	"zWiki/controller/wiki"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.Engine) *gin.Engine {
	loginRouter := r.Group("/login")
	loginRouter.POST("/getGroup", wiki.LoginGetGroupController)

	return r
}
