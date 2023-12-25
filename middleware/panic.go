package middleware

import (
	"runtime/debug"
	"zWiki/pkg/logging"

	"github.com/gin-gonic/gin"
)

func CommonCatchPanicMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if m := recover(); m != nil {
				var (
					routerStr string
				)
				logging.Error("panicRouter:" + routerStr + "/errorDetail:" + string(debug.Stack()))
			}
		}()

		c.Next()
	}
}
