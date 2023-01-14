/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/eip/v1beta1"
	v1beta11 "github.com/FrangipaneTeam/provider-flexibleengine/apis/ims/v1beta1"
	v1beta12 "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1"
	common "github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this FloatingIpAssociate.
func (mg *FloatingIpAssociate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FloatingIP),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FloatingIPRef,
		Selector:     mg.Spec.ForProvider.FloatingIPSelector,
		To: reference.To{
			List:    &v1beta1.EIPList{},
			Managed: &v1beta1.EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FloatingIP")
	}
	mg.Spec.ForProvider.FloatingIP = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FloatingIPRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ImageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ImageIDRef,
		Selector:     mg.Spec.ForProvider.ImageIDSelector,
		To: reference.To{
			List:    &v1beta11.ImageList{},
			Managed: &v1beta11.Image{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ImageID")
	}
	mg.Spec.ForProvider.ImageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ImageName),
		Extract:      common.ImageNameExtractor(),
		Reference:    mg.Spec.ForProvider.ImageNameRef,
		Selector:     mg.Spec.ForProvider.ImageNameSelector,
		To: reference.To{
			List:    &v1beta11.ImageList{},
			Managed: &v1beta11.Image{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ImageName")
	}
	mg.Spec.ForProvider.ImageName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyPairRef,
		Selector:     mg.Spec.ForProvider.KeyPairSelector,
		To: reference.To{
			List:    &KeyPairList{},
			Managed: &KeyPair{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyPair")
	}
	mg.Spec.ForProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyPairRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Network); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network[i3].UUID),
			Extract:      common.IDExtractor(),
			Reference:    mg.Spec.ForProvider.Network[i3].UUIDRef,
			Selector:     mg.Spec.ForProvider.Network[i3].UUIDSelector,
			To: reference.To{
				List:    &v1beta12.VPCSubnetList{},
				Managed: &v1beta12.VPCSubnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Network[i3].UUID")
		}
		mg.Spec.ForProvider.Network[i3].UUID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Network[i3].UUIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SchedulerHints); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SchedulerHints[i3].Group),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SchedulerHints[i3].GroupRef,
			Selector:     mg.Spec.ForProvider.SchedulerHints[i3].GroupSelector,
			To: reference.To{
				List:    &ServerGroupList{},
				Managed: &ServerGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SchedulerHints[i3].Group")
		}
		mg.Spec.ForProvider.SchedulerHints[i3].Group = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SchedulerHints[i3].GroupRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this InterfaceAttach.
func (mg *InterfaceAttach) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      common.IDExtractor(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1beta12.VPCSubnetList{},
			Managed: &v1beta12.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &v1beta12.PortList{},
			Managed: &v1beta12.Port{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	return nil
}
