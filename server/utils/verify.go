package utils

var (
	IdVerify             = Rules{"ID": {NotEmpty()}}
	ApiVerify            = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify           = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify       = Rules{"Title": {NotEmpty()}}
	LoginVerify          = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "Nickname": {NotEmpty()}, "Password": {NotEmpty()}, "RoleId": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify       = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	RoleVerify           = Rules{"RoleId": {NotEmpty()}, "Title": {NotEmpty()}, "ParentId": {NotEmpty()}}
	RoleIdVerify         = Rules{"RoleId": {NotEmpty()}}
	OldRoleVerify        = Rules{"OldRoleId": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserRoleVerify    = Rules{"RoleId": {NotEmpty()}}
)
