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

type CustomHTTPSConfigurationInitParameters struct {

	// The name of the Key Vault secret representing the full certificate PFX.
	AzureKeyVaultCertificateSecretName *string `json:"azureKeyVaultCertificateSecretName,omitempty" tf:"azure_key_vault_certificate_secret_name,omitempty"`

	// The version of the Key Vault secret representing the full certificate PFX.
	AzureKeyVaultCertificateSecretVersion *string `json:"azureKeyVaultCertificateSecretVersion,omitempty" tf:"azure_key_vault_certificate_secret_version,omitempty"`

	// The ID of the Key Vault containing the SSL certificate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta2.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	AzureKeyVaultCertificateVaultID *string `json:"azureKeyVaultCertificateVaultId,omitempty" tf:"azure_key_vault_certificate_vault_id,omitempty"`

	// Reference to a Key in keyvault to populate azureKeyVaultCertificateVaultId.
	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateVaultIDRef *v1.Reference `json:"azureKeyVaultCertificateVaultIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate azureKeyVaultCertificateVaultId.
	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateVaultIDSelector *v1.Selector `json:"azureKeyVaultCertificateVaultIdSelector,omitempty" tf:"-"`

	// Certificate source to encrypted HTTPS traffic with. Allowed values are FrontDoor or AzureKeyVault. Defaults to FrontDoor.
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`
}

