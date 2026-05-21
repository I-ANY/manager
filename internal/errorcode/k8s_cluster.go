package errorcode

// ===== 集群相关（200xxx）=====
var (
	ErrorClusterNotFound   *Error
	ErrorClusterUnhealthy  *Error
	ErrorClusterForbidden  *Error
	ErrorClusterInitFailed *Error
	ErrorClusterDeleteFail *Error // 删除失败
	ErrorClusterUpdateFail *Error // 更新失败
	ErrorClusterQueryFail  *Error // 查询失败（列表/单查）
)

func registerCluster() {
	ErrorClusterNotFound = NewError(5001, "集群名字不存在")
	ErrorClusterUnhealthy = NewError(5002, "集群不可用")
	ErrorClusterForbidden = NewError(5003, "没有访问该集群的权限")
	ErrorClusterInitFailed = NewError(5004, "K8s 集群初始化失败")

	ErrorClusterDeleteFail = NewError(5005, "删除集群失败")
	ErrorClusterUpdateFail = NewError(5006, "更新集群失败")
	ErrorClusterQueryFail = NewError(5007, "查询集群失败")
}
