package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUsUserRouter(Router *gin.RouterGroup) {
	UsUserRouter := Router.Group("usUser").Use(middleware.OperationRecord())
	{
		UsUserRouter.POST("createUsUser", v1.CreateUsUser)   // 新建UsUser
		UsUserRouter.DELETE("deleteUsUser", v1.DeleteUsUser) // 删除UsUser
		UsUserRouter.DELETE("deleteUsUserByIds", v1.DeleteUsUserByIds) // 批量删除UsUser
		UsUserRouter.PUT("updateUsUser", v1.UpdateUsUser)    // 更新UsUser
		UsUserRouter.GET("findUsUser", v1.FindUsUser)        // 根据ID获取UsUser
		UsUserRouter.GET("getUsUserList", v1.GetUsUserList)  // 获取UsUser列表
	}
}
