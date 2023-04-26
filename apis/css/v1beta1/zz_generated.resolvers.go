/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/vpc/v1beta1"
	tools "github.com/FlexibleEngineCloud/provider-flexibleengine/pkg/tools"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NodeConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SecurityGroupID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SecurityGroupIDRef,
				Selector:     mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SecurityGroupIDSelector,
				To: reference.To{
					List:    &v1beta1.SecurityGroupList{},
					Managed: &v1beta1.SecurityGroup{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SecurityGroupID")
			}
			mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SecurityGroupIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NodeConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SubnetID),
				Extract:      tools.ExtractorParamPathfunc(true, "id"),
				Reference:    mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1beta1.VPCSubnetList{},
					Managed: &v1beta1.VPCSubnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SubnetID")
			}
			mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NodeConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].VPCID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].VPCIDRef,
				Selector:     mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].VPCIDSelector,
				To: reference.To{
					List:    &v1beta1.VPCList{},
					Managed: &v1beta1.VPC{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].VPCID")
			}
			mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NodeConfig[i3].NetworkInfo[i4].VPCIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}
