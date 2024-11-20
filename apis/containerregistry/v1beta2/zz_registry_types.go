// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EncryptionInitParameters struct {

	// Boolean value that indicates whether the policy is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// The client ID of the managed identity associated with the encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("client_id",true)
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id"`

	// Reference to a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDRef *v1.Reference `json:"identityClientIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDSelector *v1.Selector `json:"identityClientIdSelector,omitempty" tf:"-"`

	// The ID of the Key Vault Key.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id"`
}

type EncryptionObservation struct {

	// Boolean value that indicates whether the policy is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The client ID of the managed identity associated with the encryption key.
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// The ID of the Key Vault Key.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`
}

type EncryptionParameters struct {

	// Boolean value that indicates whether the policy is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// The client ID of the managed identity associated with the encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("client_id",true)
	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id"`

	// Reference to a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDRef *v1.Reference `json:"identityClientIdRef,omitempty" tf:"-"`

	// Selector for a UserAssignedIdentity in managedidentity to populate identityClientId.
	// +kubebuilder:validation:Optional
	IdentityClientIDSelector *v1.Selector `json:"identityClientIdSelector,omitempty" tf:"-"`

	// The ID of the Key Vault Key.
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id"`
}

type GeoreplicationsInitParameters struct {

	// A location where the container registry should be geo-replicated.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether regional endpoint is enabled for this Container Registry?
	RegionalEndpointEnabled *bool `json:"regionalEndpointEnabled,omitempty" tf:"regional_endpoint_enabled,omitempty"`

	// A mapping of tags to assign to this replication location.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false.
	ZoneRedundancyEnabled *bool `json:"zoneRedundancyEnabled,omitempty" tf:"zone_redundancy_enabled,omitempty"`
}

type GeoreplicationsObservation struct {

	// A location where the container registry should be geo-replicated.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether regional endpoint is enabled for this Container Registry?
	RegionalEndpointEnabled *bool `json:"regionalEndpointEnabled,omitempty" tf:"regional_endpoint_enabled,omitempty"`

	// A mapping of tags to assign to this replication location.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false.
	ZoneRedundancyEnabled *bool `json:"zoneRedundancyEnabled,omitempty" tf:"zone_redundancy_enabled,omitempty"`
}

type GeoreplicationsParameters struct {

	// A location where the container registry should be geo-replicated.
	// +kubebuilder:validation:Optional
	Location *string `json:"location" tf:"location,omitempty"`

	// Whether regional endpoint is enabled for this Container Registry?
	// +kubebuilder:validation:Optional
	RegionalEndpointEnabled *bool `json:"regionalEndpointEnabled,omitempty" tf:"regional_endpoint_enabled,omitempty"`

	// A mapping of tags to assign to this replication location.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false.
	// +kubebuilder:validation:Optional
	ZoneRedundancyEnabled *bool `json:"zoneRedundancyEnabled,omitempty" tf:"zone_redundancy_enabled,omitempty"`
}

type IPRuleInitParameters struct {

	// The behaviour for requests matching this rule. At this time the only supported value is Allow
	Action *string `json:"action,omitempty" tf:"action"`

	// The CIDR block from which requests will match the rule.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range"`
}

type IPRuleObservation struct {

	// The behaviour for requests matching this rule. At this time the only supported value is Allow
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The CIDR block from which requests will match the rule.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`
}

type IPRuleParameters struct {

	// The behaviour for requests matching this rule. At this time the only supported value is Allow
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// The CIDR block from which requests will match the rule.
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Container Registry.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// References to UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsRefs []v1.Reference `json:"identityIdsRefs,omitempty" tf:"-"`

	// Selector for a list of UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsSelector *v1.Selector `json:"identityIdsSelector,omitempty" tf:"-"`

	// Specifies the type of Managed Service Identity that should be configured on this Container Registry. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Container Registry.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Container Registry. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Container Registry.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// References to UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsRefs []v1.Reference `json:"identityIdsRefs,omitempty" tf:"-"`

	// Selector for a list of UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsSelector *v1.Selector `json:"identityIdsSelector,omitempty" tf:"-"`

	// Specifies the type of Managed Service Identity that should be configured on this Container Registry. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkRuleSetInitParameters struct {

	// The behaviour for requests matching no rules. Either Allow or Deny. Defaults to Allow
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action"`

	// One or more ip_rule blocks as defined below.
	IPRule []IPRuleInitParameters `json:"ipRule,omitempty" tf:"ip_rule"`

	VirtualNetwork []VirtualNetworkInitParameters `json:"virtualNetwork,omitempty" tf:"virtual_network"`
}

type NetworkRuleSetObservation struct {

	// The behaviour for requests matching no rules. Either Allow or Deny. Defaults to Allow
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more ip_rule blocks as defined below.
	IPRule []IPRuleObservation `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`

	VirtualNetwork []VirtualNetworkObservation `json:"virtualNetwork,omitempty" tf:"virtual_network,omitempty"`
}

