package returnMsg

import (
	"net/http"
	"zWiki/pkg/e"

	"github.com/gin-gonic/gin"
)

//200返回成功
func ReturnSuccessMsg(code int, val interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"val":  val,
	})
}

//400状态码
func ReturnFailedMsg(code int, val interface{}, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"val":  val,
	})
}
