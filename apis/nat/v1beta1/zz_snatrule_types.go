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

type SnatRuleObservation struct {

	// The actual floating IP address.
	FloatingIPAddress *string `json:"floatingIpAddress,omitempty" tf:"floating_ip_address,omitempty"`

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The status of the snat rule.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SnatRuleParameters struct {

	// Specifies CIDR, which can be in the format of a network segment or a host IP address.
	// This parameter and subnet_id are alternative. Changing this creates a new snat rule.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// ID of the floating ip this snat rule connets to.
	// Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/eip/v1beta1.EIP
	// +kubebuilder:validation:Optional
	FloatingIPID *string `json:"floatingIpId,omitempty" tf:"floating_ip_id,omitempty"`

	// Reference to a EIP in eip to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDRef *v1.Reference `json:"floatingIpIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate floatingIpId.
	// +kubebuilder:validation:Optional
	FloatingIPIDSelector *v1.Selector `json:"floatingIpIdSelector,omitempty" tf:"-"`

	// ID of the nat gateway this snat rule belongs to.
	// Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=Gateway
	// +kubebuilder:validation:Optional
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Reference to a Gateway to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDRef *v1.Reference `json:"natGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway to populate natGatewayId.
	// +kubebuilder:validation:Optional
	NATGatewayIDSelector *v1.Selector `json:"natGatewayIdSelector,omitempty" tf:"-"`

	// The resource ID in UUID format.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(true, "id")
	// +crossplane:generate:reference:refFieldName=NetworkIDRef
	// +crossplane:generate:reference:selectorFieldName=NetworkIDSelector
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The region in which to obtain the V2 nat client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new snat rule.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the scenario. The valid value is 0 (VPC scenario) and 1 (Direct Connect scenario).
	// Only cidr can be specified over a Direct Connect connection.
	// If no value is entered, the default value 0 (VPC scenario) is used.
	// Changing this creates a new snat rule.
	// +kubebuilder:validation:Optional
	SourceType *float64 `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// ID of the VPC Subnet this snat rule connects to.
	// This parameter and cidr are alternative. Changing this creates a new snat rule.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(true, "id")
	// +crossplane:generate:reference:refFieldName=SubnetIDRef
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// SnatRuleSpec defines the desired state of SnatRule
type SnatRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnatRuleParameters `json:"forProvider"`
}

// SnatRuleStatus defines the observed state of SnatRule.
type SnatRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnatRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnatRule is the Schema for the SnatRules API. ""page_title: "flexibleengine_nat_snat_rule_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type SnatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnatRuleSpec   `json:"spec"`
	Status            SnatRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnatRuleList contains a list of SnatRules
type SnatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnatRule `json:"items"`
}

// Repository type metadata.
var (
	SnatRule_Kind             = "SnatRule"
	SnatRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnatRule_Kind}.String()
	SnatRule_KindAPIVersion   = SnatRule_Kind + "." + CRDGroupVersion.String()
	SnatRule_GroupVersionKind = CRDGroupVersion.WithKind(SnatRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SnatRule{}, &SnatRuleList{})
}
