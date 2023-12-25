package request

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"account.required":  "账号号码不能为空",
		"password.required": "用户密码不能为空",
	}
}

type Login struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"account.required":  "账号号码不能为空",
		"password.required": "用户密码不能为空",
	}
}
