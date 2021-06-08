import service from '@/utils/request'

// @Tags ExaExampleTest
// @Summary 创建ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "创建ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaExampleTest/createExaExampleTest [post]
export const createExaExampleTest = (data) => {
     return service({
         url: "/exaExampleTest/createExaExampleTest",
         method: 'post',
         data
     })
 }


// @Tags ExaExampleTest
// @Summary 删除ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "删除ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaExampleTest/deleteExaExampleTest [delete]
 export const deleteExaExampleTest = (data) => {
     return service({
         url: "/exaExampleTest/deleteExaExampleTest",
         method: 'delete',
         data
     })
 }

// @Tags ExaExampleTest
// @Summary 删除ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaExampleTest/deleteExaExampleTest [delete]
 export const deleteExaExampleTestByIds = (data) => {
     return service({
         url: "/exaExampleTest/deleteExaExampleTestByIds",
         method: 'delete',
         data
     })
 }

// @Tags ExaExampleTest
// @Summary 更新ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "更新ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exaExampleTest/updateExaExampleTest [put]
 export const updateExaExampleTest = (data) => {
     return service({
         url: "/exaExampleTest/updateExaExampleTest",
         method: 'put',
         data
     })
 }


// @Tags ExaExampleTest
// @Summary 用id查询ExaExampleTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaExampleTest true "用id查询ExaExampleTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exaExampleTest/findExaExampleTest [get]
 export const findExaExampleTest = (params) => {
     return service({
         url: "/exaExampleTest/findExaExampleTest",
         method: 'get',
         params
     })
 }


// @Tags ExaExampleTest
// @Summary 分页获取ExaExampleTest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ExaExampleTest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaExampleTest/getExaExampleTestList [get]
 export const getExaExampleTestList = (params) => {
     return service({
         url: "/exaExampleTest/getExaExampleTestList",
         method: 'get',
         params
     })
 }