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

type InstanceObservation struct {

	// Time when the APIG instance is created, in RFC-3339 format.
	// schema: Deprecated; Time when the dedicated instance is created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Time when the dedicated instance is created, in RFC-3339 format.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The egress (nat) public ip address.
	// The egress (NAT) public IP address.
	EgressAddress *string `json:"egressAddress,omitempty" tf:"egress_address,omitempty"`

	// ID of the APIG dedicated instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ingress eip address.
	// The ingress EIP address.
	IngressAddress *string `json:"ingressAddress,omitempty" tf:"ingress_address,omitempty"`

	// End time of the maintenance time window, 4-hour difference between the start time and end time.
	// End time of the maintenance time window, 4-hour difference between the start time and end time.
	MaintainEnd *string `json:"maintainEnd,omitempty" tf:"maintain_end,omitempty"`

	// Status of the APIG dedicated instance.
	// Status of the dedicated instance.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The supported features of the APIG dedicated instance.
	// The supported features of the dedicated instance.
	SupportedFeatures []*string `json:"supportedFeatures,omitempty" tf:"supported_features,omitempty"`

	// The ingress private ip address of vpc.
	// The ingress private IP address of the VPC.
	VPCIngressAddress *string `json:"vpcIngressAddress,omitempty" tf:"vpc_ingress_address,omitempty"`
}

type InstanceParameters struct {

	// schema: Required; The name list of availability zones for the dedicated instance.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// Specifies an array of available zone names for the APIG dedicated
	// instance. Changing this will create a new APIG dedicated instance resource.
	// schema: Deprecated; The name list of availability zones for the dedicated instance.
	// +kubebuilder:validation:Optional
	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// Specifies the egress bandwidth size of the APIG dedicated instance. The range of
	// valid value is from 1 to 2000.
	// The egress bandwidth size of the dedicated instance.
	// +kubebuilder:validation:Optional
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// Specifies the description about the APIG dedicated instance. The description
	// contain a maximum of 255 characters and the angle brackets (< and >) are not allowed.
	// The description of the dedicated instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the eip ID associated with the APIG dedicated instance.
	// The EIP ID associated with the dedicated instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/eip/v1beta1.EIP
	// +kubebuilder:validation:Optional
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// Reference to a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDRef *v1.Reference `json:"eipIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDSelector *v1.Selector `json:"eipIdSelector,omitempty" tf:"-"`

	// Specifies the edition of the APIG dedicated instance. The supported editions
	// are as follows: BASIC, PROFESSIONAL, ENTERPRISE, PLATINUM. Changing this will create a new APIG dedicated instance
	// resource.
	// The edition of the dedicated instance.
	// +kubebuilder:validation:Required
	Edition *string `json:"edition" tf:"edition,omitempty"`

	// ID of the APIG dedicated instance.
	// The enterprise project ID to which the dedicated instance belongs.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Whether public access with an IPv6 address is supported.
	// +kubebuilder:validation:Optional
	IPv6Enable *bool `json:"ipv6Enable,omitempty" tf:"ipv6_enable,omitempty"`

	// The type of loadbalancer provider used by the instance.
	// +kubebuilder:validation:Optional
	LoadbalancerProvider *string `json:"loadbalancerProvider,omitempty" tf:"loadbalancer_provider,omitempty"`

	// Specifies a start time of the maintenance time window in the format 'xx:00:00'.
	// The value of xx can be 02, 06, 10, 14, 18 or 22.
	// The start time of the maintenance time window.
	// +kubebuilder:validation:Optional
	MaintainBegin *string `json:"maintainBegin,omitempty" tf:"maintain_begin,omitempty"`

	// Specifies the name of the API dedicated instance. The API group name consists of 3 to 64
	// characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.
	// The name of the dedicated instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the region in which to create the APIG dedicated instance resource.
	// If omitted, the provider-level region will be used. Changing this will create a new APIG dedicated instance resource.
	// The region in which to create the dedicated instance resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies an ID of the security group to which the APIG dedicated instance
	// belongs to.
	// The ID of the security group to which the dedicated instance belongs to.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies an ID of the VPC subnet used to create the APIG dedicated
	// instance. Changing this will create a new APIG dedicated instance resource.
	// The ID of the VPC subnet used to create the dedicated instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Specifies an ID of the VPC used to create the APIG dedicated instance.
	// Changing this will create a new APIG dedicated instance resource.
	// The ID of the VPC used to create the dedicated instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. ""page_title: "flexibleengine_apig_instance"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
