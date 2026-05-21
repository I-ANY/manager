package services

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8soperation/internal/app/requests"
	"k8soperation/pkg/k8s/configmap"
)

func (s *Services) KubeCreateConfigMap(ctx context.Context,
	req *requests.KubeConfigMapCreateRequest,
) (*corev1.ConfigMap, error) {
	return configmap.CreateConfigMap(s.App().K8sClient(), ctx, req)
}

// KubeConfigMapList 获取 ConfigMap 列表（支持名称过滤 + 分页）
func (s *Services) KubeConfigMapList(ctx context.Context, param *requests.KubeConfigMapListRequest) ([]corev1.ConfigMap, int, error) {
	return configmap.GetConfigMapList(s.App().K8sClient(), ctx, param.Name, param.Namespace, param.Page, param.Limit)
}

func (s *Services) KubeConfigMapDetail(ctx context.Context, param *requests.KubeConfigMapDetailRequest) (*corev1.ConfigMap, error) {
	return configmap.GetConfigMapDetail(s.App().K8sClient(), ctx, param.Name, param.Namespace)
}

func (s *Services) KubeConfigMapDelete(ctx context.Context, param *requests.KubeConfigMapDeleteRequest) error {
	return configmap.DeleteConfigMap(s.App().K8sClient(), ctx, param.Name, param.Namespace)
}
