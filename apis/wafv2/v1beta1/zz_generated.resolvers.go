// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *WebACLAssociation) ResolveReferences( // ResolveReferences of this WebACLAssociation.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "Stage", "StageList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ResourceArnRef,
			Selector:     mg.Spec.ForProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("wafv2.aws.upbound.io", "v1beta1", "WebACL", "WebACLList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WebACLArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.WebACLArnRef,
			Selector:     mg.Spec.ForProvider.WebACLArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WebACLArn")
	}
	mg.Spec.ForProvider.WebACLArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WebACLArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WebACLLoggingConfiguration.
func (mg *WebACLLoggingConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.LogDestinationConfigs),
			Extract:       resource.ExtractParamPath("arn", false),
			References:    mg.Spec.ForProvider.LogDestinationConfigsRefs,
			Selector:      mg.Spec.ForProvider.LogDestinationConfigsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LogDestinationConfigs")
	}
	mg.Spec.ForProvider.LogDestinationConfigs = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.LogDestinationConfigsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("wafv2.aws.upbound.io", "v1beta1", "WebACL", "WebACLList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ResourceArnRef,
			Selector:     mg.Spec.ForProvider.ResourceArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceArn")
	}
	mg.Spec.ForProvider.ResourceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.LogDestinationConfigs),
			Extract:       resource.ExtractParamPath("arn", false),
			References:    mg.Spec.InitProvider.LogDestinationConfigsRefs,
			Selector:      mg.Spec.InitProvider.LogDestinationConfigsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LogDestinationConfigs")
	}
	mg.Spec.InitProvider.LogDestinationConfigs = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.LogDestinationConfigsRefs = mrsp.ResolvedReferences

	return nil
}
