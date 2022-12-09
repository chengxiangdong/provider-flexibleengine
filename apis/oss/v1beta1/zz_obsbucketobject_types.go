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

type OBSBucketObjectObservation struct {

	// the key of the resource supplied above.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// the size of the object in bytes.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// A unique version ID value for the object, if bucket versioning is enabled.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type OBSBucketObjectParameters struct {

	// The ACL policy to apply. Defaults to private.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/oss/v1beta1.S3Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The literal content being uploaded to the bucket.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream.
	// All Valid MIME Types are valid for this input.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Whether enable server-side encryption of the object in SSE-KMS mode.
	// +kubebuilder:validation:Optional
	Encryption *bool `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Specifies the unique identifier of the object content. It can be used to trigger updates.
	// The only meaningful value is md5(file("path_to_file")).
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// the key of the resource supplied above.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The name of the object once it is in the bucket.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The path to the source file being uploaded to the bucket.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (Optioanl) Specifies the storage class of the object. Defaults to STANDARD.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

// OBSBucketObjectSpec defines the desired state of OBSBucketObject
type OBSBucketObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OBSBucketObjectParameters `json:"forProvider"`
}

// OBSBucketObjectStatus defines the observed state of OBSBucketObject.
type OBSBucketObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OBSBucketObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OBSBucketObject is the Schema for the OBSBucketObjects API. ""page_title: "flexibleengine_obs_bucket_object"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type OBSBucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OBSBucketObjectSpec   `json:"spec"`
	Status            OBSBucketObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OBSBucketObjectList contains a list of OBSBucketObjects
type OBSBucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OBSBucketObject `json:"items"`
}

// Repository type metadata.
var (
	OBSBucketObject_Kind             = "OBSBucketObject"
	OBSBucketObject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OBSBucketObject_Kind}.String()
	OBSBucketObject_KindAPIVersion   = OBSBucketObject_Kind + "." + CRDGroupVersion.String()
	OBSBucketObject_GroupVersionKind = CRDGroupVersion.WithKind(OBSBucketObject_Kind)
)

func init() {
	SchemeBuilder.Register(&OBSBucketObject{}, &OBSBucketObjectList{})
}
