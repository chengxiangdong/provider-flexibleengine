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

type RuleCcProtectionObservation struct {

	// The rule ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleCcProtectionParameters struct {

	// Specifies the action when the number of requests reaches the upper limit. Valid Options are:
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Specifies the content of the returned page.
	// +kubebuilder:validation:Optional
	BlockPageContent *string `json:"blockPageContent,omitempty" tf:"block_page_content,omitempty"`

	// Specifies the type of the returned page.
	// The options are application/json, text/html, and text/xml.
	// +kubebuilder:validation:Optional
	BlockPageType *string `json:"blockPageType,omitempty" tf:"block_page_type,omitempty"`

	// Specifies the lock duration. The value ranges from 0 seconds to 2^32 seconds.
	// +kubebuilder:validation:Optional
	BlockTime *float64 `json:"blockTime,omitempty" tf:"block_time,omitempty"`

	// Specifies the category content. The format is as follows: "http://www.example.com/path".
	// This field is mandatory when mode is set to other.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Specifies the cookie name. This field is mandatory when mode is set to cookie.
	// +kubebuilder:validation:Optional
	Cookie *string `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// Specifies the number of requests allowed from a web visitor in a rate limiting period.
	// The value ranges from 0 to 2^32.
	// +kubebuilder:validation:Required
	LimitNum *float64 `json:"limitNum" tf:"limit_num,omitempty"`

	// Specifies the rate limiting period. The value ranges from 0 seconds to 2^32 seconds.
	// +kubebuilder:validation:Required
	LimitPeriod *float64 `json:"limitPeriod" tf:"limit_period,omitempty"`

	// Specifies the rate limit mode. Valid Options are:
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// Specifies the URL to which the rule applies. The path ending with * indicates
	// that the path is used as a prefix. For example, if the path to be protected is /admin/test.php or /adminabc,
	// set Path to /admin*.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Specifies the WAF policy ID. Changing this creates a new rule.
<<<<<<< HEAD
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/waf/v1beta1.Policy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a Policy in waf to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a Policy in waf to populate policyId.
=======
	// +crossplane:generate:reference:type=Policy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a Policy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a Policy to populate policyId.
>>>>>>> 54cb9bd (Chore(cce): Fix NameSpace)
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`
}

// RuleCcProtectionSpec defines the desired state of RuleCcProtection
type RuleCcProtectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleCcProtectionParameters `json:"forProvider"`
}

// RuleCcProtectionStatus defines the observed state of RuleCcProtection.
type RuleCcProtectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleCcProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleCcProtection is the Schema for the RuleCcProtections API. ""page_title: "flexibleengine_waf_rule_cc_protection"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type RuleCcProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleCcProtectionSpec   `json:"spec"`
	Status            RuleCcProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleCcProtectionList contains a list of RuleCcProtections
type RuleCcProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleCcProtection `json:"items"`
}

// Repository type metadata.
var (
	RuleCcProtection_Kind             = "RuleCcProtection"
	RuleCcProtection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleCcProtection_Kind}.String()
	RuleCcProtection_KindAPIVersion   = RuleCcProtection_Kind + "." + CRDGroupVersion.String()
	RuleCcProtection_GroupVersionKind = CRDGroupVersion.WithKind(RuleCcProtection_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleCcProtection{}, &RuleCcProtectionList{})
}
