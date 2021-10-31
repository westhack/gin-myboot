package request

import (
    "gin-myboot/modules/{{.ModuleName}}/model"
    "gin-myboot/modules/common/model/request"
)

type {{.StructName}}Search struct{
    model.{{.StructName}}
    request.PageInfo
}