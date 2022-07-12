/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MSSQLManagedInstanceFailoverGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PartnerRegion []PartnerRegionObservation `json:"partnerRegion,omitempty" tf:"partner_region,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type MSSQLManagedInstanceFailoverGroupParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedInstanceIDRef *v1.Reference `json:"managedInstanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ManagedInstanceIDSelector *v1.Selector `json:"managedInstanceIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PartnerManagedInstanceID *string `json:"partnerManagedInstanceId,omitempty" tf:"partner_managed_instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	PartnerManagedInstanceIDRef *v1.Reference `json:"partnerManagedInstanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PartnerManagedInstanceIDSelector *v1.Selector `json:"partnerManagedInstanceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ReadWriteEndpointFailoverPolicy []ReadWriteEndpointFailoverPolicyParameters `json:"readWriteEndpointFailoverPolicy" tf:"read_write_endpoint_failover_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ReadonlyEndpointFailoverPolicyEnabled *bool `json:"readonlyEndpointFailoverPolicyEnabled,omitempty" tf:"readonly_endpoint_failover_policy_enabled,omitempty"`
}

type PartnerRegionObservation struct {
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type PartnerRegionParameters struct {
}

type ReadWriteEndpointFailoverPolicyObservation struct {
}

type ReadWriteEndpointFailoverPolicyParameters struct {

	// +kubebuilder:validation:Optional
	GraceMinutes *float64 `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

// MSSQLManagedInstanceFailoverGroupSpec defines the desired state of MSSQLManagedInstanceFailoverGroup
type MSSQLManagedInstanceFailoverGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLManagedInstanceFailoverGroupParameters `json:"forProvider"`
}

// MSSQLManagedInstanceFailoverGroupStatus defines the observed state of MSSQLManagedInstanceFailoverGroup.
type MSSQLManagedInstanceFailoverGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLManagedInstanceFailoverGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedInstanceFailoverGroup is the Schema for the MSSQLManagedInstanceFailoverGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLManagedInstanceFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLManagedInstanceFailoverGroupSpec   `json:"spec"`
	Status            MSSQLManagedInstanceFailoverGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedInstanceFailoverGroupList contains a list of MSSQLManagedInstanceFailoverGroups
type MSSQLManagedInstanceFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLManagedInstanceFailoverGroup `json:"items"`
}

// Repository type metadata.
var (
	MSSQLManagedInstanceFailoverGroup_Kind             = "MSSQLManagedInstanceFailoverGroup"
	MSSQLManagedInstanceFailoverGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLManagedInstanceFailoverGroup_Kind}.String()
	MSSQLManagedInstanceFailoverGroup_KindAPIVersion   = MSSQLManagedInstanceFailoverGroup_Kind + "." + CRDGroupVersion.String()
	MSSQLManagedInstanceFailoverGroup_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLManagedInstanceFailoverGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLManagedInstanceFailoverGroup{}, &MSSQLManagedInstanceFailoverGroupList{})
}
