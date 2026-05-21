package errorcode

var (
	ErrorK8sDeploymentCreateFail   *Error
	ErrorK8sDeploymentDeleteFail   *Error
	ErrorK8sDeploymentListFail     *Error
	ErrorK8sDeploymentDetailFail   *Error
	ErrorK8sDeploymentUpdateFail   *Error
	ErrorK8sDeploymentRollbackFail *Error
	ErrorK8sDeploymentScaleFail    *Error
	ErrorK8sDeploymentRestartFail  *Error
	ErrorK8sDeploymentGetPodFail   *Error
)

func Register_k8s_Deployment() {
	ErrorK8sDeploymentCreateFail = NewError(5040, "创建K8s Deployment失败")
	ErrorK8sDeploymentDeleteFail = NewError(5041, "删除K8s Deployment失败")
	ErrorK8sDeploymentListFail = NewError(5042, "获取K8s Deployment列表失败")
	ErrorK8sDeploymentDetailFail = NewError(5043, "获取K8s Deployment详情失败")
	ErrorK8sDeploymentUpdateFail = NewError(5044, "更新K8s Deployment失败")
	ErrorK8sDeploymentRollbackFail = NewError(5045, "回滚K8s Deployment失败")
	ErrorK8sDeploymentScaleFail = NewError(5046, "扩缩容K8s Deployment失败")
	ErrorK8sDeploymentRestartFail = NewError(5047, "重启K8s Deployment失败")
	ErrorK8sDeploymentGetPodFail = NewError(5048, "获取K8s Deployment Pod列表失败")
}
