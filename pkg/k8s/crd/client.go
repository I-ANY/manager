package crd

import (
	"context"

	"k8soperation/internal/app/dao"
	"k8soperation/pkg/app"
	"k8soperation/pkg/cluster"
)

func BuildAppConfig(ctx context.Context, a *app.App, clusterID uint32) (*dao.KubeAppConfig, error) {
	cfg, err := cluster.GetRestConfig(ctx, a, clusterID)
	if err != nil {
		return nil, err
	}

	cli, err := cluster.NewAppConfigClient(cfg)
	if err != nil {
		return nil, err
	}
	return dao.NewKubeAppConfig(cli), nil
}
