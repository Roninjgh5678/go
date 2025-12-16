package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用于基础返回
type BaseController struct{}

func (con BaseController) success(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]any{

		"msg": "成功返回",
	})
}
func (con BaseController) failed(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]any{

		"msg": "请求失败",
	})
}
