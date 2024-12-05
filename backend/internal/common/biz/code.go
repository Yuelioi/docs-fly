package biz

const Ok = 200

var (
	DBError         = NewError(-300, "数据库错误")
	AlreadyRegister = NewError(-301, "用户已注册")
	AuthError       = NewError(-401, "验证失败")
)
