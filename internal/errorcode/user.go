package errorcode

var (
	ErrorUserCreateFail *Error
	ErrorUserDeleteFail *Error
	ErrorUserUpdateFail *Error
	ErrorUserListFail   *Error
)

func registerUser() {
	ErrorUserCreateFail = NewError(2011, "创建用户失败")
	ErrorUserDeleteFail = NewError(2012, "删除用户失败,用户不存在")
	ErrorUserUpdateFail = NewError(2013, "更新用户失败,用户不存在")
	ErrorUserListFail = NewError(2014, "列出用户失败")
}
