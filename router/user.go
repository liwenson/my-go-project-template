package router

import (
	"github.com/gin-gonic/gin"
	v1 "nav/api/v1"
)

func InitUserRouers(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("get_user_list", v1.GetUserList)
	}

}
