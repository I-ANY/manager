package kube_cluster

import (
	"github.com/gin-gonic/gin"
	v1 "k8soperation/internal/app/controllers/api/v1/k8s_cluster"
)

type K8sClusterRouter struct{}

func NewK8sRouter() *K8sClusterRouter {
	return &K8sClusterRouter{}
}

func (r *K8sClusterRouter) Inject(router *gin.RouterGroup) {
	cluster := v1.NewK8sClusterController()

	router.POST("/cluster/create", cluster.Create)
	router.GET("/cluster/list", cluster.List)
	router.POST("/cluster/update", cluster.Update)
	router.POST("/cluster/delete", cluster.Delete)
	router.POST("/cluster/init", cluster.Init)
}
