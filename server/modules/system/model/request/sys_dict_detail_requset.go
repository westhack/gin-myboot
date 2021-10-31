package request

import (
	"gin-myboot/modules/common/model/request"
	system "gin-myboot/modules/system/model"
)

type SysDictDetailSearch struct {
	system.SysDictDetail
	request.PageInfo
}
