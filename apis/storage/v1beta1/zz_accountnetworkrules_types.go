// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccountNetworkRulesInitParameters struct {

	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of Logging, Metrics, AzureServices, or None. Defaults to ["AzureServices"].
	// +listType=set
	Bypass []*string `json:"bypass,omitempty" tf:"bypass,omitempty"`

	// Specifies the default action of allow or deny when no other rules match. Valid options are Deny or Allow.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// List of public IP or IP ranges in CIDR Format. Only IPv4 addresses are allowed. Private IP address ranges (as defined in RFC 1918) are not allowed.
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// One or more private_link_access block as defined below.
	PrivateLinkAccess []AccountNetworkRulesPrivateLinkAccessInitParameters `json:"privateLinkAccess,omitempty" tf:"private_link_access,omitempty"`

	// Specifies the ID of the storage account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`

	// A list of virtual network subnet ids to secure the storage account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +listType=set
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds,omitempty" tf:"virtual_network_subnet_ids,omitempty"`

	// References to Subnet in network to populate virtualNetworkSubnetIds.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIdsRefs []v1.Reference `json:"virtualNetworkSubnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in network to populate virtualNetworkSubnetIds.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIdsSelector *v1.Selector `json:"virtualNetworkSubnetIdsSelector,omitempty" tf:"-"`
}

type AccountNetworkRulesObservation struct {

	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of Logging, Metrics, AzureServices, or None. Defaults to ["AzureServices"].
	// +listType=set
	Bypass []*string `json:"bypass,omitempty" tf:"bypass,omitempty"`

	// Specifies the default action of allow or deny when no other rules match. Valid options are Deny or Allow.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// The ID of the Storage Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of public IP or IP ranges in CIDR Format. Only IPv4 addresses are allowed. Private IP address ranges (as defined in RFC 1918) are not allowed.
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// One or more private_link_access block as defined below.
	PrivateLinkAccess []AccountNetworkRulesPrivateLinkAccessObservation `json:"privateLinkAccess,omitempty" tf:"private_link_access,omitempty"`

	// Specifies the ID of the storage account. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// A list of virtual network subnet ids to secure the storage account.
	// +listType=set
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type AccountNetworkRulesParameters struct {

	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of Logging, Metrics, AzureServices, or None. Defaults to ["AzureServices"].
	// +kubebuilder:validation:Optional
	// +listType=set
	Bypass []*string `json:"bypass,omitempty" tf:"bypass,omitempty"`

	// Specifies the default action of allow or deny when no other rules match. Valid options are Deny or Allow.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// List of public IP or IP ranges in CIDR Format. Only IPv4 addresses are allowed. Private IP address ranges (as defined in RFC 1918) are not allowed.
	// +kubebuilder:validation:Optional
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// One or more private_link_access block as defined below.
	// +kubebuilder:validation:Optional
	PrivateLinkAccess []AccountNetworkRulesPrivateLinkAccessParameters `json:"privateLinkAccess,omitempty" tf:"private_link_access,omitempty"`

	// Specifies the ID of the storage account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`

	// A list of virtual network subnet ids to secure the storage account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds,omitempty" tf:"virtual_network_subnet_ids,omitempty"`

	// References to Subnet in network to populate virtualNetworkSubnetIds.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIdsRefs []v1.Reference `json:"virtualNetworkSubnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in network to populate virtualNetworkSubnetIds.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIdsSelector *v1.Selector `json:"virtualNetworkSubnetIdsSelector,omitempty" tf:"-"`
}

type AccountNetworkRulesPrivateLinkAccessInitParameters struct {

	// The resource id of the resource access rule to be granted access.
	EndpointResourceID *string `json:"endpointResourceId,omitempty" tf:"endpoint_resource_id,omitempty"`

	// The tenant id of the resource of the resource access rule to be granted access. Defaults to the current tenant id.
	EndpointTenantID *string `json:"endpointTenantId,omitempty" tf:"endpoint_tenant_id,omitempty"`
}

type AccountNetworkRulesPrivateLinkAccessObservation struct {

	// The resource id of the resource access rule to be granted access.
	EndpointResourceID *string `json:"endpointResourceId,omitempty" tf:"endpoint_resource_id,omitempty"`

	// The tenant id of the resource of the resource access rule to be granted access. Defaults to the current tenant id.
	EndpointTenantID *string `json:"endpointTenantId,omitempty" tf:"endpoint_tenant_id,omitempty"`
}

type AccountNetworkRulesPrivateLinkAccessParameters struct {

	// The resource id of the resource access rule to be granted access.
	// +kubebuilder:validation:Optional
	EndpointResourceID *string `json:"endpointResourceId" tf:"endpoint_resource_id,omitempty"`

	// The tenant id of the resource of the resource access rule to be granted access. Defaults to the current tenant id.
	// +kubebuilder:validation:Optional
	EndpointTenantID *string `json:"endpointTenantId,omitempty" tf:"endpoint_tenant_id,omitempty"`
}

// AccountNetworkRulesSpec defines the desired state of AccountNetworkRules
type AccountNetworkRulesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountNetworkRulesParameters `json:"forProvider"`
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
	InitProvider AccountNetworkRulesInitParameters `json:"initProvider,omitempty"`
}

// AccountNetworkRulesStatus defines the observed state of AccountNetworkRules.
type AccountNetworkRulesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountNetworkRulesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccountNetworkRules is the Schema for the AccountNetworkRuless API. Manages network rules inside of a Azure Storage Account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AccountNetworkRules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultAction) || (has(self.initProvider) && has(self.initProvider.defaultAction))",message="spec.forProvider.defaultAction is a required parameter"
	Spec   AccountNetworkRulesSpec   `json:"spec"`
	Status AccountNetworkRulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountNetworkRulesList contains a list of AccountNetworkRuless
type AccountNetworkRulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountNetworkRules `json:"items"`
}

// Repository type metadata.
var (
	AccountNetworkRules_Kind             = "AccountNetworkRules"
	AccountNetworkRules_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountNetworkRules_Kind}.String()
	AccountNetworkRules_KindAPIVersion   = AccountNetworkRules_Kind + "." + CRDGroupVersion.String()
	AccountNetworkRules_GroupVersionKind = CRDGroupVersion.WithKind(AccountNetworkRules_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountNetworkRules{}, &AccountNetworkRulesList{})
}
