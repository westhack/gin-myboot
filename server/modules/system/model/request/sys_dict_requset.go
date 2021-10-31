package request

import (
	system "gin-myboot/modules/system/model"
)

type SysDictSearch struct {
	system.SysDict
}

type SaveDetailRequest struct {
	DictId      uint64 `json:"dictId"`
	DictDetails []struct {
		Label string `json:"label"`
		Value string `json:"value"`
		Color string `json:"color"`
	} `json:"dictDetails"`
}
