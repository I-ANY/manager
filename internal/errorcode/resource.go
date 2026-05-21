package errorcode

// ===== 资源/对象（300xxx）=====
var (
	ResourceNotFound      *Error
	ResourceConflict      *Error
	ResourceStatusInvalid *Error
)

func registerResource() {
	ResourceNotFound = NewError(5008, "资源不存在")
	ResourceConflict = NewError(5009, "资源冲突（重复/唯一键冲突）")
	ResourceStatusInvalid = NewError(5010, "资源当前状态不允许此操作")
}
