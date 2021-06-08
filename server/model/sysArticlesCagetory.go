// 自动生成模板SysArticlesCagetory
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type SysArticlesCagetory struct {
	global.GVA_MODEL
	CategoryName string `json:"categoryName" form:"categoryName" gorm:"column:category_name;comment:文章分类名称;type:varchar(255);size:255;"`
	CategoryNo   int    `json:"categoryNo" form:"categoryNo" gorm:"column:category_no;comment:文章分类排序;type:int;size:10;"`
}

func (SysArticlesCagetory) TableName() string {
	return "sys_articles_cagetory"
}
