package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EnvVarKV struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// +kubebuilder:validation:MinLength=1
	Value string `json:"value"`
}

type AppConfigSpec struct {
	// +kubebuilder:validation:MinLength=1
	AppName string `json:"appName"`

	// +kubebuilder:validation:MinLength=1
	Image string `json:"image"`

	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// +optional
	Env []EnvVarKV `json:"env,omitempty"`

	// +optional
	EnableMetrics bool `json:"enableMetrics,omitempty"`

	// +kubebuilder:validation:Enum=RollingUpdate;Recreate
	// +optional
	Strategy string `json:"strategy,omitempty"`
}

type AppConfigStatus struct {
	// +optional
	Phase string `json:"phase,omitempty"`

	// +optional
	Message string `json:"message,omitempty"`

	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// +listType=map
	// +listMapKey=types
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type AppConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec AppConfigSpec `json:"spec"`

	// +optional
	Status AppConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type AppConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppConfig{}, &AppConfigList{})
}
