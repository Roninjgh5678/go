package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// user界面的Controller
type UserController struct {
	BaseController
}

func (con UserController) Login(c *gin.Context) {
	querNname := c.DefaultQuery("name", "ronin")
	fmt.Println("请求参数获取为", querNname)
	c.HTML(http.StatusOK, "default.html", map[string]any{
		"titile": "后台首页",
	})
}

// 用户的上传文件界面
func (con UserController) UploadIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", map[string]any{
		"titile": "上传界面",
	})

}

// 用户的上传文件界面
func (con UserController) UploadFinished(c *gin.Context) {
	// file,err:=c.FormFile("")
	c.JSON(http.StatusOK, map[string]any{
		"titile": "上传完成界面",
	})

}
