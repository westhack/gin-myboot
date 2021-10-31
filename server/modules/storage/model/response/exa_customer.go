package response

import (
	storage "gin-myboot/modules/storage/model"
)

type ExaCustomerResponse struct {
	Customer storage.ExaCustomer `json:"customer"`
}
