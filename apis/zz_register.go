/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/gaetanars/provider-flexibleengine/apis/dedicatedelb/v1beta1"
	v1beta1ecs "github.com/gaetanars/provider-flexibleengine/apis/ecs/v1beta1"
	v1beta1eip "github.com/gaetanars/provider-flexibleengine/apis/eip/v1beta1"
	v1beta1elb "github.com/gaetanars/provider-flexibleengine/apis/elb/v1beta1"
	v1beta1iam "github.com/gaetanars/provider-flexibleengine/apis/iam/v1beta1"
	v1beta1oss "github.com/gaetanars/provider-flexibleengine/apis/oss/v1beta1"
	v1alpha1 "github.com/gaetanars/provider-flexibleengine/apis/v1alpha1"
	v1beta1apis "github.com/gaetanars/provider-flexibleengine/apis/v1beta1"
	v1beta1vpc "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1"
	v1beta1vpcep "github.com/gaetanars/provider-flexibleengine/apis/vpcep/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1ecs.SchemeBuilder.AddToScheme,
		v1beta1eip.SchemeBuilder.AddToScheme,
		v1beta1elb.SchemeBuilder.AddToScheme,
		v1beta1iam.SchemeBuilder.AddToScheme,
		v1beta1oss.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
		v1beta1vpc.SchemeBuilder.AddToScheme,
		v1beta1vpcep.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
