package routers

import (
	"zWiki/middleware"
	"zWiki/routers/wiki"

	"github.com/gin-gonic/gin"
)

func WikiRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CherryTokenValidMiddleware())
	r = wiki.LoginRouter(r)     //登录
	r = wiki.CommentRouter(r)   //评论
	r = wiki.DirectoryRouter(r) //目录
	r = wiki.RecordRouter(r)    //历史记录

	return r
}
