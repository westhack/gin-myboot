package sms

import (
	"gin-myboot/global"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func AlidySendSms(phoneNumbers string, code string) (_err error, str string) {

	accessKeyId := global.Config.Alidy.AccessKeyId
	accessKeySecret := global.Config.Alidy.AccessKeySecret
	signName := global.Config.Alidy.SignName
	//templateParam := global.Config.Alidy.TemplateParam
	templateCode := global.Config.Alidy.TemplateCode

	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err, ""
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{}
	sendSmsRequest.SetPhoneNumbers(phoneNumbers)   // PhoneNumbers接收短信的手机号码
	sendSmsRequest.SetTemplateCode(templateCode)   // TemplateCode短信模板ID
	sendSmsRequest.SetSignName(signName)           // SignName短信签名名称
	sendSmsRequest.SetTemplateParam("{\"code\":\"" +code+ "\"}") // TemplateParam短信模板变量对应的实际值
	response, _err := client.SendSms(sendSmsRequest)
	if _err != nil {
		global.Error("=====>AlidySendSms", _err)
		return _err, ""
	}

	return _err, tea.StringValue(response.Body.RequestId)

}
