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

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {

	// origin requests.
	// Only CORS requests matching the allowed header are valid.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies the acceptable operation type of buckets and objects.
	// The methods include GET, PUT, POST, DELETE or HEAD.
	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Requests from this origin can access the bucket. Multiple matching rules are allowed.
	// One rule occupies one line, and allows one wildcard character (*) at most.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Specifies the exposed header in CORS responses, providing additional information for clients.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies the duration that your browser can cache CORS responses, expressed in seconds.
	// The default value is 100.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {

	// Specifies the number of days when objects that have been last updated are automatically deleted.
	// The expiration time must be greater than the transition times.
	// +kubebuilder:validation:Required
	Days *float64 `json:"days" tf:"days,omitempty"`
}

type LifecycleRuleObservation struct {
}

type LifecycleRuleParameters struct {

	// Specifies lifecycle rule status.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period when objects that have been last updated are automatically deleted. (documented below).
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies a period when noncurrent object versions are automatically deleted. (documented below).
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Specifies a period when noncurrent object versions are automatically transitioned to STANDARD_IA or GLACIER storage class (documented below).
	// +kubebuilder:validation:Optional
	NoncurrentVersionTransition []NoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// If omitted, all objects in the bucket will be managed by the lifecycle rule.
	// The prefix cannot start or end with a slash (/), cannot have consecutive slashes (/), and cannot contain the following special characters: :*?"<>|.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies a period when objects that have been last updated are automatically transitioned to STANDARD_IA or GLACIER storage class (documented below).
	// +kubebuilder:validation:Optional
	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// The name of the bucket that will receive the log objects.
	// The acl policy of the target bucket should be log-delivery-write.
	// +kubebuilder:validation:Required
	TargetBucket *string `json:"targetBucket" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	// +kubebuilder:validation:Optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {
}

type NoncurrentVersionExpirationParameters struct {

	// Specifies the number of days when objects that have been last updated are automatically deleted.
	// The expiration time must be greater than the transition times.
	// +kubebuilder:validation:Required
	Days *float64 `json:"days" tf:"days,omitempty"`
}

type NoncurrentVersionTransitionObservation struct {
}

type NoncurrentVersionTransitionParameters struct {

	// Specifies the number of days when objects that have been last updated are automatically deleted.
	// The expiration time must be greater than the transition times.
	// +kubebuilder:validation:Required
	Days *float64 `json:"days" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// "STANDARD", "STANDARD_IA" (Infrequent Access) and "GLACIER" (Archive). Defaults to STANDARD.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type OBSBucketObservation struct {

	// The bucket domain name. Will be of format bucketname.oss.region.prod-cloud-ocb.orange-business.com.
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	// The name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OBSBucketParameters struct {

	// Specifies the ACL policy for a bucket. The predefined common policies are as follows:
	// "private", "public-read", "public-read-write" and "log-delivery-write". Defaults to private.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Specifies the name of the bucket. Changing this parameter will create a new resource.
	// A bucket must be named according to the globally applied DNS naming regulations as follows:
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/oss/v1beta1.S3Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// Whether enable default server-side encryption of the bucket in SSE-KMS mode.
	// +kubebuilder:validation:Optional
	Encryption *bool `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket so that
	// the bucket can be destroyed without error. Default to false.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Specifies the ID of a kms key. If omitted, the default master key will be used.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// A configuration of object lifecycle management (documented below).
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// Whether enable the multi-AZ mode for the bucket. When the multi-AZ mode is enabled,
	// data in the bucket is duplicated and stored in multiple AZs.
	// Changing this creates a new bucket.
	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// If specified, the region this bucket should reside in. Otherwise, the region used by the provider.
	// Changing this creates a new bucket.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// "STANDARD", "STANDARD_IA" (Infrequent Access) and "GLACIER" (Archive). Defaults to STANDARD.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Whether enable versioning. Once you version-enable a bucket, it can never return to an unversioned state.
	// You can, however, suspend versioning on that bucket.
	// +kubebuilder:validation:Optional
	Versioning *bool `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`
}

type TransitionObservation struct {
}

type TransitionParameters struct {

	// Specifies the number of days when objects that have been last updated are automatically deleted.
	// The expiration time must be greater than the transition times.
	// +kubebuilder:validation:Required
	Days *float64 `json:"days" tf:"days,omitempty"`

	// Specifies the storage class of the bucket. OBS provides three storage classes:
	// "STANDARD", "STANDARD_IA" (Infrequent Access) and "GLACIER" (Archive). Defaults to STANDARD.
	// +kubebuilder:validation:Required
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type WebsiteObservation struct {
}

type WebsiteParameters struct {

	// Specifies the error page returned when an error occurs during static website access.
	// Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.
	// +kubebuilder:validation:Optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Specifies the default homepage of the static website, only HTML web pages are supported.
	// OBS only allows files such as index.html in the root directory of a bucket to function as the default homepage.
	// That is to say, do not set the default homepage with a multi-level directory structure (for example, /page/index.html).
	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting requests. The default is the protocol that is used in the original request.
	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A JSON or XML format containing routing rules describing redirect behavior and when redirects are applied.
	// Each rule contains a Condition and a Redirect as shown in the following table:
	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

// OBSBucketSpec defines the desired state of OBSBucket
type OBSBucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OBSBucketParameters `json:"forProvider"`
}

// OBSBucketStatus defines the observed state of OBSBucket.
type OBSBucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OBSBucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OBSBucket is the Schema for the OBSBuckets API. ""page_title: "flexibleengine_obs_bucket"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type OBSBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OBSBucketSpec   `json:"spec"`
	Status            OBSBucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OBSBucketList contains a list of OBSBuckets
type OBSBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OBSBucket `json:"items"`
}

// Repository type metadata.
var (
	OBSBucket_Kind             = "OBSBucket"
	OBSBucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OBSBucket_Kind}.String()
	OBSBucket_KindAPIVersion   = OBSBucket_Kind + "." + CRDGroupVersion.String()
	OBSBucket_GroupVersionKind = CRDGroupVersion.WithKind(OBSBucket_Kind)
)

func init() {
	SchemeBuilder.Register(&OBSBucket{}, &OBSBucketList{})
}
