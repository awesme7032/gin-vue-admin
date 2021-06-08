import service from '@/utils/request'

// @Tags SysArticles
// @Summary 创建SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "创建SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticles/createSysArticles [post]
export const createSysArticles = (data) => {
     return service({
         url: "/sysArticles/createSysArticles",
         method: 'post',
         data
     })
 }


// @Tags SysArticles
// @Summary 删除SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "删除SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysArticles/deleteSysArticles [delete]
 export const deleteSysArticles = (data) => {
     return service({
         url: "/sysArticles/deleteSysArticles",
         method: 'delete',
         data
     })
 }

// @Tags SysArticles
// @Summary 删除SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysArticles/deleteSysArticles [delete]
 export const deleteSysArticlesByIds = (data) => {
     return service({
         url: "/sysArticles/deleteSysArticlesByIds",
         method: 'delete',
         data
     })
 }

// @Tags SysArticles
// @Summary 更新SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "更新SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysArticles/updateSysArticles [put]
 export const updateSysArticles = (data) => {
     return service({
         url: "/sysArticles/updateSysArticles",
         method: 'put',
         data
     })
 }


// @Tags SysArticles
// @Summary 用id查询SysArticles
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticles true "用id查询SysArticles"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysArticles/findSysArticles [get]
 export const findSysArticles = (params) => {
     return service({
         url: "/sysArticles/findSysArticles",
         method: 'get',
         params
     })
 }


// @Tags SysArticles
// @Summary 分页获取SysArticles列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SysArticles列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysArticles/getSysArticlesList [get]
 export const getSysArticlesList = (params) => {
     return service({
         url: "/sysArticles/getSysArticlesList",
         method: 'get',
         params
     })
 }