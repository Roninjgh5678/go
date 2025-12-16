package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//传入指针并分组
	adminRouter := r.Group("/admin")
	{
		//登录的初始
		adminRouter.GET("/", func(ctx *gin.Context) {
			querNname := ctx.DefaultQuery("name", "ronin")
			fmt.Println("请求参数获取为", querNname)
			ctx.JSON(http.StatusOK, map[string]any{
				"titile": "管理员页面",
			})
		})
		//登录的注册界面
		adminRouter.GET("/Login", func(ctx *gin.Context) {
			querNname := ctx.DefaultQuery("name", "ronin")
			fmt.Println("请求参数获取为", querNname)
			ctx.HTML(http.StatusOK, "default.html", map[string]any{
				"titile": "后台首页",
			})
		})
		//登录的注册界面
		adminRouter.GET("/Register", func(ctx *gin.Context) {
			querNname := ctx.DefaultQuery("name", "ronin")
			fmt.Println("请求参数获取为", querNname)
			ctx.HTML(http.StatusOK, "Register.html", map[string]any{
				"titile": "后台首页",
			})
		})
	}

}
