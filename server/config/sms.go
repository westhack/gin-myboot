package config

type Alidy struct {
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	SignName        string `mapstructure:"sign-name" json:"signName" yaml:"sign-name"`
	TemplateCode    string `mapstructure:"template-code" json:"templateCode" yaml:"template-code"`
	TemplateParam   string `mapstructure:"template-param" json:"templateParam" yaml:"template-param"`
}


type Sms struct {
	SmsType    string `mapstructure:"sms-type" json:"smsType" yaml:"sms-type"`           // 短信类型
	ExpireTime int64 `mapstructure:"expire-time" json:"expireTime" yaml:"expire-time"` // 过期时间
}
