package errorcode

var (
	ErrorK8sServiceCreateFail *Error
	ErrorK8sServiceDeleteFail *Error
	ErrorK8sServiceListFail   *Error
	ErrorK8sServiceDetailFail *Error
	ErrorK8sServiceUpdateFail *Error
	ErrorK8sServicePatchFail  *Error
	ErrorK8sServiceSyncFail   *Error
	ErrorK8sServiceSelectFail *Error
)

func Register_k8s_Service() {
	ErrorK8sServiceCreateFail = NewError(5031, "创建K8s Service失败")
	ErrorK8sServiceDeleteFail = NewError(5032, "删除K8s Service失败")
	ErrorK8sServiceListFail = NewError(5033, "获取K8s Service列表失败")
	ErrorK8sServiceDetailFail = NewError(5034, "获取K8s Service详情失败")
	ErrorK8sServiceUpdateFail = NewError(5035, "更新K8s Service失败")
	ErrorK8sServicePatchFail = NewError(5036, "Patch K8s Service失败")
	ErrorK8sServiceSyncFail = NewError(5037, "同步K8s Service失败")
	ErrorK8sServiceSelectFail = NewError(5038, "选择K8s Service目标失败")
}
