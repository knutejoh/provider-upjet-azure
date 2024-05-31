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

type ExpirationInitParameters struct {

	// The duration of the role assignment in days. Conflicts with schedule[0].expiration[0].duration_hours,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Active Role Assignment to be created.
	// The duration of the assignment in days.
	DurationDays *float64 `json:"durationDays,omitempty" tf:"duration_days,omitempty"`

	// The duration of the role assignment in hours. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Active Role Assignment to be created.
	// The duration of the assignment in hours.
	DurationHours *float64 `json:"durationHours,omitempty" tf:"duration_hours,omitempty"`

	// The end date time of the role assignment. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].duration_hours Changing this forces a new Pim Active Role Assignment to be created.
	// The end date time of the assignment.
	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`
}

type ExpirationObservation struct {

	// The duration of the role assignment in days. Conflicts with schedule[0].expiration[0].duration_hours,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Active Role Assignment to be created.
	// The duration of the assignment in days.
	DurationDays *float64 `json:"durationDays,omitempty" tf:"duration_days,omitempty"`

	// The duration of the role assignment in hours. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Active Role Assignment to be created.
	// The duration of the assignment in hours.
	DurationHours *float64 `json:"durationHours,omitempty" tf:"duration_hours,omitempty"`

	// The end date time of the role assignment. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].duration_hours Changing this forces a new Pim Active Role Assignment to be created.
	// The end date time of the assignment.
	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`
}

type ExpirationParameters struct {

	// The duration of the role assignment in days. Conflicts with schedule[0].expiration[0].duration_hours,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Active Role Assignment to be created.
	// The duration of the assignment in days.
	// +kubebuilder:validation:Optional
	DurationDays *float64 `json:"durationDays,omitempty" tf:"duration_days,omitempty"`

	// The duration of the role assignment in hours. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].end_date_time Changing this forces a new Pim Active Role Assignment to be created.
	// The duration of the assignment in hours.
	// +kubebuilder:validation:Optional
	DurationHours *float64 `json:"durationHours,omitempty" tf:"duration_hours,omitempty"`

	// The end date time of the role assignment. Conflicts with schedule[0].expiration[0].duration_days,schedule[0].expiration[0].duration_hours Changing this forces a new Pim Active Role Assignment to be created.
	// The end date time of the assignment.
	// +kubebuilder:validation:Optional
	EndDateTime *string `json:"endDateTime,omitempty" tf:"end_date_time,omitempty"`
}

type PimActiveRoleAssignmentInitParameters struct {

	// The justification of the role assignment. Changing this forces a new Pim Active Role Assignment to be created.
	// The justification of the role assignment.
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// The principal id. Changing this forces a new Pim Active Role Assignment to be created.
	// The principal id.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The role definition id. Changing this forces a new Pim Active Role Assignment to be created.
	// The role definition id.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// A schedule block as defined below. Changing this forces a new Pim Active Role Assignment to be created.
	// The schedule details of this role assignment.
	Schedule *ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The scope. Changing this forces a new Pim Active Role Assignment to be created.
	// The scope.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/management/v1beta1.ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Reference to a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeRef *v1.Reference `json:"scopeRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeSelector *v1.Selector `json:"scopeSelector,omitempty" tf:"-"`

	// A ticket block as defined below. Changing this forces a new Pim Active Role Assignment to be created.
	// The ticket details.
	Ticket *TicketInitParameters `json:"ticket,omitempty" tf:"ticket,omitempty"`
}

type PimActiveRoleAssignmentObservation struct {

	// The ID of the Pim Active Role Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The justification of the role assignment. Changing this forces a new Pim Active Role Assignment to be created.
	// The justification of the role assignment.
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// The principal id. Changing this forces a new Pim Active Role Assignment to be created.
	// The principal id.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The type of principal.
	// The type of principal.
	PrincipalType *string `json:"principalType,omitempty" tf:"principal_type,omitempty"`

	// The role definition id. Changing this forces a new Pim Active Role Assignment to be created.
	// The role definition id.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// A schedule block as defined below. Changing this forces a new Pim Active Role Assignment to be created.
	// The schedule details of this role assignment.
	Schedule *ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The scope. Changing this forces a new Pim Active Role Assignment to be created.
	// The scope.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// A ticket block as defined below. Changing this forces a new Pim Active Role Assignment to be created.
	// The ticket details.
	Ticket *TicketObservation `json:"ticket,omitempty" tf:"ticket,omitempty"`
}

