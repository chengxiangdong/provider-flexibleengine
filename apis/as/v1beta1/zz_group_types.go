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

type GroupObservation struct {

	// Indicates the number of current instances in the AS group.
	CurrentInstanceNumber *float64 `json:"currentInstanceNumber,omitempty" tf:"current_instance_number,omitempty"`

	// The UUID of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instances IDs of the AS group.
	// The instances id list in the as group.
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// Indicates the status of the AS group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type GroupParameters struct {

	// The availability zones in which to create
	// the instances in the autoscaling group.
	// +kubebuilder:validation:Optional
	AvailableZones []*string `json:"availableZones,omitempty" tf:"available_zones,omitempty"`

	// The cooling duration (in seconds). The value ranges
	// from 0 to 86400, and is 900 by default.
	// The cooling duration, in seconds.
	// +kubebuilder:validation:Optional
	CoolDownTime *float64 `json:"coolDownTime,omitempty" tf:"cool_down_time,omitempty"`

	// Whether to delete the instances in the AS group
	// when deleting the AS group. The options are yes and no.
	// Whether to delete instances when they are removed from the AS group.
	// +kubebuilder:validation:Optional
	DeleteInstances *string `json:"deleteInstances,omitempty" tf:"delete_instances,omitempty"`

	// Whether to delete the elastic IP address bound to the
	// instances of AS group when deleting the instances. The options are true and false.
	// +kubebuilder:validation:Optional
	DeletePublicip *bool `json:"deletePublicip,omitempty" tf:"delete_publicip,omitempty"`

	// The expected number of instances. The default
	// value is the minimum number of instances. The value ranges from the minimum number of
	// instances to the maximum number of instances.
	// +kubebuilder:validation:Optional
	DesireInstanceNumber *float64 `json:"desireInstanceNumber,omitempty" tf:"desire_instance_number,omitempty"`

	// Whether to forcibly delete the AS group, remove the ECS instances and release them.
	// The default value is false.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// The health check method for instances
	// in the AS group. The health check methods include ELB_AUDIT and NOVA_AUDIT.
	// If load balancing is configured, the default value of this parameter is ELB_AUDIT.
	// Otherwise, the default value is NOVA_AUDIT.
	// +kubebuilder:validation:Optional
	HealthPeriodicAuditMethod *string `json:"healthPeriodicAuditMethod,omitempty" tf:"health_periodic_audit_method,omitempty"`

	// The health check period for instances.
	// The period has four options: 5 minutes (default), 15 minutes, 60 minutes, and 180 minutes.
	// The health check period for instances, in minutes.
	// +kubebuilder:validation:Optional
	HealthPeriodicAuditTime *float64 `json:"healthPeriodicAuditTime,omitempty" tf:"health_periodic_audit_time,omitempty"`

	// The instance removal policy. The policy has
	// four options: OLD_CONFIG_OLD_INSTANCE (default), OLD_CONFIG_NEW_INSTANCE,
	// OLD_INSTANCE, and NEW_INSTANCE.
	// +kubebuilder:validation:Optional
	InstanceTerminatePolicy *string `json:"instanceTerminatePolicy,omitempty" tf:"instance_terminate_policy,omitempty"`

	// The ELB listener IDs. The system supports up to
	// six ELB listeners, the IDs of which are separated using a comma (,).
	// The system supports the binding of up to six ELB listeners, the IDs of which are separated using a comma.
	// +kubebuilder:validation:Optional
	LBListenerID *string `json:"lbListenerId,omitempty" tf:"lb_listener_id,omitempty"`

	// An array of one or more enhanced load balancer.
	// The system supports the binding of up to six load balancers. The field is
	// alternative to lb_listener_id.  The object structure is documented below.
	// +kubebuilder:validation:Optional
	LbaasListeners []LbaasListenersParameters `json:"lbaasListeners,omitempty" tf:"lbaas_listeners,omitempty"`

	// The maximum number of instances.
	// The default value is 0.
	// +kubebuilder:validation:Optional
	MaxInstanceNumber *float64 `json:"maxInstanceNumber,omitempty" tf:"max_instance_number,omitempty"`

	// The minimum number of instances.
	// The default value is 0.
	// +kubebuilder:validation:Optional
	MinInstanceNumber *float64 `json:"minInstanceNumber,omitempty" tf:"min_instance_number,omitempty"`

	// An array of one or more network IDs.
	// The system supports up to five networks. The networks object structure
	// is documented below.
	// +kubebuilder:validation:Required
	Networks []NetworksParameters `json:"networks" tf:"networks,omitempty"`

	// The notification mode. The system only supports EMAIL
	// mode which refers to notification by email.
	// +kubebuilder:validation:Optional
	Notifications []*string `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// The region in which to create the AS group. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new AS group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The configuration ID which defines
	// configurations of instances in the AS group.
	// +crossplane:generate:reference:type=Configuration
	// +kubebuilder:validation:Optional
	ScalingConfigurationID *string `json:"scalingConfigurationId,omitempty" tf:"scaling_configuration_id,omitempty"`

	// Reference to a Configuration to populate scalingConfigurationId.
	// +kubebuilder:validation:Optional
	ScalingConfigurationIDRef *v1.Reference `json:"scalingConfigurationIdRef,omitempty" tf:"-"`

	// Selector for a Configuration to populate scalingConfigurationId.
	// +kubebuilder:validation:Optional
	ScalingConfigurationIDSelector *v1.Selector `json:"scalingConfigurationIdSelector,omitempty" tf:"-"`

	// The name of the scaling group. The name can contain letters,
	// digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.
	// +kubebuilder:validation:Required
	ScalingGroupName *string `json:"scalingGroupName" tf:"scaling_group_name,omitempty"`

	// An array of one security group ID to
	// associate with the group. The security_groups object structure is
	// documented below.
	// +kubebuilder:validation:Required
	SecurityGroups []SecurityGroupsParameters `json:"securityGroups" tf:"security_groups,omitempty"`

	// The key/value pairs to associate with the scaling group.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID. Changing this creates a new group.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type LbaasListenersObservation struct {
}

type LbaasListenersParameters struct {

	// Specifies the backend ECS group ID.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/elb/v1beta1.Pool
	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// Reference to a Pool in elb to populate poolId.
	// +kubebuilder:validation:Optional
	PoolIDRef *v1.Reference `json:"poolIdRef,omitempty" tf:"-"`

	// Selector for a Pool in elb to populate poolId.
	// +kubebuilder:validation:Optional
	PoolIDSelector *v1.Selector `json:"poolIdSelector,omitempty" tf:"-"`

	// Specifies the backend protocol, which is the port on which
	// a backend ECS listens for traffic. The number of the port ranges from 1 to 65535.
	// +kubebuilder:validation:Required
	ProtocolPort *float64 `json:"protocolPort" tf:"protocol_port,omitempty"`

	// Specifies the weight, which determines the portion of requests a
	// backend ECS processes compared to other backend ECSs added to the same listener. The value
	// of this parameter ranges from 0 to 100. The default value is 1.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type NetworksObservation struct {
}

type NetworksParameters struct {

	// The network UUID.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(true, "id")
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type SecurityGroupsObservation struct {
}

type SecurityGroupsParameters struct {

	// The UUID of the security group.
	// +crossplane:generate:reference:type=github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. ""page_title: "flexibleengine_as_group_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