type CustomHTTPSConfigurationObservation struct {

	// The name of the Key Vault secret representing the full certificate PFX.
	AzureKeyVaultCertificateSecretName *string `json:"azureKeyVaultCertificateSecretName,omitempty" tf:"azure_key_vault_certificate_secret_name,omitempty"`

	// The version of the Key Vault secret representing the full certificate PFX.
	AzureKeyVaultCertificateSecretVersion *string `json:"azureKeyVaultCertificateSecretVersion,omitempty" tf:"azure_key_vault_certificate_secret_version,omitempty"`

	// The ID of the Key Vault containing the SSL certificate.
	AzureKeyVaultCertificateVaultID *string `json:"azureKeyVaultCertificateVaultId,omitempty" tf:"azure_key_vault_certificate_vault_id,omitempty"`

	// Certificate source to encrypted HTTPS traffic with. Allowed values are FrontDoor or AzureKeyVault. Defaults to FrontDoor.
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`

	// Minimum client TLS version supported.
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	ProvisioningState *string `json:"provisioningState,omitempty" tf:"provisioning_state,omitempty"`

	ProvisioningSubstate *string `json:"provisioningSubstate,omitempty" tf:"provisioning_substate,omitempty"`
}

type CustomHTTPSConfigurationParameters struct {

	// The name of the Key Vault secret representing the full certificate PFX.
	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateSecretName *string `json:"azureKeyVaultCertificateSecretName,omitempty" tf:"azure_key_vault_certificate_secret_name,omitempty"`

	// The version of the Key Vault secret representing the full certificate PFX.
	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateSecretVersion *string `json:"azureKeyVaultCertificateSecretVersion,omitempty" tf:"azure_key_vault_certificate_secret_version,omitempty"`

	// The ID of the Key Vault containing the SSL certificate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta2.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateVaultID *string `json:"azureKeyVaultCertificateVaultId,omitempty" tf:"azure_key_vault_certificate_vault_id,omitempty"`

	// Reference to a Key in keyvault to populate azureKeyVaultCertificateVaultId.
	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateVaultIDRef *v1.Reference `json:"azureKeyVaultCertificateVaultIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate azureKeyVaultCertificateVaultId.
	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateVaultIDSelector *v1.Selector `json:"azureKeyVaultCertificateVaultIdSelector,omitempty" tf:"-"`

	// Certificate source to encrypted HTTPS traffic with. Allowed values are FrontDoor or AzureKeyVault. Defaults to FrontDoor.
	// +kubebuilder:validation:Optional
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`
}

type FrontdoorCustomHTTPSConfigurationInitParameters struct {

	// A custom_https_configuration block as defined above.
	CustomHTTPSConfiguration *CustomHTTPSConfigurationInitParameters `json:"customHttpsConfiguration,omitempty" tf:"custom_https_configuration,omitempty"`

	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	CustomHTTPSProvisioningEnabled *bool `json:"customHttpsProvisioningEnabled,omitempty" tf:"custom_https_provisioning_enabled,omitempty"`

	// The ID of the Front Door Frontend Endpoint which this configuration refers to. Changing this forces a new resource to be created.
	FrontendEndpointID *string `json:"frontendEndpointId,omitempty" tf:"frontend_endpoint_id,omitempty"`
}

type FrontdoorCustomHTTPSConfigurationObservation struct {

	// A custom_https_configuration block as defined above.
	CustomHTTPSConfiguration *CustomHTTPSConfigurationObservation `json:"customHttpsConfiguration,omitempty" tf:"custom_https_configuration,omitempty"`

	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	CustomHTTPSProvisioningEnabled *bool `json:"customHttpsProvisioningEnabled,omitempty" tf:"custom_https_provisioning_enabled,omitempty"`

	// The ID of the Front Door Frontend Endpoint which this configuration refers to. Changing this forces a new resource to be created.
	FrontendEndpointID *string `json:"frontendEndpointId,omitempty" tf:"frontend_endpoint_id,omitempty"`

	// The ID of the Azure Front Door Custom HTTPS Configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FrontdoorCustomHTTPSConfigurationParameters struct {

	// A custom_https_configuration block as defined above.
	// +kubebuilder:validation:Optional
	CustomHTTPSConfiguration *CustomHTTPSConfigurationParameters `json:"customHttpsConfiguration,omitempty" tf:"custom_https_configuration,omitempty"`

	// Should the HTTPS protocol be enabled for this custom domain associated with the Front Door?
	// +kubebuilder:validation:Optional
	CustomHTTPSProvisioningEnabled *bool `json:"customHttpsProvisioningEnabled,omitempty" tf:"custom_https_provisioning_enabled,omitempty"`

	// The ID of the Front Door Frontend Endpoint which this configuration refers to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	FrontendEndpointID *string `json:"frontendEndpointId,omitempty" tf:"frontend_endpoint_id,omitempty"`
}

// FrontdoorCustomHTTPSConfigurationSpec defines the desired state of FrontdoorCustomHTTPSConfiguration
type FrontdoorCustomHTTPSConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorCustomHTTPSConfigurationParameters `json:"forProvider"`
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
	InitProvider FrontdoorCustomHTTPSConfigurationInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorCustomHTTPSConfigurationStatus defines the observed state of FrontdoorCustomHTTPSConfiguration.
type FrontdoorCustomHTTPSConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorCustomHTTPSConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FrontdoorCustomHTTPSConfiguration is the Schema for the FrontdoorCustomHTTPSConfigurations API. Manages the Custom Https Configuration for an Azure Front Door (classic) Frontend Endpoint.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorCustomHTTPSConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.customHttpsProvisioningEnabled) || (has(self.initProvider) && has(self.initProvider.customHttpsProvisioningEnabled))",message="spec.forProvider.customHttpsProvisioningEnabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.frontendEndpointId) || (has(self.initProvider) && has(self.initProvider.frontendEndpointId))",message="spec.forProvider.frontendEndpointId is a required parameter"
	Spec   FrontdoorCustomHTTPSConfigurationSpec   `json:"spec"`
	Status FrontdoorCustomHTTPSConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorCustomHTTPSConfigurationList contains a list of FrontdoorCustomHTTPSConfigurations
type FrontdoorCustomHTTPSConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorCustomHTTPSConfiguration `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorCustomHTTPSConfiguration_Kind             = "FrontdoorCustomHTTPSConfiguration"
	FrontdoorCustomHTTPSConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorCustomHTTPSConfiguration_Kind}.String()
	FrontdoorCustomHTTPSConfiguration_KindAPIVersion   = FrontdoorCustomHTTPSConfiguration_Kind + "." + CRDGroupVersion.String()
	FrontdoorCustomHTTPSConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorCustomHTTPSConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorCustomHTTPSConfiguration{}, &FrontdoorCustomHTTPSConfigurationList{})
}