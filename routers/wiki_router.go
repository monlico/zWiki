package routers

import (
	"zWiki/routers/wiki"

	"github.com/gin-gonic/gin"
)

func WikiRouter(r *gin.Engine) *gin.Engine {
	r = wiki.LoginRouter(r)

	return r
}
