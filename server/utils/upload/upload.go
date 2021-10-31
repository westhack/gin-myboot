package upload

import (
	"gin-myboot/global"
	"mime/multipart"
)

// OSS
// @author: [ccfish86](https://github.com/ccfish86)
// @author: [SliverHorn](https://github.com/SliverHorn)
// @interface_name: OSS
// @description: OSS接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss
// @author: [ccfish86](https://github.com/ccfish86)
// @author: [SliverHorn](https://github.com/SliverHorn)
// @function: NewOss
// @description: OSS接口
// @description: OSS的实例化方法
// @return: OSS
func NewOss() OSS {
	switch global.Config.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	default:
		return &Local{}
	}
}


// NewOssByType
// @function: NewOssByType
// @description: OSS接口
// @description: OSS的实例化方法
// @return: OSS
func NewOssByType(ossType string) OSS {
	switch ossType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	default:
		return &Local{}
	}
}
