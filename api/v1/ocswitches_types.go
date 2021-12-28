/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OcswitchesSpec defines the desired state of Ocswitches
type OcswitchesSpec struct {
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	Host       string        `json:"host"`
	Port       string        `json:"port"`
	Bgp        []BgpNeighbor `json:"bgpneighbors"`
	BgpAS      uint32        `json:"bgpas"`
	BgpReplace bool          `json:"bgpreplace"`
}

type BgpNeighbor struct {
	NeighborIP string `json:"neighbor"`
	RemoteAs   uint32    `json:"remoteas"`
}

// OcswitchesStatus defines the observed state of Ocswitches
type OcswitchesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Ocswitches is the Schema for the ocswitches API
type Ocswitches struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OcswitchesSpec   `json:"spec,omitempty"`
	Status OcswitchesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OcswitchesList contains a list of Ocswitches
type OcswitchesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ocswitches `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Ocswitches{}, &OcswitchesList{})
}
