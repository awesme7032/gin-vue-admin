import service from '@/utils/request'

// @Tags SysArticlesCagetory
// @Summary 创建SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "创建SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticlesCagetory/createSysArticlesCagetory [post]
export const createSysArticlesCagetory = (data) => {
     return service({
         url: "/sysArticlesCagetory/createSysArticlesCagetory",
         method: 'post',
         data
     })
 }


// @Tags SysArticlesCagetory
// @Summary 删除SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "删除SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysArticlesCagetory/deleteSysArticlesCagetory [delete]
 export const deleteSysArticlesCagetory = (data) => {
     return service({
         url: "/sysArticlesCagetory/deleteSysArticlesCagetory",
         method: 'delete',
         data
     })
 }

// @Tags SysArticlesCagetory
// @Summary 删除SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysArticlesCagetory/deleteSysArticlesCagetory [delete]
 export const deleteSysArticlesCagetoryByIds = (data) => {
     return service({
         url: "/sysArticlesCagetory/deleteSysArticlesCagetoryByIds",
         method: 'delete',
         data
     })
 }

// @Tags SysArticlesCagetory
// @Summary 更新SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "更新SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysArticlesCagetory/updateSysArticlesCagetory [put]
 export const updateSysArticlesCagetory = (data) => {
     return service({
         url: "/sysArticlesCagetory/updateSysArticlesCagetory",
         method: 'put',
         data
     })
 }


// @Tags SysArticlesCagetory
// @Summary 用id查询SysArticlesCagetory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticlesCagetory true "用id查询SysArticlesCagetory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysArticlesCagetory/findSysArticlesCagetory [get]
 export const findSysArticlesCagetory = (params) => {
     return service({
         url: "/sysArticlesCagetory/findSysArticlesCagetory",
         method: 'get',
         params
     })
 }


// @Tags SysArticlesCagetory
// @Summary 分页获取SysArticlesCagetory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysArticlesCagetory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticlesCagetory/getSysArticlesCagetoryList [get]
 export const getSysArticlesCagetoryList = (params) => {
     return service({
         url: "/sysArticlesCagetory/getSysArticlesCagetoryList",
         method: 'get',
         params
     })
 }