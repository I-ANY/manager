package crd

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8soperation/internal/app/requests"
	appv1alpha1 "k8soperation/pkg/apis/appconfig/v1alpha1"
	"k8soperation/pkg/app"
)

func CreateAppConfig(ctx context.Context, a *app.App, req *requests.KubeAppConfigCreateRequest) (*appv1alpha1.AppConfig, error) {
	d, err := BuildAppConfig(ctx, a, req.ClusterID)
	if err != nil {
		return nil, err
	}

	app := &appv1alpha1.AppConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.AppName,
			Namespace: req.Namespace,
		},
		Spec: appv1alpha1.AppConfigSpec{
			AppName:  req.AppName,
			Image:    req.Image,
			Replicas: req.Replicas,
		},
	}

	if err := d.Create(ctx, app); err != nil {
		return nil, err
	}
	return app, nil
}
