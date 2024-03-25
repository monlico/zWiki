package routers

import (
	"zWiki/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//wiki的路由
	r = WikiRouter(r)

	return r
}
