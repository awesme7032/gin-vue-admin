package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitExaExampleTestRouter(Router *gin.RouterGroup) {
	ExaExampleTestRouter := Router.Group("exaExampleTest").Use(middleware.OperationRecord())
	{
		ExaExampleTestRouter.POST("createExaExampleTest", v1.CreateExaExampleTest)   // 新建ExaExampleTest
		ExaExampleTestRouter.DELETE("deleteExaExampleTest", v1.DeleteExaExampleTest) // 删除ExaExampleTest
		ExaExampleTestRouter.DELETE("deleteExaExampleTestByIds", v1.DeleteExaExampleTestByIds) // 批量删除ExaExampleTest
		ExaExampleTestRouter.PUT("updateExaExampleTest", v1.UpdateExaExampleTest)    // 更新ExaExampleTest
		ExaExampleTestRouter.GET("findExaExampleTest", v1.FindExaExampleTest)        // 根据ID获取ExaExampleTest
		ExaExampleTestRouter.GET("getExaExampleTestList", v1.GetExaExampleTestList)  // 获取ExaExampleTest列表
	}
}
