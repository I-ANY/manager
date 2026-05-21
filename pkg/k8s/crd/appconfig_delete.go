package crd

import (
	"context"
	"k8soperation/internal/app/requests"
	"k8soperation/pkg/app"
)

func DeleteAppConfig(ctx context.Context, a *app.App, req *requests.KubeAppConfigNameRequest) error {
	d, err := BuildAppConfig(ctx, a, req.ClusterID)
	if err != nil {
		return err
	}
	return d.Delete(ctx, req.Namespace, req.AppName)
}
