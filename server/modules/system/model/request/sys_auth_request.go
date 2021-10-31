package request

type UserUpdateAccountRequest struct {
	Id       uint64 `json:"id"`
	Username string `json:"username" binding:"required" required_err:"用户名不能为空"`
	Nickname string `json:"nickname" binding:"required" required_err:"昵称不能为空"`
	Mobile   string `json:"mobile" binding:"required" required_err:"手机号不能为空"`
	Email    string `json:"email" binding:"required" required_err:"邮箱不能为空"`
	Avatar   string `json:"avatar"   binding:"required" required_err:"头像不能为空"`
}

type SmsCodeRequest struct {
	Mobile string `json:"mobile"`
}
