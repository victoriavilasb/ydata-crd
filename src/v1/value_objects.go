package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// YdataSpec defines all attributes in spec
type YdataSpec struct {
	Replicas int `json:"replicas"`
}

// Ydata defines the structure of the resource
type Ydata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec YdataSpec `json:"spec"`
}
