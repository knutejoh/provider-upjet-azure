// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	rconfig "github.com/upbound/provider-azure/apis/rconfig"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ComputeCluster) ResolveReferences( // ResolveReferences of this ComputeCluster.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.ForProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.ForProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetResourceID")
	}
	mg.Spec.ForProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.InitProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.InitProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.InitProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetResourceID")
	}
	mg.Spec.InitProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ComputeInstance.
func (mg *ComputeInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.ForProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.ForProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetResourceID")
	}
	mg.Spec.ForProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SubnetResourceIDRef,
			Selector:     mg.Spec.InitProvider.SubnetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetResourceID")
	}
	mg.Spec.InitProvider.SubnetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SynapseSpark.
func (mg *SynapseSpark) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("machinelearningservices.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MachineLearningWorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.MachineLearningWorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.MachineLearningWorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MachineLearningWorkspaceID")
	}
	mg.Spec.ForProvider.MachineLearningWorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MachineLearningWorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("synapse.azure.upbound.io", "v1beta2", "SparkPool", "SparkPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SynapseSparkPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SynapseSparkPoolIDRef,
			Selector:     mg.Spec.ForProvider.SynapseSparkPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SynapseSparkPoolID")
	}
	mg.Spec.ForProvider.SynapseSparkPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SynapseSparkPoolIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("synapse.azure.upbound.io", "v1beta2", "SparkPool", "SparkPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SynapseSparkPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SynapseSparkPoolIDRef,
			Selector:     mg.Spec.InitProvider.SynapseSparkPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SynapseSparkPoolID")
	}
	mg.Spec.InitProvider.SynapseSparkPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SynapseSparkPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Workspace.
func (mg *Workspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("insights.azure.upbound.io", "v1beta1", "ApplicationInsights", "ApplicationInsightsList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationInsightsID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ApplicationInsightsIDRef,
			Selector:     mg.Spec.ForProvider.ApplicationInsightsIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationInsightsID")
	}
	mg.Spec.ForProvider.ApplicationInsightsID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationInsightsIDRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption.KeyID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Encryption.KeyIDRef,
				Selector:     mg.Spec.ForProvider.Encryption.KeyIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption.KeyID")
		}
		mg.Spec.ForProvider.Encryption.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption.KeyIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption.KeyVaultID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Encryption.KeyVaultIDRef,
				Selector:     mg.Spec.ForProvider.Encryption.KeyVaultIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption.KeyVaultID")
		}
		mg.Spec.ForProvider.Encryption.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption.KeyVaultIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Encryption.UserAssignedIdentityID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Encryption.UserAssignedIdentityIDRef,
				Selector:     mg.Spec.ForProvider.Encryption.UserAssignedIdentityIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Encryption.UserAssignedIdentityID")
		}
		mg.Spec.ForProvider.Encryption.UserAssignedIdentityID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Encryption.UserAssignedIdentityIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.Identity != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Identity.IdentityIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.ForProvider.Identity.IdentityIdsRefs,
				Selector:      mg.Spec.ForProvider.Identity.IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Identity.IdentityIds")
		}
		mg.Spec.ForProvider.Identity.IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Identity.IdentityIdsRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyVaultID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyVaultIDRef,
			Selector:     mg.Spec.ForProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyVaultID")
	}
	mg.Spec.ForProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrimaryUserAssignedIdentity),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PrimaryUserAssignedIdentityRef,
			Selector:     mg.Spec.ForProvider.PrimaryUserAssignedIdentitySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrimaryUserAssignedIdentity")
	}
	mg.Spec.ForProvider.PrimaryUserAssignedIdentity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrimaryUserAssignedIdentityRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.ServerlessCompute != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerlessCompute.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.ServerlessCompute.SubnetIDRef,
				Selector:     mg.Spec.ForProvider.ServerlessCompute.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ServerlessCompute.SubnetID")
		}
		mg.Spec.ForProvider.ServerlessCompute.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ServerlessCompute.SubnetIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("insights.azure.upbound.io", "v1beta1", "ApplicationInsights", "ApplicationInsightsList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationInsightsID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ApplicationInsightsIDRef,
			Selector:     mg.Spec.InitProvider.ApplicationInsightsIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationInsightsID")
	}
	mg.Spec.InitProvider.ApplicationInsightsID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationInsightsIDRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption.KeyID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Encryption.KeyIDRef,
				Selector:     mg.Spec.InitProvider.Encryption.KeyIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption.KeyID")
		}
		mg.Spec.InitProvider.Encryption.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption.KeyIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption.KeyVaultID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Encryption.KeyVaultIDRef,
				Selector:     mg.Spec.InitProvider.Encryption.KeyVaultIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption.KeyVaultID")
		}
		mg.Spec.InitProvider.Encryption.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption.KeyVaultIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Encryption != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Encryption.UserAssignedIdentityID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Encryption.UserAssignedIdentityIDRef,
				Selector:     mg.Spec.InitProvider.Encryption.UserAssignedIdentityIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Encryption.UserAssignedIdentityID")
		}
		mg.Spec.InitProvider.Encryption.UserAssignedIdentityID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Encryption.UserAssignedIdentityIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Identity != nil {
		{
			m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Identity.IdentityIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.InitProvider.Identity.IdentityIdsRefs,
				Selector:      mg.Spec.InitProvider.Identity.IdentityIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Identity.IdentityIds")
		}
		mg.Spec.InitProvider.Identity.IdentityIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Identity.IdentityIdsRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("keyvault.azure.upbound.io", "v1beta2", "Vault", "VaultList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyVaultID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyVaultIDRef,
			Selector:     mg.Spec.InitProvider.KeyVaultIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyVaultID")
	}
	mg.Spec.InitProvider.KeyVaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyVaultIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("managedidentity.azure.upbound.io", "v1beta1", "UserAssignedIdentity", "UserAssignedIdentityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrimaryUserAssignedIdentity),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PrimaryUserAssignedIdentityRef,
			Selector:     mg.Spec.InitProvider.PrimaryUserAssignedIdentitySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrimaryUserAssignedIdentity")
	}
	mg.Spec.InitProvider.PrimaryUserAssignedIdentity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrimaryUserAssignedIdentityRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.ServerlessCompute != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerlessCompute.SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.ServerlessCompute.SubnetIDRef,
				Selector:     mg.Spec.InitProvider.ServerlessCompute.SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ServerlessCompute.SubnetID")
		}
		mg.Spec.InitProvider.ServerlessCompute.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ServerlessCompute.SubnetIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.StorageAccountIDRef,
			Selector:     mg.Spec.InitProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccountID")
	}
	mg.Spec.InitProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageAccountIDRef = rsp.ResolvedReference

	return nil
}
