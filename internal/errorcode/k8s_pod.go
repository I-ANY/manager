package errorcode

// 说明：将变量声明为包级全局，实际赋值在 registerPod() 里完成，
// 便于控制 AllowOverride 等策略，也避免 import 时的初始化顺序问题。
var (
	ErrorK8sPodUpdateFail         *Error
	ErrorK8sPodDeleteFail         *Error
	ErrorK8sPodListFail           *Error
	ErrorK8sPodDetailFail         *Error
	ErrorK8sGetContainerName      *Error
	ErrorK8sGetContainerImage     *Error
	ErrorK8sGetInitContainerName  *Error
	ErrorK8sGetInitContainerImage *Error
	ErrorK8sGetContainerLog       *Error
)

// 内部注册函数（由 Register() 调用）
func register_k8s_Pod() {
	// 如果你有“是否允许覆盖”的开关，这里统一由 NewError/内部 register 方法处理
	ErrorK8sPodUpdateFail = NewError(5018, "更新K8s Pod失败")
	ErrorK8sPodDeleteFail = NewError(5019, "删除K8s Pod失败")
	ErrorK8sPodListFail = NewError(5020, "获取K8s Pod列表失败")
	ErrorK8sPodDetailFail = NewError(5021, "获取K8s Pod详情失败")
	ErrorK8sGetContainerName = NewError(5022, "获取K8s Pod容器名失败")
	ErrorK8sGetContainerImage = NewError(5023, "获取K8s Pod容器镜像失败")
	ErrorK8sGetInitContainerName = NewError(5024, "获取K8s Pod Init容器名失败")
	ErrorK8sGetInitContainerImage = NewError(5025, "获取K8s Pod Init容器镜像失败")
	ErrorK8sGetContainerLog = NewError(5026, "获取K8s Pod 容器日志失败")
}
