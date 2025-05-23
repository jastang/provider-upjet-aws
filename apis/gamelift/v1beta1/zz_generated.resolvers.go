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
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Build.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Build) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Bucket),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.StorageLocation[i3].BucketRef,
				Selector:     mg.Spec.ForProvider.StorageLocation[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Bucket")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Key),
				Extract:      resource.ExtractParamPath("key", false),
				Reference:    mg.Spec.ForProvider.StorageLocation[i3].KeyRef,
				Selector:     mg.Spec.ForProvider.StorageLocation[i3].KeySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Key")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].KeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].RoleArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef,
				Selector:     mg.Spec.ForProvider.StorageLocation[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].RoleArn")
		}
		mg.Spec.ForProvider.StorageLocation[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageLocation[i3].Bucket),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.StorageLocation[i3].BucketRef,
				Selector:     mg.Spec.InitProvider.StorageLocation[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageLocation[i3].Bucket")
		}
		mg.Spec.InitProvider.StorageLocation[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageLocation[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageLocation[i3].Key),
				Extract:      resource.ExtractParamPath("key", false),
				Reference:    mg.Spec.InitProvider.StorageLocation[i3].KeyRef,
				Selector:     mg.Spec.InitProvider.StorageLocation[i3].KeySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageLocation[i3].Key")
		}
		mg.Spec.InitProvider.StorageLocation[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageLocation[i3].KeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageLocation[i3].RoleArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.StorageLocation[i3].RoleArnRef,
				Selector:     mg.Spec.InitProvider.StorageLocation[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageLocation[i3].RoleArn")
		}
		mg.Spec.InitProvider.StorageLocation[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageLocation[i3].RoleArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Fleet.
func (mg *Fleet) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("gamelift.aws.upbound.io", "v1beta1", "Build", "BuildList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.BuildIDRef,
			Selector:     mg.Spec.ForProvider.BuildIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BuildID")
	}
	mg.Spec.ForProvider.BuildID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BuildIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.InstanceRoleArnRef,
			Selector:     mg.Spec.ForProvider.InstanceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceRoleArn")
	}
	mg.Spec.ForProvider.InstanceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("gamelift.aws.upbound.io", "v1beta1", "Build", "BuildList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.BuildIDRef,
			Selector:     mg.Spec.InitProvider.BuildIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BuildID")
	}
	mg.Spec.InitProvider.BuildID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BuildIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.InstanceRoleArnRef,
			Selector:     mg.Spec.InitProvider.InstanceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceRoleArn")
	}
	mg.Spec.InitProvider.InstanceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GameSessionQueue.
func (mg *GameSessionQueue) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("gamelift.aws.upbound.io", "v1beta2", "Fleet", "FleetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Destinations),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.ForProvider.DestinationsRefs,
			Selector:      mg.Spec.ForProvider.DestinationsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Destinations")
	}
	mg.Spec.ForProvider.Destinations = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.DestinationsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NotificationTarget),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.NotificationTargetRef,
			Selector:     mg.Spec.ForProvider.NotificationTargetSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NotificationTarget")
	}
	mg.Spec.ForProvider.NotificationTarget = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NotificationTargetRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("gamelift.aws.upbound.io", "v1beta2", "Fleet", "FleetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Destinations),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.InitProvider.DestinationsRefs,
			Selector:      mg.Spec.InitProvider.DestinationsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Destinations")
	}
	mg.Spec.InitProvider.Destinations = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.DestinationsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NotificationTarget),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.NotificationTargetRef,
			Selector:     mg.Spec.InitProvider.NotificationTargetSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NotificationTarget")
	}
	mg.Spec.InitProvider.NotificationTarget = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NotificationTargetRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Script.
func (mg *Script) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Bucket),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.StorageLocation[i3].BucketRef,
				Selector:     mg.Spec.ForProvider.StorageLocation[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Bucket")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].Key),
				Extract:      resource.ExtractParamPath("key", false),
				Reference:    mg.Spec.ForProvider.StorageLocation[i3].KeyRef,
				Selector:     mg.Spec.ForProvider.StorageLocation[i3].KeySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].Key")
		}
		mg.Spec.ForProvider.StorageLocation[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].KeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLocation[i3].RoleArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef,
				Selector:     mg.Spec.ForProvider.StorageLocation[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StorageLocation[i3].RoleArn")
		}
		mg.Spec.ForProvider.StorageLocation[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StorageLocation[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageLocation[i3].Bucket),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.StorageLocation[i3].BucketRef,
				Selector:     mg.Spec.InitProvider.StorageLocation[i3].BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageLocation[i3].Bucket")
		}
		mg.Spec.InitProvider.StorageLocation[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageLocation[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Object", "ObjectList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageLocation[i3].Key),
				Extract:      resource.ExtractParamPath("key", false),
				Reference:    mg.Spec.InitProvider.StorageLocation[i3].KeyRef,
				Selector:     mg.Spec.InitProvider.StorageLocation[i3].KeySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageLocation[i3].Key")
		}
		mg.Spec.InitProvider.StorageLocation[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageLocation[i3].KeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLocation); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageLocation[i3].RoleArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.StorageLocation[i3].RoleArnRef,
				Selector:     mg.Spec.InitProvider.StorageLocation[i3].RoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StorageLocation[i3].RoleArn")
		}
		mg.Spec.InitProvider.StorageLocation[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StorageLocation[i3].RoleArnRef = rsp.ResolvedReference

	}

	return nil
}
