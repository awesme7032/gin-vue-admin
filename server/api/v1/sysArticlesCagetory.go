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

// @Tags SysArticlesCagetory
// @Summary 创建SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "创建SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticlesCagetory/createSysArticlesCagetory [post]
func CreateSysArticlesCagetory(c *gin.Context) {
	var sysArticlesCagetory model.SysArticlesCagetory
	_ = c.ShouldBindJSON(&sysArticlesCagetory)
	if err := service.CreateSysArticlesCagetory(sysArticlesCagetory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysArticlesCagetory
// @Summary 删除SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "删除SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysArticlesCagetory/deleteSysArticlesCagetory [delete]
func DeleteSysArticlesCagetory(c *gin.Context) {
	var sysArticlesCagetory model.SysArticlesCagetory
	_ = c.ShouldBindJSON(&sysArticlesCagetory)
	if err := service.DeleteSysArticlesCagetory(sysArticlesCagetory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysArticlesCagetory
// @Summary 批量删除SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysArticlesCagetory/deleteSysArticlesCagetoryByIds [delete]
func DeleteSysArticlesCagetoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSysArticlesCagetoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SysArticlesCagetory
// @Summary 更新SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "更新SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysArticlesCagetory/updateSysArticlesCagetory [put]
func UpdateSysArticlesCagetory(c *gin.Context) {
	var sysArticlesCagetory model.SysArticlesCagetory
	_ = c.ShouldBindJSON(&sysArticlesCagetory)
	if err := service.UpdateSysArticlesCagetory(sysArticlesCagetory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysArticlesCagetory
// @Summary 用id查询SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "用id查询SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysArticlesCagetory/findSysArticlesCagetory [get]
func FindSysArticlesCagetory(c *gin.Context) {
	var sysArticlesCagetory model.SysArticlesCagetory
	_ = c.ShouldBindQuery(&sysArticlesCagetory)
	if err, resysArticlesCagetory := service.GetSysArticlesCagetory(sysArticlesCagetory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysArticlesCagetory": resysArticlesCagetory}, c)
	}
}

// @Tags SysArticlesCagetory
// @Summary 分页获取SysArticlesCagetory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysArticlesCagetorySearch true "分页获取SysArticlesCagetory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticlesCagetory/getSysArticlesCagetoryList [get]
func GetSysArticlesCagetoryList(c *gin.Context) {
	var pageInfo request.SysArticlesCagetorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSysArticlesCagetoryInfoList(pageInfo); err != nil {
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
