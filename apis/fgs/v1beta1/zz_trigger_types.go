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

type ApigObservation struct {
}

type ApigParameters struct {

	// Specifies the API name. Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	APIName *string `json:"apiName" tf:"api_name,omitempty"`

	// Specifies the API environment name.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	EnvName *string `json:"envName" tf:"env_name,omitempty"`

	// Specifies the ID of the APIG group to which the API belongs.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	GroupID *string `json:"groupId" tf:"group_id,omitempty"`

	// resource ID in UUID format.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the request protocol of the API. The valid value are
	// HTTP and HTTPS. Default to HTTPS. Changing this will create a new trigger resource.
	// +kubebuilder:validation:Optional
	RequestProtocol *string `json:"requestProtocol,omitempty" tf:"request_protocol,omitempty"`

	// Specifies the security authentication mode. The valid values
	// are NONE, APP and IAM, default to IAM. Changing this will create a new trigger resource.
	// +kubebuilder:validation:Optional
	SecurityAuthentication *string `json:"securityAuthentication,omitempty" tf:"security_authentication,omitempty"`

	// Specifies the timeout for request sending. The valid value is range form
	// 1 to 60,000, default to 5,000. Changing this will create a new trigger resource.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type DisObservation struct {
}

type DisParameters struct {

	// Specifies the maximum volume of data that can be obtained for a single
	// request, in Byte. Only the records with a size smaller than this value can be obtained.
	// The valid value is range from 1,024 to 4,194,304.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	MaxFetchBytes *float64 `json:"maxFetchBytes" tf:"max_fetch_bytes,omitempty"`

	// Specifies the interval at which data is pulled from the specified stream.
	// The valid value is range from 2 to 60,000.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	PullPeriod *float64 `json:"pullPeriod" tf:"pull_period,omitempty"`

	// Specifies the determines whether to pull data only after the data pulled
	// in the last period has been processed.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	SerialEnable *bool `json:"serialEnable" tf:"serial_enable,omitempty"`

	// Specifies the type of starting position for DIS queue.
	// The valid values are as follows:
	// +kubebuilder:validation:Required
	StartingPosition *string `json:"startingPosition" tf:"starting_position,omitempty"`

	// Specifies the name of the DIS stream resource.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	StreamName *string `json:"streamName" tf:"stream_name,omitempty"`
}

type KafkaObservation struct {
}

type KafkaParameters struct {

	// +kubebuilder:validation:Optional
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// resource ID in UUID format.
	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Required
	TopicIds []*string `json:"topicIds" tf:"topic_ids,omitempty"`
}

type LtsObservation struct {
}

type LtsParameters struct {

	// resource ID in UUID format.
	// +kubebuilder:validation:Required
	LogGroupID *string `json:"logGroupId" tf:"log_group_id,omitempty"`

	// resource ID in UUID format.
	// +kubebuilder:validation:Required
	LogTopicID *string `json:"logTopicId" tf:"log_topic_id,omitempty"`
}

type ObsObservation struct {
}

type ObsParameters struct {

	// Specifies the OBS bucket name.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// Specifies the event notification name.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	EventNotificationName *string `json:"eventNotificationName" tf:"event_notification_name,omitempty"`

	// Specifies the events that can trigger functions.
	// Changing this will create a new trigger resource.
	// The valid values are as follows:
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// Specifies the prefix to limit notifications to objects beginning with this keyword.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Specifies the suffix to limit notifications to objects ending with this keyword.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type SmnObservation struct {
}

type SmnParameters struct {

	// Specifies the Uniform Resource Name (URN) for SMN topic.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	TopicUrn *string `json:"topicUrn" tf:"topic_urn,omitempty"`
}

type TimerObservation struct {
}

type TimerParameters struct {

	// Specifies the event used by the timer to trigger the function.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Optional
	AdditionalInformation *string `json:"additionalInformation,omitempty" tf:"additional_information,omitempty"`

	// Specifies the trigger name, which can contains of 1 to 64 characters.
	// The name must start with a letter, only letters, digits, hyphens (-) and underscores (_) are allowed.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the time schedule.
	// For the rate type, schedule is composed of time and time unit.
	// The time unit supports minutes (m), hours (h) and days (d).
	// For the corn expression, please refer to the
	// User Guide.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`

	// Specifies the type of the time schedule.
	// The valid values are Rate and Cron.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	ScheduleType *string `json:"scheduleType" tf:"schedule_type,omitempty"`
}

type TriggerObservation struct {

	// resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TriggerParameters struct {

	// Specifies the configuration of the shared APIG trigger.
	// Changing this will create a new trigger resource.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	Apig []ApigParameters `json:"apig,omitempty" tf:"apig,omitempty"`

	// Specifies the configuration of the DIS trigger.
	// Changing this will create a new trigger resource.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	Dis []DisParameters `json:"dis,omitempty" tf:"dis,omitempty"`

	// Specifies the Uniform Resource Name (URN) of the function.
	// Changing this will create a new trigger resource.
	// +crossplane:generate:reference:type=Function
	// +kubebuilder:validation:Optional
	FunctionUrn *string `json:"functionUrn,omitempty" tf:"function_urn,omitempty"`

	// Reference to a Function to populate functionUrn.
	// +kubebuilder:validation:Optional
	FunctionUrnRef *v1.Reference `json:"functionUrnRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionUrn.
	// +kubebuilder:validation:Optional
	FunctionUrnSelector *v1.Selector `json:"functionUrnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Kafka []KafkaParameters `json:"kafka,omitempty" tf:"kafka,omitempty"`

	// +kubebuilder:validation:Optional
	Lts []LtsParameters `json:"lts,omitempty" tf:"lts,omitempty"`

	// Specifies the configuration of the OBS trigger.
	// Changing this will create a new trigger resource.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	Obs []ObsParameters `json:"obs,omitempty" tf:"obs,omitempty"`

	// Specifies the region in which to create the trigger resource.
	// If omitted, the provider-level region will be used.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the configuration of the SMN trigger.
	// Changing this will create a new trigger resource.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	Smn []SmnParameters `json:"smn,omitempty" tf:"smn,omitempty"`

	// Specifies whether trigger is enabled. The valid values are ACTIVE and DISABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the configuration of the timing trigger.
	// Changing this will create a new trigger resource.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	Timer []TimerParameters `json:"timer,omitempty" tf:"timer,omitempty"`

	// Specifies the type of the function.
	// The valid values currently only support TIMER, OBS, SMN, DIS, and APIG.
	// Changing this will create a new trigger resource.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// TriggerSpec defines the desired state of Trigger
type TriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerParameters `json:"forProvider"`
}

// TriggerStatus defines the observed state of Trigger.
type TriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trigger is the Schema for the Triggers API. ""page_title: "flexibleengine_fgs_trigger"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec   `json:"spec"`
	Status            TriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerList contains a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trigger `json:"items"`
}

// Repository type metadata.
var (
	Trigger_Kind             = "Trigger"
	Trigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Trigger_Kind}.String()
	Trigger_KindAPIVersion   = Trigger_Kind + "." + CRDGroupVersion.String()
	Trigger_GroupVersionKind = CRDGroupVersion.WithKind(Trigger_Kind)
)

func init() {
	SchemeBuilder.Register(&Trigger{}, &TriggerList{})
}
