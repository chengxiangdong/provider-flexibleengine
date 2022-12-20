/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/iam/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Subscription.
func (mg *Subscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Owner),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OwnerRef,
		Selector:     mg.Spec.ForProvider.OwnerSelector,
		To: reference.To{
			List:    &v1beta1.ProjectList{},
			Managed: &v1beta1.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Owner")
	}
	mg.Spec.ForProvider.Owner = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OwnerRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicUrn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TopicUrnRef,
		Selector:     mg.Spec.ForProvider.TopicUrnSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicUrn")
	}
	mg.Spec.ForProvider.TopicUrn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicUrnRef = rsp.ResolvedReference

	return nil
}
