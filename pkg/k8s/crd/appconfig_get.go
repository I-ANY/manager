package crd

import (
	"context"
	"k8soperation/internal/app/requests"
	appv1alpha1 "k8soperation/pkg/apis/appconfig/v1alpha1"
	"k8soperation/pkg/app"
)

func GetAppConfig(ctx context.Context, a *app.App, req *requests.KubeAppConfigNameRequest) (*appv1alpha1.AppConfig, error) {
	d, err := BuildAppConfig(ctx, a, req.ClusterID)
	if err != nil {
		return nil, err
	}
	return d.Get(ctx, req.Namespace, req.AppName)
}
