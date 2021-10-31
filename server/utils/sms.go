package utils

import (
	"context"
	"errors"
	"gin-myboot/global"
	"gin-myboot/utils/sms"
	"time"
)

func SendSmsCode (phoneNumbers string) (_err error, code string) {

	if phoneNumbers == "" {
		return errors.New("手机号不能为空"), ""
	}

	smsCode := CreateCaptcha()

	key := "sms::" + phoneNumbers
	ctx := context.Background()

	expireTime := global.Config.Sms.ExpireTime
	smsType := global.Config.Sms.SmsType
	if smsType == "alidy" {
		err, _ := sms.AlidySendSms(phoneNumbers, smsCode)
		if err != nil {
			// return err, ""
		}
	}

	t := time.Duration(expireTime) * time.Second

	_err = global.Redis.Set(ctx, key, smsCode, t).Err()

	return _err, smsCode
}

func CheckSmsCode(mobile string, smsCode string, del bool) (error, string) {
	ctx := context.Background()
	key := "sms::" + mobile
	val, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		return errors.New("短信验证码错误"), ""
	}

	if smsCode == val {
		if del {
			//global.Redis.Del(ctx, key)
		}

		return nil, val
	}

	return errors.New("短信验证码错误"), ""
}