package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysArticlesCagetory
//@description: 创建SysArticlesCagetory记录
//@param: sysArticlesCagetory model.SysArticlesCagetory
//@return: err error

func CreateSysArticlesCagetory(sysArticlesCagetory model.SysArticlesCagetory) (err error) {
	err = global.GVA_DB.Create(&sysArticlesCagetory).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysArticlesCagetory
//@description: 删除SysArticlesCagetory记录
//@param: sysArticlesCagetory model.SysArticlesCagetory
//@return: err error

func DeleteSysArticlesCagetory(sysArticlesCagetory model.SysArticlesCagetory) (err error) {
	err = global.GVA_DB.Delete(&sysArticlesCagetory).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysArticlesCagetoryByIds
//@description: 批量删除SysArticlesCagetory记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSysArticlesCagetoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysArticlesCagetory{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysArticlesCagetory
//@description: 更新SysArticlesCagetory记录
//@param: sysArticlesCagetory *model.SysArticlesCagetory
//@return: err error

func UpdateSysArticlesCagetory(sysArticlesCagetory model.SysArticlesCagetory) (err error) {
	err = global.GVA_DB.Save(&sysArticlesCagetory).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysArticlesCagetory
//@description: 根据id获取SysArticlesCagetory记录
//@param: id uint
//@return: err error, sysArticlesCagetory model.SysArticlesCagetory

func GetSysArticlesCagetory(id uint) (err error, sysArticlesCagetory model.SysArticlesCagetory) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysArticlesCagetory).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysArticlesCagetoryInfoList
//@description: 分页获取SysArticlesCagetory记录
//@param: info request.SysArticlesCagetorySearch
//@return: err error, list interface{}, total int64

func GetSysArticlesCagetoryInfoList(info request.SysArticlesCagetorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysArticlesCagetory{})
	var sysArticlesCagetorys []model.SysArticlesCagetory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysArticlesCagetorys).Error
	return err, sysArticlesCagetorys, total
}
