package errorcode

// ===== 依赖/三方（600xxx）=====
var (
	DBError             *Error
	CacheError          *Error
	UpstreamTimeout     *Error
	UpstreamBadResponse *Error
	RPCInvokeError      *Error
)

func registerDependency() {
	DBError = NewError(6001, "数据库操作失败")
	CacheError = NewError(6002, "缓存操作失败")
	UpstreamTimeout = NewError(6003, "上游接口超时")
	UpstreamBadResponse = NewError(6004, "上游响应异常")
	RPCInvokeError = NewError(6005, "RPC 调用失败")
}
