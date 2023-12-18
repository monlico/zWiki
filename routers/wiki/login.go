package wiki

import "github.com/gin-gonic/gin"

func LoginRouter(r *gin.Engine) *gin.Engine {
	loginRouter := r.Group("/login")
	loginRouter.GET("/getGroup")

	return r
}
