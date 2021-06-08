package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags ExaExampleTest
// @Summary 创建ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "创建ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaExampleTest/createExaExampleTest [post]
func CreateExaExampleTest(c *gin.Context) {
	var exaExampleTest model.ExaExampleTest
	_ = c.ShouldBindJSON(&exaExampleTest)
	if err := service.CreateExaExampleTest(exaExampleTest); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ExaExampleTest
// @Summary 删除ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "删除ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaExampleTest/deleteExaExampleTest [delete]
func DeleteExaExampleTest(c *gin.Context) {
	var exaExampleTest model.ExaExampleTest
	_ = c.ShouldBindJSON(&exaExampleTest)
	if err := service.DeleteExaExampleTest(exaExampleTest); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ExaExampleTest
// @Summary 批量删除ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exaExampleTest/deleteExaExampleTestByIds [delete]
func DeleteExaExampleTestByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteExaExampleTestByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ExaExampleTest
// @Summary 更新ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "更新ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exaExampleTest/updateExaExampleTest [put]
func UpdateExaExampleTest(c *gin.Context) {
	var exaExampleTest model.ExaExampleTest
	_ = c.ShouldBindJSON(&exaExampleTest)
	global.GVA_LOG.Error("更新!", zap.Any("data", exaExampleTest))
	if err := service.UpdateExaExampleTest(exaExampleTest); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ExaExampleTest
// @Summary 用id查询ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "用id查询ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exaExampleTest/findExaExampleTest [get]
func FindExaExampleTest(c *gin.Context) {
	var exaExampleTest model.ExaExampleTest
	_ = c.ShouldBindQuery(&exaExampleTest)
	if err, reexaExampleTest := service.GetExaExampleTest(exaExampleTest.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexaExampleTest": reexaExampleTest}, c)
	}
}

// @Tags ExaExampleTest
// @Summary 分页获取ExaExampleTest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ExaExampleTestSearch true "分页获取ExaExampleTest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaExampleTest/getExaExampleTestList [get]
func GetExaExampleTestList(c *gin.Context) {
	var pageInfo request.ExaExampleTestSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExaExampleTestInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
