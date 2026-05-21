package cluster

import (
	"context"
	"encoding/base64"
	"os"
	"strings"
	"time"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8soperation/internal/app/dao"
	"k8soperation/internal/errorcode"
	"k8soperation/pkg/app"
)

func GetRestConfig(ctx context.Context, a *app.App, clusterID uint32) (*rest.Config, error) {
	if clusterID == 0 {
		clusterID = a.AppSetting.DefaultClusterID
	}

	if cfg, ok := restConfigFromDB(ctx, a, clusterID); ok {
		return cfg, nil
	}
	if cfg, ok := restConfigFromGlobalPath(a); ok {
		return cfg, nil
	}
	if cfg, ok := restConfigFromInCluster(); ok {
		return cfg, nil
	}

	return nil, errorcode.ErrorClusterInitFailed
}

func restConfigFromDB(ctx context.Context, a *app.App, clusterID uint32) (*rest.Config, bool) {
	if a.DB == nil {
		return nil, false
	}

	kc, err := dao.NewDao(a.DB).K8sClusterGetByID(clusterID)
	if err != nil {
		return nil, false
	}

	raw := strings.TrimSpace(kc.KubeConfig)
	if raw == "" {
		return nil, false
	}

	if cfg, err := clientcmd.RESTConfigFromKubeConfig([]byte(raw)); err == nil {
		tuneRESTConfig(cfg)
		return cfg, true
	}
	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, false
	}
	cfg, err := clientcmd.RESTConfigFromKubeConfig(decoded)
	if err != nil {
		return nil, false
	}
	tuneRESTConfig(cfg)
	return cfg, true
}

func restConfigFromGlobalPath(a *app.App) (*rest.Config, bool) {
	p := strings.TrimSpace(a.AppSetting.GlobalKubeConfigPath)
	if p == "" {
		return nil, false
	}
	if _, err := os.Stat(p); err != nil {
		return nil, false
	}
	cfg, err := clientcmd.BuildConfigFromFlags("", p)
	if err != nil {
		return nil, false
	}
	tuneRESTConfig(cfg)
	return cfg, true
}

func restConfigFromInCluster() (*rest.Config, bool) {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, false
	}
	tuneRESTConfig(cfg)
	return cfg, true
}

func tuneRESTConfig(cfg *rest.Config) {
	cfg.UserAgent = "k8soperation/1.0"
	cfg.QPS = 50
	cfg.Burst = 100
	cfg.Timeout = 30 * time.Second
}