type NetworkRuleSetParameters struct {

	// The behaviour for requests matching no rules. Either Allow or Deny. Defaults to Allow
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action"`

	// One or more ip_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	IPRule []IPRuleParameters `json:"ipRule,omitempty" tf:"ip_rule"`

	// +kubebuilder:validation:Optional
	VirtualNetwork []VirtualNetworkParameters `json:"virtualNetwork,omitempty" tf:"virtual_network"`
}

type RegistryInitParameters struct {

	// Specifies whether the admin user is enabled. Defaults to false.
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// Whether allows anonymous (unauthenticated) pull access to this Container Registry? This is only supported on resources with the Standard or Premium SKU.
	AnonymousPullEnabled *bool `json:"anonymousPullEnabled,omitempty" tf:"anonymous_pull_enabled,omitempty"`

	// Whether to enable dedicated data endpoints for this Container Registry? This is only supported on resources with the Premium SKU.
	DataEndpointEnabled *bool `json:"dataEndpointEnabled,omitempty" tf:"data_endpoint_enabled,omitempty"`

	// An encryption block as documented below.
	Encryption *EncryptionInitParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Boolean value that indicates whether export policy is enabled. Defaults to true. In order to set it to false, make sure the public_network_access_enabled is also set to false.
	ExportPolicyEnabled *bool `json:"exportPolicyEnabled,omitempty" tf:"export_policy_enabled,omitempty"`

	// A georeplications block as documented below.
	Georeplications []GeoreplicationsInitParameters `json:"georeplications,omitempty" tf:"georeplications,omitempty"`

	// An identity block as defined below.
	Identity *IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether to allow trusted Azure services to access a network restricted Container Registry? Possible values are None and AzureServices. Defaults to AzureServices.
	NetworkRuleBypassOption *string `json:"networkRuleBypassOption,omitempty" tf:"network_rule_bypass_option,omitempty"`

	// A network_rule_set block as documented below.
	NetworkRuleSet *NetworkRuleSetInitParameters `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`

	// Whether public network access is allowed for the container registry. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Boolean value that indicates whether quarantine policy is enabled.
	QuarantinePolicyEnabled *bool `json:"quarantinePolicyEnabled,omitempty" tf:"quarantine_policy_enabled,omitempty"`

	// A retention_policy block as documented below.
	RetentionPolicy *RetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The SKU name of the container registry. Possible values are Basic, Standard and Premium.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A trust_policy block as documented below.
	TrustPolicy *TrustPolicyInitParameters `json:"trustPolicy,omitempty" tf:"trust_policy,omitempty"`

	// Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false.
	ZoneRedundancyEnabled *bool `json:"zoneRedundancyEnabled,omitempty" tf:"zone_redundancy_enabled,omitempty"`
}

