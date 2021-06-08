// 自动生成模板SysArticles
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type SysArticles struct {
	global.GVA_MODEL
	Username     string `json:"username" form:"username" gorm:"column:username;comment:作者;type:varchar(191);size:191;"`
	CategoryId   int    `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:文章分类外键;type:bigint;size:19;"`
	CommentCount int    `json:"commentCount" form:"commentCount" gorm:"column:comment_count;comment:评论次数
评论次数;type:int;size:10;"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:文章内容;type:longtext;size:500;"`
	Status    int    `json:"status" form:"status" gorm:"column:status;comment:状态,正常为1;type:int;size:10;"`
	Summary   string `json:"summary" form:"summary" gorm:"column:summary;comment:文章摘要;type:varchar(255);size:255;"`
	ViewCount int    `json:"viewCount" form:"viewCount" gorm:"column:view_count;comment:阅读次数;type:int;size:10;"`
}

func (SysArticles) TableName() string {
	return "sys_articles"
}
