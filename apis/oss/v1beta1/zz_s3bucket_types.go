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

type LifecycleRuleExpirationObservation struct {
}

type LifecycleRuleExpirationParameters struct {

	// Specifies the date after which you want the corresponding action to take effect.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type LifecycleRuleNoncurrentVersionExpirationObservation struct {
}

type LifecycleRuleNoncurrentVersionExpirationParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type S3BucketCorsRuleObservation struct {
}

type S3BucketCorsRuleParameters struct {

	// Specifies which headers are allowed.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type S3BucketLifecycleRuleObservation struct {
}

type S3BucketLifecycleRuleParameters struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	// +kubebuilder:validation:Optional
	Expiration []LifecycleRuleExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies when noncurrent object versions expire (documented below).
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration []LifecycleRuleNoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3BucketLoggingObservation struct {
}

type S3BucketLoggingParameters struct {

	// The name of the bucket that will receive the log objects.
	// +kubebuilder:validation:Required
	TargetBucket *string `json:"targetBucket" tf:"target_bucket,omitempty"`

	// To specify a key prefix for log objects.
	// +kubebuilder:validation:Optional
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type S3BucketObservation struct {

	// The bucket domain name. Will be of format bucketname.s3.amazonaws.com.
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	// Unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type S3BucketParameters struct {

	// The canned ACL to apply. Defaults to "private".
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ARN of the bucket. Will be of format arn:aws:s3:::bucketname.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/oss/v1beta1.S3Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Creates a unique bucket name beginning with the specified prefix. Conflicts with bucket.
	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// Reference to a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	// +kubebuilder:validation:Optional
	CorsRule []S3BucketCorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The Route 53 Hosted Zone ID for this bucket's region.
	// +kubebuilder:validation:Optional
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// A configuration of object lifecycle management (documented below).
	// +kubebuilder:validation:Optional
	LifecycleRule []S3BucketLifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// A settings of bucket logging (documented below).
	// +kubebuilder:validation:Optional
	Logging []S3BucketLoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// A valid bucket policy JSON document. In this case, please make sure you use the verbose/specific version of the policy.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// If specified, the region this bucket should reside in. Otherwise, the region used by the callee.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A state of versioning (documented below)
	// +kubebuilder:validation:Optional
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// A website object (documented below).
	// +kubebuilder:validation:Optional
	Website []S3BucketWebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`

	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	// +kubebuilder:validation:Optional
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	// +kubebuilder:validation:Optional
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type S3BucketWebsiteObservation struct {
}

type S3BucketWebsiteParameters struct {

	// An absolute path to the document to return in case of a 4XX error.
	// +kubebuilder:validation:Optional
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// Amazon S3 returns this index document when requests are made to the root domain or any of the subfolders.
	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (http:// or https://) to use when redirecting requests. The default is the protocol that is used in the original request.
	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// A json array containing routing rules
	// describing redirect behavior and when redirects are applied.
	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type VersioningObservation struct {
}

type VersioningParameters struct {

	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable MFA delete for either Change the versioning state of your bucket or Permanently delete an object version. Default is false.
	// +kubebuilder:validation:Optional
	MfaDelete *bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

// S3BucketSpec defines the desired state of S3Bucket
type S3BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3BucketParameters `json:"forProvider"`
}

// S3BucketStatus defines the observed state of S3Bucket.
type S3BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3Bucket is the Schema for the S3Buckets API. ""page_title: "flexibleengine_s3_bucket"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type S3Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketSpec   `json:"spec"`
	Status            S3BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketList contains a list of S3Buckets
type S3BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Bucket `json:"items"`
}

// Repository type metadata.
var (
	S3Bucket_Kind             = "S3Bucket"
	S3Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: S3Bucket_Kind}.String()
	S3Bucket_KindAPIVersion   = S3Bucket_Kind + "." + CRDGroupVersion.String()
	S3Bucket_GroupVersionKind = CRDGroupVersion.WithKind(S3Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&S3Bucket{}, &S3BucketList{})
}
