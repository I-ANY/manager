package crd

import (
	"context"

	"k8soperation/internal/app/requests"
	appv1alpha1 "k8soperation/pkg/apis/appconfig/v1alpha1"
	"k8soperation/pkg/app"
)

func UpdateAppConfig(ctx context.Context, a *app.App, req *requests.KubeAppConfigUpdateRequest) (*appv1alpha1.AppConfig, error) {
	d, err := BuildAppConfig(ctx, a, req.ClusterID)
	if err != nil {
		return nil, err
	}
	app, err := d.Get(ctx, req.Namespace, req.AppName)
	if err != nil {
		return nil, err
	}

	if req.Image != "" {
		app.Spec.Image = req.Image
	}
	if req.Replicas != nil {
		app.Spec.Replicas = req.Replicas
	}

	if err := d.Update(ctx, app); err != nil {
		return nil, err
	}
	return app, nil
}
