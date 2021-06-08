package request

import "gin-vue-admin/model"

type ExaExampleTestSearch struct{
    model.ExaExampleTest
    PageInfo
}