type RegistryObservation struct {

	// Specifies whether the admin user is enabled. Defaults to false.
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// The Username associated with the Container Registry Admin account - if the admin account is enabled.
	AdminUsername *string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`

	// Whether allows anonymous (unauthenticated) pull access to this Container Registry? This is only supported on resources with the Standard or Premium SKU.
	AnonymousPullEnabled *bool `json:"anonymousPullEnabled,omitempty" tf:"anonymous_pull_enabled,omitempty"`

	// Whether to enable dedicated data endpoints for this Container Registry? This is only supported on resources with the Premium SKU.
	DataEndpointEnabled *bool `json:"dataEndpointEnabled,omitempty" tf:"data_endpoint_enabled,omitempty"`

	// An encryption block as documented below.
	Encryption *EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Boolean value that indicates whether export policy is enabled. Defaults to true. In order to set it to false, make sure the public_network_access_enabled is also set to false.
	ExportPolicyEnabled *bool `json:"exportPolicyEnabled,omitempty" tf:"export_policy_enabled,omitempty"`

	// A georeplications block as documented below.
	Georeplications []GeoreplicationsObservation `json:"georeplications,omitempty" tf:"georeplications,omitempty"`

	// The ID of the Container Registry.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity *IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The URL that can be used to log into the container registry.
	LoginServer *string `json:"loginServer,omitempty" tf:"login_server,omitempty"`

	// Whether to allow trusted Azure services to access a network restricted Container Registry? Possible values are None and AzureServices. Defaults to AzureServices.
	NetworkRuleBypassOption *string `json:"networkRuleBypassOption,omitempty" tf:"network_rule_bypass_option,omitempty"`

	// A network_rule_set block as documented below.
	NetworkRuleSet *NetworkRuleSetObservation `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`

	// Whether public network access is allowed for the container registry. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Boolean value that indicates whether quarantine policy is enabled.
	QuarantinePolicyEnabled *bool `json:"quarantinePolicyEnabled,omitempty" tf:"quarantine_policy_enabled,omitempty"`

	// The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A retention_policy block as documented below.
	RetentionPolicy *RetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The SKU name of the container registry. Possible values are Basic, Standard and Premium.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A trust_policy block as documented below.
	TrustPolicy *TrustPolicyObservation `json:"trustPolicy,omitempty" tf:"trust_policy,omitempty"`

	// Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false.
	ZoneRedundancyEnabled *bool `json:"zoneRedundancyEnabled,omitempty" tf:"zone_redundancy_enabled,omitempty"`
}

type RegistryParameters struct {

	// Specifies whether the admin user is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// Whether allows anonymous (unauthenticated) pull access to this Container Registry? This is only supported on resources with the Standard or Premium SKU.
	// +kubebuilder:validation:Optional
	AnonymousPullEnabled *bool `json:"anonymousPullEnabled,omitempty" tf:"anonymous_pull_enabled,omitempty"`

	// Whether to enable dedicated data endpoints for this Container Registry? This is only supported on resources with the Premium SKU.
	// +kubebuilder:validation:Optional
	DataEndpointEnabled *bool `json:"dataEndpointEnabled,omitempty" tf:"data_endpoint_enabled,omitempty"`

	// An encryption block as documented below.
	// +kubebuilder:validation:Optional
	Encryption *EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Boolean value that indicates whether export policy is enabled. Defaults to true. In order to set it to false, make sure the public_network_access_enabled is also set to false.
	// +kubebuilder:validation:Optional
	ExportPolicyEnabled *bool `json:"exportPolicyEnabled,omitempty" tf:"export_policy_enabled,omitempty"`

	// A georeplications block as documented below.
	// +kubebuilder:validation:Optional
	Georeplications []GeoreplicationsParameters `json:"georeplications,omitempty" tf:"georeplications,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether to allow trusted Azure services to access a network restricted Container Registry? Possible values are None and AzureServices. Defaults to AzureServices.
	// +kubebuilder:validation:Optional
	NetworkRuleBypassOption *string `json:"networkRuleBypassOption,omitempty" tf:"network_rule_bypass_option,omitempty"`

	// A network_rule_set block as documented below.
	// +kubebuilder:validation:Optional
	NetworkRuleSet *NetworkRuleSetParameters `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`

	// Whether public network access is allowed for the container registry. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// Boolean value that indicates whether quarantine policy is enabled.
	// +kubebuilder:validation:Optional
	QuarantinePolicyEnabled *bool `json:"quarantinePolicyEnabled,omitempty" tf:"quarantine_policy_enabled,omitempty"`

	// The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A retention_policy block as documented below.
	// +kubebuilder:validation:Optional
	RetentionPolicy *RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The SKU name of the container registry. Possible values are Basic, Standard and Premium.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A trust_policy block as documented below.
	// +kubebuilder:validation:Optional
	TrustPolicy *TrustPolicyParameters `json:"trustPolicy,omitempty" tf:"trust_policy,omitempty"`

	// Whether zone redundancy is enabled for this Container Registry? Changing this forces a new resource to be created. Defaults to false.
	// +kubebuilder:validation:Optional
	ZoneRedundancyEnabled *bool `json:"zoneRedundancyEnabled,omitempty" tf:"zone_redundancy_enabled,omitempty"`
}

type RetentionPolicyInitParameters struct {

	// The number of days to retain an untagged manifest after which it gets purged. Default is 7.
	Days *float64 `json:"days,omitempty" tf:"days"`

	// Boolean value that indicates whether the policy is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type RetentionPolicyObservation struct {

	// The number of days to retain an untagged manifest after which it gets purged. Default is 7.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Boolean value that indicates whether the policy is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RetentionPolicyParameters struct {

	// The number of days to retain an untagged manifest after which it gets purged. Default is 7.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days"`

	// Boolean value that indicates whether the policy is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type TrustPolicyInitParameters struct {

	// Boolean value that indicates whether the policy is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type TrustPolicyObservation struct {

	// Boolean value that indicates whether the policy is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TrustPolicyParameters struct {

	// Boolean value that indicates whether the policy is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type VirtualNetworkInitParameters struct {

	// The behaviour for requests matching this rule. At this time the only supported value is Allow
	Action *string `json:"action,omitempty" tf:"action"`

	// The ID of the Container Registry.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualNetworkObservation struct {

	// The behaviour for requests matching this rule. At this time the only supported value is Allow
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of the Container Registry.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type VirtualNetworkParameters struct {

	// The behaviour for requests matching this rule. At this time the only supported value is Allow
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// The ID of the Container Registry.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// RegistrySpec defines the desired state of Registry
type RegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RegistryInitParameters `json:"initProvider,omitempty"`
}

// RegistryStatus defines the observed state of Registry.
type RegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Registry is the Schema for the Registrys API. Manages an Azure Container Registry.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   RegistrySpec   `json:"spec"`
	Status RegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryList contains a list of Registrys
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Repository type metadata.
var (
	Registry_Kind             = "Registry"
	Registry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Registry_Kind}.String()
	Registry_KindAPIVersion   = Registry_Kind + "." + CRDGroupVersion.String()
	Registry_GroupVersionKind = CRDGroupVersion.WithKind(Registry_Kind)
)

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
