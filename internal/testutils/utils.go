package testutils

import (
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/types"
	// "sigs.k8s.io/controller-runtime/pkg/log/zap"

	// gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"

	stnrconfv1a1 "github.com/l7mp/stunner/pkg/apis/v1alpha1"

	"github.com/l7mp/stunner-gateway-operator/internal/config"
	// "github.com/l7mp/stunner-gateway-operator/internal/operator"
	// stnrv1a1 "github.com/l7mp/stunner-gateway-operator/api/v1alpha1"
)

func UnpackConfigMap(cm *corev1.ConfigMap) (stnrconfv1a1.StunnerConfig, error) {
	conf := stnrconfv1a1.StunnerConfig{}

	jsonConf, found := cm.Data[config.DefaultStunnerdConfigfileName]
	if !found {
		return conf, fmt.Errorf("error unpacking configmap data: %s not found",
			config.DefaultStunnerdConfigfileName)
	}

	if err := json.Unmarshal([]byte(jsonConf), &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
