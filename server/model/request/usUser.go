package request

import "gin-vue-admin/model"

type UsUserSearch struct{
    model.UsUser
    PageInfo
}