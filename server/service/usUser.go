package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUsUser
//@description: 创建UsUser记录
//@param: usUser model.UsUser
//@return: err error

func CreateUsUser(usUser model.UsUser) (err error) {
	err = global.GVA_DB.Create(&usUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsUser
//@description: 删除UsUser记录
//@param: usUser model.UsUser
//@return: err error

func DeleteUsUser(usUser model.UsUser) (err error) {
	err = global.GVA_DB.Delete(&usUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsUserByIds
//@description: 批量删除UsUser记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUsUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.UsUser{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUsUser
//@description: 更新UsUser记录
//@param: usUser *model.UsUser
//@return: err error

func UpdateUsUser(usUser model.UsUser) (err error) {
	err = global.GVA_DB.Save(&usUser).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsUser
//@description: 根据id获取UsUser记录
//@param: id uint
//@return: err error, usUser model.UsUser

func GetUsUser(id uint) (err error, usUser model.UsUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&usUser).Error
	return
}

func GetUsUserByUsername(username string) (err error,usUser model.UsUser) {
	err = global.GVA_DB.Where("username = ?", username).First(&usUser).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsUserInfoList
//@description: 分页获取UsUser记录
//@param: info request.UsUserSearch
//@return: err error, list interface{}, total int64

func GetUsUserInfoList(info request.UsUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.UsUser{})
    var usUsers []model.UsUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Username != "" {
    	db = db.Where("username = ?",info.Username)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&usUsers).Error
	return err, usUsers, total
}