package request

import (
    "gin-myboot/modules/article/model"
    "gin-myboot/modules/common/model/request"
)

type SysArticleSearch struct{
    model.SysArticle
    request.PageInfo
}