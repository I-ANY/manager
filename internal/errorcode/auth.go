package errorcode

// ===== 鉴权/账号（200xxx）=====
var (
	Unauthorized         *Error
	PermissionDenied     *Error
	UserFrozenOrDisabled *Error
)

func registerAuth() {
	Unauthorized = NewError(2001, "未认证或认证失效")
	PermissionDenied = NewError(2002, "无权限执行该操作")
	UserFrozenOrDisabled = NewError(2003, "用户状态异常")
}
