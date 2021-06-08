package request

import "gin-vue-admin/model"

type SysArticlesCagetorySearch struct {
	model.SysArticlesCagetory
	PageInfo
}
