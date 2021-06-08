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

// @Tags SysArticles
// @Summary 创建SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "创建SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticles/createSysArticles [post]
func CreateSysArticles(c *gin.Context) {
	var sysArticles model.SysArticles
	_ = c.ShouldBindJSON(&sysArticles)
	if err := service.CreateSysArticles(sysArticles); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysArticles
// @Summary 删除SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "删除SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysArticles/deleteSysArticles [delete]
func DeleteSysArticles(c *gin.Context) {
	var sysArticles model.SysArticles
	_ = c.ShouldBindJSON(&sysArticles)
	if err := service.DeleteSysArticles(sysArticles); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysArticles
// @Summary 批量删除SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysArticles/deleteSysArticlesByIds [delete]
func DeleteSysArticlesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSysArticlesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SysArticles
// @Summary 更新SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "更新SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysArticles/updateSysArticles [put]
func UpdateSysArticles(c *gin.Context) {
	var sysArticles model.SysArticles
	_ = c.ShouldBindJSON(&sysArticles)
	if err := service.UpdateSysArticles(sysArticles); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysArticles
// @Summary 用id查询SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "用id查询SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysArticles/findSysArticles [get]
func FindSysArticles(c *gin.Context) {
	var sysArticles model.SysArticles
	_ = c.ShouldBindQuery(&sysArticles)
	if err, resysArticles := service.GetSysArticles(sysArticles.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysArticles": resysArticles}, c)
	}
}

// @Tags SysArticles
// @Summary 分页获取SysArticles列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysArticlesSearch true "分页获取SysArticles列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticles/getSysArticlesList [get]
func GetSysArticlesList(c *gin.Context) {
	var pageInfo request.SysArticlesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSysArticlesInfoList(pageInfo); err != nil {
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
