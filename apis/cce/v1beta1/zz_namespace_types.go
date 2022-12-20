/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NamespaceObservation struct {

	// The server time when namespace was created.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// The namespace ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The current phase of the namespace.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NamespaceParameters struct {

	// Specifies an unstructured key value map for external parameters.
	// Changing this will create a new namespace resource.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Specifies the cluster ID to which the CCE namespace belongs.
	// Changing this will create a new namespace resource.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Specifies the map of string keys and values for labels.
	// Changing this will create a new namespace resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Specifies the unique name of the namespace.
	// This parameter can contain a maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-),
	// and must start and end with lowercase letters and digits. Changing this will create a new namespace resource.
	// Exactly one of name or prefix must be provided.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies a prefix used by the server to generate a unique name.
	// This parameter can contain a maximum of 63 characters, which may consist of lowercase letters, digits and
	// hyphens (-), and must start and end with lowercase letters and digits.
	// Changing this will create a new namespace resource. Exactly one of name or prefix must be provided.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the region in which to create the namespace resource.
	// If omitted, the provider-level region will be used. Changing this will create a new namespace resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// NamespaceSpec defines the desired state of Namespace
type NamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamespaceParameters `json:"forProvider"`
}

// NamespaceStatus defines the observed state of Namespace.
type NamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Namespace is the Schema for the Namespaces API. ""page_title: "flexibleengine_cce_namespace"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Namespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespaceSpec   `json:"spec"`
	Status            NamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceList contains a list of Namespaces
type NamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Namespace `json:"items"`
}

// Repository type metadata.
var (
	Namespace_Kind             = "Namespace"
	Namespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Namespace_Kind}.String()
	Namespace_KindAPIVersion   = Namespace_Kind + "." + CRDGroupVersion.String()
	Namespace_GroupVersionKind = CRDGroupVersion.WithKind(Namespace_Kind)
)

func init() {
	SchemeBuilder.Register(&Namespace{}, &NamespaceList{})
}
