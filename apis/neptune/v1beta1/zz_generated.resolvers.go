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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Cluster.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.IAMRoles),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.IAMRoleRefs,
			Selector:      mg.Spec.ForProvider.IAMRoleSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoles")
	}
	mg.Spec.ForProvider.IAMRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.IAMRoleRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
			Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterParameterGroup", "ClusterParameterGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneClusterParameterGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef,
			Selector:     mg.Spec.ForProvider.NeptuneClusterParameterGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneClusterParameterGroupName")
	}
	mg.Spec.ForProvider.NeptuneClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneSubnetGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NeptuneSubnetGroupNameRef,
			Selector:     mg.Spec.ForProvider.NeptuneSubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.ForProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicationSourceIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ReplicationSourceIdentifierRef,
			Selector:     mg.Spec.ForProvider.ReplicationSourceIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicationSourceIdentifier")
	}
	mg.Spec.ForProvider.ReplicationSourceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicationSourceIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterSnapshot", "ClusterSnapshotList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SnapshotIdentifierRef,
			Selector:     mg.Spec.ForProvider.SnapshotIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotIdentifier")
	}
	mg.Spec.ForProvider.SnapshotIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
			Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.IAMRoles),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.IAMRoleRefs,
			Selector:      mg.Spec.InitProvider.IAMRoleSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRoles")
	}
	mg.Spec.InitProvider.IAMRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.IAMRoleRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.KMSKeyArnRef,
			Selector:     mg.Spec.InitProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyArn")
	}
	mg.Spec.InitProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterParameterGroup", "ClusterParameterGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NeptuneClusterParameterGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NeptuneClusterParameterGroupNameRef,
			Selector:     mg.Spec.InitProvider.NeptuneClusterParameterGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NeptuneClusterParameterGroupName")
	}
	mg.Spec.InitProvider.NeptuneClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NeptuneClusterParameterGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NeptuneSubnetGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NeptuneSubnetGroupNameRef,
			Selector:     mg.Spec.InitProvider.NeptuneSubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.InitProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ReplicationSourceIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ReplicationSourceIdentifierRef,
			Selector:     mg.Spec.InitProvider.ReplicationSourceIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ReplicationSourceIdentifier")
	}
	mg.Spec.InitProvider.ReplicationSourceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ReplicationSourceIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterSnapshot", "ClusterSnapshotList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnapshotIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.SnapshotIdentifierRef,
			Selector:     mg.Spec.InitProvider.SnapshotIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SnapshotIdentifier")
	}
	mg.Spec.InitProvider.SnapshotIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SnapshotIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCSecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.VPCSecurityGroupIDRefs,
			Selector:      mg.Spec.InitProvider.VPCSecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCSecurityGroupIds")
	}
	mg.Spec.InitProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ClusterEndpoint.
func (mg *ClusterEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
			Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterIdentifierRef,
			Selector:     mg.Spec.InitProvider.ClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterIdentifier")
	}
	mg.Spec.InitProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterInstance.
func (mg *ClusterInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
			Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ParameterGroup", "ParameterGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneParameterGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NeptuneParameterGroupNameRef,
			Selector:     mg.Spec.ForProvider.NeptuneParameterGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneParameterGroupName")
	}
	mg.Spec.ForProvider.NeptuneParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneParameterGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneSubnetGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NeptuneSubnetGroupNameRef,
			Selector:     mg.Spec.ForProvider.NeptuneSubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.ForProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ClusterIdentifierRef,
			Selector:     mg.Spec.InitProvider.ClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterIdentifier")
	}
	mg.Spec.InitProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ParameterGroup", "ParameterGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NeptuneParameterGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NeptuneParameterGroupNameRef,
			Selector:     mg.Spec.InitProvider.NeptuneParameterGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NeptuneParameterGroupName")
	}
	mg.Spec.InitProvider.NeptuneParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NeptuneParameterGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NeptuneSubnetGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NeptuneSubnetGroupNameRef,
			Selector:     mg.Spec.InitProvider.NeptuneSubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.InitProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterSnapshot.
func (mg *ClusterSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DBClusterIdentifierRef,
			Selector:     mg.Spec.ForProvider.DBClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBClusterIdentifier")
	}
	mg.Spec.ForProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBClusterIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DBClusterIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DBClusterIdentifierRef,
			Selector:     mg.Spec.InitProvider.DBClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DBClusterIdentifier")
	}
	mg.Spec.InitProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventSubscription.
func (mg *EventSubscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnsTopicArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.SnsTopicArnRef,
			Selector:     mg.Spec.ForProvider.SnsTopicArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnsTopicArn")
	}
	mg.Spec.ForProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnsTopicArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterInstance", "ClusterInstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SourceIds),
			Extract:       resource.ExtractResourceID(),
			References:    mg.Spec.ForProvider.SourceIdsRefs,
			Selector:      mg.Spec.ForProvider.SourceIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceIds")
	}
	mg.Spec.ForProvider.SourceIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SourceIdsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io", "v1beta1", "Topic", "TopicList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnsTopicArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.SnsTopicArnRef,
			Selector:     mg.Spec.InitProvider.SnsTopicArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SnsTopicArn")
	}
	mg.Spec.InitProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SnsTopicArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterInstance", "ClusterInstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SourceIds),
			Extract:       resource.ExtractResourceID(),
			References:    mg.Spec.InitProvider.SourceIdsRefs,
			Selector:      mg.Spec.InitProvider.SourceIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SourceIds")
	}
	mg.Spec.InitProvider.SourceIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SourceIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this GlobalCluster.
func (mg *GlobalCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceDBClusterIdentifier),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.SourceDBClusterIdentifierRef,
			Selector:     mg.Spec.ForProvider.SourceDBClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceDBClusterIdentifier")
	}
	mg.Spec.ForProvider.SourceDBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceDBClusterIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceDBClusterIdentifier),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.SourceDBClusterIdentifierRef,
			Selector:     mg.Spec.InitProvider.SourceDBClusterIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SourceDBClusterIdentifier")
	}
	mg.Spec.InitProvider.SourceDBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SourceDBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetIDRefs,
			Selector:      mg.Spec.ForProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetIDRefs,
			Selector:      mg.Spec.InitProvider.SubnetIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}
