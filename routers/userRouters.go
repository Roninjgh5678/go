package routers

import (
	"fmt"
	"golangproject/controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("我是中间件start")
	c.Next()
	fmt.Println("end我是中间件")
	end := time.Now().UnixNano()
	fmt.Println("使用时间", end-start)

}
func UserRoutersInit(r *gin.Engine) {
	//传入指针并分组
	userRouter := r.Group("/user")
	{
		//用户的初始
		userRouter.GET("/", initMiddleware, func(ctx *gin.Context) {
			querNname := ctx.DefaultQuery("name", "ronin")
			fmt.Println("请求参数获取为", querNname)
			ctx.JSON(http.StatusOK, map[string]any{
				"titile": "初始页面",
			})
		})
		//用户的注册界面
		userRouter.GET("/Login", controller.UserController{}.Login)

		//用户的文件上传界面
		userRouter.GET("/uploadIndex", controller.UserController{}.UploadIndex)
		//用户文件上传之后的界面
		userRouter.GET("/uploadFinished", controller.UserController{}.UploadFinished)

		//获取所有用户信息的接口
		userRouter.GET("/showAllUserData", controller.UserController{}.FindAllUserData)

		//增加用户数据的方法
		userRouter.GET("/addUser", controller.UserController{}.AddUserData)
		//删除用户数据的方法
		userRouter.GET("/deleteUser", controller.UserController{}.DeleteUserData)
		//增加用户数据的方法
		userRouter.GET("/updateUser", controller.UserController{}.UpdateUserData)
	}

}
