package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysArticlesRouter(Router *gin.RouterGroup) {
	SysArticlesRouter := Router.Group("sysArticles").Use(middleware.OperationRecord())
	{
		SysArticlesRouter.POST("createSysArticles", v1.CreateSysArticles)             // 新建SysArticles
		SysArticlesRouter.DELETE("deleteSysArticles", v1.DeleteSysArticles)           // 删除SysArticles
		SysArticlesRouter.DELETE("deleteSysArticlesByIds", v1.DeleteSysArticlesByIds) // 批量删除SysArticles
		SysArticlesRouter.PUT("updateSysArticles", v1.UpdateSysArticles)              // 更新SysArticles
		SysArticlesRouter.GET("findSysArticles", v1.FindSysArticles)                  // 根据ID获取SysArticles
		SysArticlesRouter.GET("getSysArticlesList", v1.GetSysArticlesList)            // 获取SysArticles列表
	}
}
