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

type ProtectedInstanceObservation struct {

	// ID of the protected instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the target server.
	TargetServer *string `json:"targetServer,omitempty" tf:"target_server,omitempty"`
}

type ProtectedInstanceParameters struct {

	// Specifies the ID of a storage pool. Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Specifies whether to delete the EIP of the target server. The default value is false.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	DeleteTargetEIP *bool `json:"deleteTargetEip,omitempty" tf:"delete_target_eip,omitempty"`

	// Specifies whether to delete the target server. The default value is false.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	DeleteTargetServer *bool `json:"deleteTargetServer,omitempty" tf:"delete_target_server,omitempty"`

	// The description of a protected instance. Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the ID of the protection group where a protected instance is added.
	// Changing this creates a new instance.
	// +crossplane:generate:reference:type=ProtectionGroup
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a ProtectionGroup to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a ProtectionGroup to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The name of a protected instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the IP address of the primary NIC on the target server.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	PrimaryIPAddress *string `json:"primaryIpAddress,omitempty" tf:"primary_ip_address,omitempty"`

	// Specifies the subnet ID of the primary NIC on the target server.
	// Changing this creates a new instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +kubebuilder:validation:Optional
	PrimarySubnetID *string `json:"primarySubnetId,omitempty" tf:"primary_subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate primarySubnetId.
	// +kubebuilder:validation:Optional
	PrimarySubnetIDRef *v1.Reference `json:"primarySubnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate primarySubnetId.
	// +kubebuilder:validation:Optional
	PrimarySubnetIDSelector *v1.Selector `json:"primarySubnetIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the source server. Changing this creates a new instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1.Instance
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Instance in ecs to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// ProtectedInstanceSpec defines the desired state of ProtectedInstance
type ProtectedInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtectedInstanceParameters `json:"forProvider"`
}

// ProtectedInstanceStatus defines the observed state of ProtectedInstance.
type ProtectedInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtectedInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectedInstance is the Schema for the ProtectedInstances API. ""page_title: "flexibleengine_sdrs_protectedinstance_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type ProtectedInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtectedInstanceSpec   `json:"spec"`
	Status            ProtectedInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectedInstanceList contains a list of ProtectedInstances
type ProtectedInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtectedInstance `json:"items"`
}

// Repository type metadata.
var (
	ProtectedInstance_Kind             = "ProtectedInstance"
	ProtectedInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtectedInstance_Kind}.String()
	ProtectedInstance_KindAPIVersion   = ProtectedInstance_Kind + "." + CRDGroupVersion.String()
	ProtectedInstance_GroupVersionKind = CRDGroupVersion.WithKind(ProtectedInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtectedInstance{}, &ProtectedInstanceList{})
}
