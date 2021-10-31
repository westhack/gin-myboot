package service

import (
	"gin-myboot/utils"
)

type EmailService struct {
}

// EmailTest
// @function: EmailTest
// @description: 发送邮件测试
// @return: err error
func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}
