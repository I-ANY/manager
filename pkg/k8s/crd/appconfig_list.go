package crd

import (
	"context"
	"k8soperation/internal/app/requests"
	appv1alpha1 "k8soperation/pkg/apis/appconfig/v1alpha1"
	"k8soperation/pkg/app"
	"k8soperation/pkg/cluster"
)

func List(ctx context.Context, a *app.App, req *requests.KubeAppConfigListRequest) ([]appv1alpha1.AppConfig, error) {
	cfg, err := cluster.GetRestConfig(ctx, a, req.ClusterID)
	if err != nil {
		return nil, err
	}

	cli, err := cluster.NewAppConfigClient(cfg)
	if err != nil {
		return nil, err
	}

	var list appv1alpha1.AppConfigList
	if err := cli.List(ctx, &list); err != nil {
		return nil, err
	}
	return list.Items, nil
}
