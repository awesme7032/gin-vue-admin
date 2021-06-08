package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysArticles
//@description: 创建SysArticles记录
//@param: sysArticles model.SysArticles
//@return: err error

func CreateSysArticles(sysArticles model.SysArticles) (err error) {
	err = global.GVA_DB.Create(&sysArticles).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysArticles
//@description: 删除SysArticles记录
//@param: sysArticles model.SysArticles
//@return: err error

func DeleteSysArticles(sysArticles model.SysArticles) (err error) {
	err = global.GVA_DB.Delete(&sysArticles).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysArticlesByIds
//@description: 批量删除SysArticles记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSysArticlesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysArticles{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysArticles
//@description: 更新SysArticles记录
//@param: sysArticles *model.SysArticles
//@return: err error

func UpdateSysArticles(sysArticles model.SysArticles) (err error) {
	err = global.GVA_DB.Save(&sysArticles).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysArticles
//@description: 根据id获取SysArticles记录
//@param: id uint
//@return: err error, sysArticles model.SysArticles

func GetSysArticles(id uint) (err error, sysArticles model.SysArticles) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysArticles).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysArticlesInfoList
//@description: 分页获取SysArticles记录
//@param: info request.SysArticlesSearch
//@return: err error, list interface{}, total int64

func GetSysArticlesInfoList(info request.SysArticlesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysArticles{})
	var sysArticless []model.SysArticles
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Username != "" {
		db = db.Where("`username` LIKE ?", "%"+info.Username+"%")
	}
	// 手动新增文章摘要搜索
	if info.Summary != "" {
		db = db.Where("`summary` LIKE ?", "%"+info.Summary+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysArticless).Error
	return err, sysArticless, total
}
