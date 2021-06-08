package request

import "gin-vue-admin/model"

type SysArticlesSearch struct {
	model.SysArticles
	PageInfo
}