type PimActiveRoleAssignmentParameters struct {

	// The justification of the role assignment. Changing this forces a new Pim Active Role Assignment to be created.
	// The justification of the role assignment.
	// +kubebuilder:validation:Optional
	Justification *string `json:"justification,omitempty" tf:"justification,omitempty"`

	// The principal id. Changing this forces a new Pim Active Role Assignment to be created.
	// The principal id.
	// +kubebuilder:validation:Optional
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The role definition id. Changing this forces a new Pim Active Role Assignment to be created.
	// The role definition id.
	// +kubebuilder:validation:Optional
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// A schedule block as defined below. Changing this forces a new Pim Active Role Assignment to be created.
	// The schedule details of this role assignment.
	// +kubebuilder:validation:Optional
	Schedule *ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The scope. Changing this forces a new Pim Active Role Assignment to be created.
	// The scope.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/management/v1beta1.ManagementGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Reference to a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeRef *v1.Reference `json:"scopeRef,omitempty" tf:"-"`

	// Selector for a ManagementGroup in management to populate scope.
	// +kubebuilder:validation:Optional
	ScopeSelector *v1.Selector `json:"scopeSelector,omitempty" tf:"-"`

	// A ticket block as defined below. Changing this forces a new Pim Active Role Assignment to be created.
	// The ticket details.
	// +kubebuilder:validation:Optional
	Ticket *TicketParameters `json:"ticket,omitempty" tf:"ticket,omitempty"`
}

type ScheduleInitParameters struct {

	// A expiration block as defined above.
	Expiration *ExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The start date time of the role assignment. Changing this forces a new Pim Active Role Assignment to be created.
	// The start date time.
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`
}

type ScheduleObservation struct {

	// A expiration block as defined above.
	Expiration *ExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The start date time of the role assignment. Changing this forces a new Pim Active Role Assignment to be created.
	// The start date time.
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`
}

type ScheduleParameters struct {

	// A expiration block as defined above.
	// +kubebuilder:validation:Optional
	Expiration *ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The start date time of the role assignment. Changing this forces a new Pim Active Role Assignment to be created.
	// The start date time.
	// +kubebuilder:validation:Optional
	StartDateTime *string `json:"startDateTime,omitempty" tf:"start_date_time,omitempty"`
}

type TicketInitParameters struct {

	// The ticket number.
	// The ticket number.
	Number *string `json:"number,omitempty" tf:"number,omitempty"`

	// The ticket system.
	// The ticket system.
	System *string `json:"system,omitempty" tf:"system,omitempty"`
}

type TicketObservation struct {

	// The ticket number.
	// The ticket number.
	Number *string `json:"number,omitempty" tf:"number,omitempty"`

	// The ticket system.
	// The ticket system.
	System *string `json:"system,omitempty" tf:"system,omitempty"`
}

type TicketParameters struct {

	// The ticket number.
	// The ticket number.
	// +kubebuilder:validation:Optional
	Number *string `json:"number,omitempty" tf:"number,omitempty"`

	// The ticket system.
	// The ticket system.
	// +kubebuilder:validation:Optional
	System *string `json:"system,omitempty" tf:"system,omitempty"`
}

// PimActiveRoleAssignmentSpec defines the desired state of PimActiveRoleAssignment
type PimActiveRoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PimActiveRoleAssignmentParameters `json:"forProvider"`
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
	InitProvider PimActiveRoleAssignmentInitParameters `json:"initProvider,omitempty"`
}

// PimActiveRoleAssignmentStatus defines the observed state of PimActiveRoleAssignment.
type PimActiveRoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PimActiveRoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PimActiveRoleAssignment is the Schema for the PimActiveRoleAssignments API. Manages a Pim Active Role Assignment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PimActiveRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalId) || (has(self.initProvider) && has(self.initProvider.principalId))",message="spec.forProvider.principalId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleDefinitionId) || (has(self.initProvider) && has(self.initProvider.roleDefinitionId))",message="spec.forProvider.roleDefinitionId is a required parameter"
	Spec   PimActiveRoleAssignmentSpec   `json:"spec"`
	Status PimActiveRoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PimActiveRoleAssignmentList contains a list of PimActiveRoleAssignments
type PimActiveRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PimActiveRoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	PimActiveRoleAssignment_Kind             = "PimActiveRoleAssignment"
	PimActiveRoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PimActiveRoleAssignment_Kind}.String()
	PimActiveRoleAssignment_KindAPIVersion   = PimActiveRoleAssignment_Kind + "." + CRDGroupVersion.String()
	PimActiveRoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(PimActiveRoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&PimActiveRoleAssignment{}, &PimActiveRoleAssignmentList{})
}
