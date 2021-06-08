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

// @Tags UsUser
// @Summary 创建UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "创建UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usUser/createUsUser [post]
func CreateUsUser(c *gin.Context) {
	var usUser model.UsUser
	_ = c.ShouldBindJSON(&usUser)
	if err := service.CreateUsUser(usUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags UsUser
// @Summary 删除UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "删除UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /usUser/deleteUsUser [delete]
func DeleteUsUser(c *gin.Context) {
	var usUser model.UsUser
	_ = c.ShouldBindJSON(&usUser)
	if err := service.DeleteUsUser(usUser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags UsUser
// @Summary 批量删除UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /usUser/deleteUsUserByIds [delete]
func DeleteUsUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUsUserByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags UsUser
// @Summary 更新UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "更新UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /usUser/updateUsUser [put]
func UpdateUsUser(c *gin.Context) {
	var usUser model.UsUser
	_ = c.ShouldBindJSON(&usUser)
	if err := service.UpdateUsUser(usUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags UsUser
// @Summary 用id查询UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "用id查询UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /usUser/findUsUser [get]
func FindUsUser(c *gin.Context) {
	var usUser model.UsUser
	_ = c.ShouldBindQuery(&usUser)
	if err, reusUser := service.GetUsUser(usUser.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reusUser": reusUser}, c)
	}
}

// @Tags UsUser
// @Summary 用username查询UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "用id查询UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /usUser/findUsUser [get]
func FindUsUserByUsername(c *gin.Context) {
	var usUser model.UsUser
	_ = c.ShouldBindQuery(&usUser)
	if err, reusUser := service.GetUsUserByUsername(usUser.Username); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reusUser": reusUser}, c)
	}
}

// @Tags UsUser
// @Summary 分页获取UsUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UsUserSearch true "分页获取UsUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usUser/getUsUserList [get]
func GetUsUserList(c *gin.Context) {
	var pageInfo request.UsUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUsUserInfoList(pageInfo); err != nil {
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
