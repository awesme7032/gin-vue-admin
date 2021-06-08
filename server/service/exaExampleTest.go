package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaExampleTest
//@description: 创建ExaExampleTest记录
//@param: exaExampleTest model.ExaExampleTest
//@return: err error

func CreateExaExampleTest(exaExampleTest model.ExaExampleTest) (err error) {
	err = global.GVA_DB.Create(&exaExampleTest).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExaExampleTest
//@description: 删除ExaExampleTest记录
//@param: exaExampleTest model.ExaExampleTest
//@return: err error

func DeleteExaExampleTest(exaExampleTest model.ExaExampleTest) (err error) {
	err = global.GVA_DB.Delete(&exaExampleTest).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExaExampleTestByIds
//@description: 批量删除ExaExampleTest记录
//@param: ids request.IdsReq
//@return: err error

func DeleteExaExampleTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ExaExampleTest{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaExampleTest
//@description: 更新ExaExampleTest记录
//@param: exaExampleTest *model.ExaExampleTest
//@return: err error

func UpdateExaExampleTest(exaExampleTest model.ExaExampleTest) (err error) {
	err = global.GVA_DB.Save(&exaExampleTest).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaExampleTest
//@description: 根据id获取ExaExampleTest记录
//@param: id uint
//@return: err error, exaExampleTest model.ExaExampleTest

func GetExaExampleTest(id uint) (err error, exaExampleTest model.ExaExampleTest) {
	err = global.GVA_DB.Where("id = ?", id).First(&exaExampleTest).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaExampleTestInfoList
//@description: 分页获取ExaExampleTest记录
//@param: info request.ExaExampleTestSearch
//@return: err error, list interface{}, total int64

func GetExaExampleTestInfoList(info request.ExaExampleTestSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ExaExampleTest{})
    var exaExampleTests []model.ExaExampleTest
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Username != "" {
    	db = db.Where("username = ?",info.Username)
	}
	// 模糊搜索关键字
	if info.Keywords != "" {
		db = db.Where("keywords like ?","%"+ info.Keywords + "%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&exaExampleTests).Error
	return err, exaExampleTests, total
}