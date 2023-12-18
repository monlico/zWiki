package main

import (
	"zWiki/routers/wiki"

	"github.com/gin-gonic/gin"
)

func MainRouter(r *gin.Engine) *gin.Engine {
	r = wiki.WikiLoginRouter(r)
	return r
}
