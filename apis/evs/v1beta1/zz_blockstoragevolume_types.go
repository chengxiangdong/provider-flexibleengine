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

type AttachmentObservation struct {
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type AttachmentParameters struct {
}

type BlockStorageVolumeObservation struct {

	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance
	// sees it.
	Attachment []AttachmentObservation `json:"attachment,omitempty" tf:"attachment,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BlockStorageVolumeParameters struct {

	// The availability zone for the volume.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Specifies to delete all snapshots associated with the EVS disk.
	// +kubebuilder:validation:Optional
	Cascade *bool `json:"cascade,omitempty" tf:"cascade,omitempty"`

	// The consistency group to place the volume
	// in.
	// +kubebuilder:validation:Optional
	ConsistencyGroupID *string `json:"consistencyGroupId,omitempty" tf:"consistency_group_id,omitempty"`

	// Reference to a  to populate consistencyGroupId.
	// +kubebuilder:validation:Optional
	ConsistencyGroupIDRef *v1.Reference `json:"consistencyGroupIdRef,omitempty" tf:"-"`

	// Selector for a  to populate consistencyGroupId.
	// +kubebuilder:validation:Optional
	ConsistencyGroupIDSelector *v1.Selector `json:"consistencyGroupIdSelector,omitempty" tf:"-"`

	// A description of the volume. Changing this updates
	// the volume's description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/ims/v1beta1.Image
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Reference to a Image in ims to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDRef *v1.Reference `json:"imageIdRef,omitempty" tf:"-"`

	// Selector for a Image in ims to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDSelector *v1.Selector `json:"imageIdSelector,omitempty" tf:"-"`

	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	// The EVS encryption capability with KMS key can be set with the following parameters:
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Specifies whether the EVS disk is shareable.
	// +kubebuilder:validation:Optional
	Multiattach *bool `json:"multiattach,omitempty" tf:"multiattach,omitempty"`

	// A unique name for the volume. Changing this updates the
	// volume's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new volume.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the volume to create (in gigabytes).
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/csbs/v1beta1.Backup
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Reference to a Backup in csbs to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDRef *v1.Reference `json:"snapshotIdRef,omitempty" tf:"-"`

	// Selector for a Backup in csbs to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDSelector *v1.Selector `json:"snapshotIdSelector,omitempty" tf:"-"`

	// The volume ID to replicate with.
	// +kubebuilder:validation:Optional
	SourceReplica *string `json:"sourceReplica,omitempty" tf:"source_replica,omitempty"`

	// The volume ID from which to create the volume.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	SourceVolID *string `json:"sourceVolId,omitempty" tf:"source_vol_id,omitempty"`

	// The key/value pairs to associate with the volume.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of volume to create.
	// Changing this creates a new volume.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

// BlockStorageVolumeSpec defines the desired state of BlockStorageVolume
type BlockStorageVolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BlockStorageVolumeParameters `json:"forProvider"`
}

// BlockStorageVolumeStatus defines the observed state of BlockStorageVolume.
type BlockStorageVolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BlockStorageVolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BlockStorageVolume is the Schema for the BlockStorageVolumes API. ""page_title: "flexibleengine_blockstorage_volume_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type BlockStorageVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BlockStorageVolumeSpec   `json:"spec"`
	Status            BlockStorageVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlockStorageVolumeList contains a list of BlockStorageVolumes
type BlockStorageVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlockStorageVolume `json:"items"`
}

// Repository type metadata.
var (
	BlockStorageVolume_Kind             = "BlockStorageVolume"
	BlockStorageVolume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BlockStorageVolume_Kind}.String()
	BlockStorageVolume_KindAPIVersion   = BlockStorageVolume_Kind + "." + CRDGroupVersion.String()
	BlockStorageVolume_GroupVersionKind = CRDGroupVersion.WithKind(BlockStorageVolume_Kind)
)

func init() {
	SchemeBuilder.Register(&BlockStorageVolume{}, &BlockStorageVolumeList{})
}
