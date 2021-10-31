package request

import (
	"gin-myboot/modules/common/model/request"
	system "gin-myboot/modules/system/model"
)

type SysLogSearch struct {
	system.SysLog
	request.PageInfo
}
