package controller

import (
	"fmt"
	"golangproject/models"
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

// 用户的查询数据库
func (con UserController) FindAllUserData(c *gin.Context) {
	userList := []models.User{}
	models.Db.Find(&userList)
	c.JSON(http.StatusOK, map[string]any{
		"result": userList,
	})

}

// 用户的增加数据库信息
func (con UserController) AddUserData(c *gin.Context) {
	user := models.User{
		Username: "创始人名字 ",
		Age:      25,
		Email:    "2222@qq.com",
		AddTime:  models.Getunix(),
	}
	models.Db.Create(&user)
	fmt.Println(user)
	c.JSON(http.StatusOK, map[string]any{
		"msg": "增加数据成功",
	})

}

// 删除用户信息
func (con UserController) DeleteUserData(c *gin.Context) {
	user := models.User{}
	models.Db.Where("id=?", 3).Delete(&user)
	fmt.Println(user)
	c.JSON(http.StatusOK, map[string]any{
		"msg": "删除数据成功",
	})

}

// 更新用户信息
func (con UserController) UpdateUserData(c *gin.Context) {
	user := models.User{}
	models.Db.Where("username=?", "Ronin").Find(&user)
	user.Username = "Ronin"
	user.Age--
	user.AddTime = models.Getunix()
	models.Db.Save(&user)
	fmt.Println(user)
	c.JSON(http.StatusOK, map[string]any{
		"msg": "修改数据成功",
	})

}
