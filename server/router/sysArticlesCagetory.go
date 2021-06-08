package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysArticlesCagetoryRouter(Router *gin.RouterGroup) {
	SysArticlesCagetoryRouter := Router.Group("sysArticlesCagetory").Use(middleware.OperationRecord())
	{
		SysArticlesCagetoryRouter.POST("createSysArticlesCagetory", v1.CreateSysArticlesCagetory)             // 新建SysArticlesCagetory
		SysArticlesCagetoryRouter.DELETE("deleteSysArticlesCagetory", v1.DeleteSysArticlesCagetory)           // 删除SysArticlesCagetory
		SysArticlesCagetoryRouter.DELETE("deleteSysArticlesCagetoryByIds", v1.DeleteSysArticlesCagetoryByIds) // 批量删除SysArticlesCagetory
		SysArticlesCagetoryRouter.PUT("updateSysArticlesCagetory", v1.UpdateSysArticlesCagetory)              // 更新SysArticlesCagetory
		SysArticlesCagetoryRouter.GET("findSysArticlesCagetory", v1.FindSysArticlesCagetory)                  // 根据ID获取SysArticlesCagetory
		SysArticlesCagetoryRouter.GET("getSysArticlesCagetoryList", v1.GetSysArticlesCagetoryList)            // 获取SysArticlesCagetory列表
	}
}
