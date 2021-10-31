package request

// User register structure
type Register struct {
	Username string `json:"username" binding:"required" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" required_err:"密码不能为空"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required" required_err:"邮箱不能为空"`
	Mobile   string `json:"mobile" binding:"required" required_err:"手机号不能为空"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"code"`      // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
	SmsCode   string `json:"smsCode"`   // 手机验证码
	Mobile    string `json:"mobile"`    // 手机号
}

// Modify password structure
type UserChangePasswordStruct struct {
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserRole struct {
	RoleId uint `json:"roleId"` // 角色ID
}

// Modify  user's auth structure
type SetUserRoles struct {
	ID      uint
	RoleIds []string `json:"roleIds"` // 角色ID
}

type UserFormRequest struct {
	Id       uint   `json:"id"`
	Username string `json:"username" binding:"required" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"checkPasswordLen" checkPasswordLen_err:"密码至少大于6位小于20位"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required" required_err:"邮箱不能为空"`
}

type AdminFormRequest struct {
	Id           uint64   `json:"id"`
	Username     string   `json:"username" binding:"required" required_err:"用户名不能为空"`
	Mobile       string   `json:"mobile" binding:"required" required_err:"手机号不能为空"`
	Password     string   `json:"password" binding:"checkPasswordLen" checkPasswordLen_err:"密码至少大于6位小于20位"`
	Nickname     string   `json:"nickname"`
	Status       int      `json:"status"`
	Avatar       string   `json:"avatar"`
	Roles        []uint64 `json:"roles" binding:"required" required_err:"角色不能为空"`
	Email        string   `json:"email" binding:"required" required_err:"邮箱不能为空"`
	DepartmentId uint64   `json:"departmentId"`
}
