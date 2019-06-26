package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CrossVersionObjectReference contains enough information to let you identify the referred resource.
type CrossVersionObjectReference struct {
	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds"
	Kind string `json:"kind"`
	// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name"`
	// API version of the referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
}

// WatermarkedPodAutoscalerSpec defines the desired state of WatermarkedPodAutoscaler
// +k8s:openapi-gen=true
type WatermarkedPodAutoscalerSpec struct {

	// +kubebuilder:validation:Minimum=0.01
	// +kubebuilder:validation:Maximum=0.99
	Tolerance float64 `json:"tolerance,omitempty"`

	HighWaterMark float64 `json:"highwatermark,omitempty"`
	LowWaterMark float64 `json:"lowwatermark,omitempty"`

	Algorithm string `json:"algorithm,omitempty"`

	// part of HorizontalPodAutoscalerSpec, see comments in the k8s-1.10.8 repo: staging/src/k8s.io/api/autoscaling/v1/types.go
	// reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption
	// and will set the desired number of pods by using its Scale subresource.
	ScaleTargetRef CrossVersionObjectReference `json:"scaleTargetRef"`
	// specifications that will be used to calculate the desired replica count
	Metrics []autoscalingv2.MetricSpec `json:"metrics,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1000
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1000
	MaxReplicas int32 `json:"maxReplicas"`

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// WatermarkedPodAutoscalerStatus defines the observed state of WatermarkedPodAutoscaler
// +k8s:openapi-gen=true
type WatermarkedPodAutoscalerStatus struct {
	ObservedGeneration *int64                                           `json:"observedGeneration,omitempty"`
	LastScaleTime      *metav1.Time                                     `json:"lastScaleTime,omitempty"`
	CurrentReplicas    int32                                            `json:"currentReplicas"`
	DesiredReplicas    int32                                            `json:"desiredReplicas"`
	CurrentMetrics     []autoscalingv2.MetricStatus                     `json:"currentMetrics"`
	Conditions         []autoscalingv2.HorizontalPodAutoscalerCondition `json:"conditions"`

	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WatermarkedPodAutoscaler is the Schema for the watermarkedpodautoscalers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type WatermarkedPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WatermarkedPodAutoscalerSpec   `json:"spec,omitempty"`
	Status WatermarkedPodAutoscalerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WatermarkedPodAutoscalerList contains a list of WatermarkedPodAutoscaler
type WatermarkedPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WatermarkedPodAutoscaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WatermarkedPodAutoscaler{}, &WatermarkedPodAutoscalerList{})
}
