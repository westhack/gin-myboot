package service

import (
	"gin-myboot/modules/common/model/request"
)

type Demo struct {
	BaseService
}


func (this *Demo) GetList(queryParams request.QueryParams) (err error, ret interface{}) {
	return nil, nil
}
