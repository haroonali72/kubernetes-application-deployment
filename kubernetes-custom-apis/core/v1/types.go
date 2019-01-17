package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type RuntimeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              interface{} `json:"spec"`
	Status            interface{} `json:"status,omitempty"`
}
type RuntimeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []RuntimeConfig `json:"items"`
}
