import service from '@/utils/request'

// @Tags UsUser
// @Summary 创建UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "创建UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usUser/createUsUser [post]
export const createUsUser = (data) => {
     return service({
         url: "/usUser/createUsUser",
         method: 'post',
         data
     })
 }


// @Tags UsUser
// @Summary 删除UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "删除UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /usUser/deleteUsUser [delete]
 export const deleteUsUser = (data) => {
     return service({
         url: "/usUser/deleteUsUser",
         method: 'delete',
         data
     })
 }

// @Tags UsUser
// @Summary 删除UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /usUser/deleteUsUser [delete]
 export const deleteUsUserByIds = (data) => {
     return service({
         url: "/usUser/deleteUsUserByIds",
         method: 'delete',
         data
     })
 }

// @Tags UsUser
// @Summary 更新UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "更新UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /usUser/updateUsUser [put]
 export const updateUsUser = (data) => {
     return service({
         url: "/usUser/updateUsUser",
         method: 'put',
         data
     })
 }


// @Tags UsUser
// @Summary 用id查询UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "用id查询UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /usUser/findUsUser [get]
 export const findUsUser = (params) => {
     return service({
         url: "/usUser/findUsUser",
         method: 'get',
         params
     })
 }

// @Tags UsUser
// @Summary 用username查询UsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UsUser true "用id查询UsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /usUser/FindUsUserByUsername [get]
export const FindUsUserByUsername = (params) => {
    return service({
        url: "/usUser/FindUsUserByUsername",
        method: 'get',
        params
    })
}

// @Tags UsUser
// @Summary 分页获取UsUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取UsUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /usUser/getUsUserList [get]
 export const getUsUserList = (params) => {
     return service({
         url: "/usUser/getUsUserList",
         method: 'get',
         params
     })
 }