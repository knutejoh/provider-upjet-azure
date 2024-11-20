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

type ConditionAlertContextInitParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionAlertContextObservation struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionAlertContextParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionAlertRuleIDInitParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionAlertRuleIDObservation struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionAlertRuleIDParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionDescriptionInitParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionDescriptionObservation struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionDescriptionParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionMonitorInitParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionMonitorObservation struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionMonitorParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionMonitorServiceInitParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionMonitorServiceObservation struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionMonitorServiceParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionSeverityInitParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionSeverityObservation struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionSeverityParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionTargetResourceTypeInitParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionTargetResourceTypeObservation struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionTargetResourceTypeParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorActionRuleSuppressionConditionInitParameters struct {

	// A alert_context block as defined below.
	AlertContext *ConditionAlertContextInitParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined below.
	AlertRuleID *ConditionAlertRuleIDInitParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A description block as defined below.
	Description *ConditionDescriptionInitParameters `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor block as defined below.
	Monitor *ConditionMonitorInitParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// A monitor_service block as defined below.
	MonitorService *ConditionMonitorServiceInitParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	Severity *ConditionSeverityInitParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// A target_resource_type block as defined below.
	TargetResourceType *ConditionTargetResourceTypeInitParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorActionRuleSuppressionConditionObservation struct {

	// A alert_context block as defined below.
	AlertContext *ConditionAlertContextObservation `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined below.
	AlertRuleID *ConditionAlertRuleIDObservation `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A description block as defined below.
	Description *ConditionDescriptionObservation `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor block as defined below.
	Monitor *ConditionMonitorObservation `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// A monitor_service block as defined below.
	MonitorService *ConditionMonitorServiceObservation `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	Severity *ConditionSeverityObservation `json:"severity,omitempty" tf:"severity,omitempty"`

	// A target_resource_type block as defined below.
	TargetResourceType *ConditionTargetResourceTypeObservation `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorActionRuleSuppressionConditionParameters struct {

	// A alert_context block as defined below.
	// +kubebuilder:validation:Optional
	AlertContext *ConditionAlertContextParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined below.
	// +kubebuilder:validation:Optional
	AlertRuleID *ConditionAlertRuleIDParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A description block as defined below.
	// +kubebuilder:validation:Optional
	Description *ConditionDescriptionParameters `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor block as defined below.
	// +kubebuilder:validation:Optional
	Monitor *ConditionMonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// A monitor_service block as defined below.
	// +kubebuilder:validation:Optional
	MonitorService *ConditionMonitorServiceParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	// +kubebuilder:validation:Optional
	Severity *ConditionSeverityParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// A target_resource_type block as defined below.
	// +kubebuilder:validation:Optional
	TargetResourceType *ConditionTargetResourceTypeParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorActionRuleSuppressionInitParameters struct {

	// A condition block as defined below.
	Condition *MonitorActionRuleSuppressionConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Action Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is the Action Rule enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A scope block as defined below.
	Scope *MonitorActionRuleSuppressionScopeInitParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// A suppression block as defined below.
	Suppression *SuppressionInitParameters `json:"suppression,omitempty" tf:"suppression,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActionRuleSuppressionObservation struct {

	// A condition block as defined below.
	Condition *MonitorActionRuleSuppressionConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Action Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is the Action Rule enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Monitor Action Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A scope block as defined below.
	Scope *MonitorActionRuleSuppressionScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`

	// A suppression block as defined below.
	Suppression *SuppressionObservation `json:"suppression,omitempty" tf:"suppression,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActionRuleSuppressionParameters struct {

	// A condition block as defined below.
	// +kubebuilder:validation:Optional
	Condition *MonitorActionRuleSuppressionConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Action Rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is the Action Rule enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A scope block as defined below.
	// +kubebuilder:validation:Optional
	Scope *MonitorActionRuleSuppressionScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// A suppression block as defined below.
	// +kubebuilder:validation:Optional
	Suppression *SuppressionParameters `json:"suppression,omitempty" tf:"suppression,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActionRuleSuppressionScopeInitParameters struct {

	// A list of resource IDs of the given scope type which will be the target of action rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +listType=set
	ResourceIds []*string `json:"resourceIds,omitempty" tf:"resource_ids,omitempty"`

	// References to ResourceGroup in azure to populate resourceIds.
	// +kubebuilder:validation:Optional
	ResourceIdsRefs []v1.Reference `json:"resourceIdsRefs,omitempty" tf:"-"`

	// Selector for a list of ResourceGroup in azure to populate resourceIds.
	// +kubebuilder:validation:Optional
	ResourceIdsSelector *v1.Selector `json:"resourceIdsSelector,omitempty" tf:"-"`

	// Specifies the type of target scope. Possible values are ResourceGroup and Resource.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MonitorActionRuleSuppressionScopeObservation struct {

	// A list of resource IDs of the given scope type which will be the target of action rule.
	// +listType=set
	ResourceIds []*string `json:"resourceIds,omitempty" tf:"resource_ids,omitempty"`

	// Specifies the type of target scope. Possible values are ResourceGroup and Resource.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MonitorActionRuleSuppressionScopeParameters struct {

	// A list of resource IDs of the given scope type which will be the target of action rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	ResourceIds []*string `json:"resourceIds,omitempty" tf:"resource_ids,omitempty"`

	// References to ResourceGroup in azure to populate resourceIds.
	// +kubebuilder:validation:Optional
	ResourceIdsRefs []v1.Reference `json:"resourceIdsRefs,omitempty" tf:"-"`

	// Selector for a list of ResourceGroup in azure to populate resourceIds.
	// +kubebuilder:validation:Optional
	ResourceIdsSelector *v1.Selector `json:"resourceIdsSelector,omitempty" tf:"-"`

	// Specifies the type of target scope. Possible values are ResourceGroup and Resource.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScheduleInitParameters struct {

	// specifies the recurrence UTC end datetime (Y-m-d'T'H:M:S'Z').
	EndDateUtc *string `json:"endDateUtc,omitempty" tf:"end_date_utc,omitempty"`

	// specifies the list of dayOfMonth to recurrence. Possible values are between 1 - 31. Required if recurrence_type is Monthly.
	// +listType=set
	RecurrenceMonthly []*float64 `json:"recurrenceMonthly,omitempty" tf:"recurrence_monthly,omitempty"`

	// specifies the list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and Saturday.
	// +listType=set
	RecurrenceWeekly []*string `json:"recurrenceWeekly,omitempty" tf:"recurrence_weekly,omitempty"`

	// specifies the recurrence UTC start datetime (Y-m-d'T'H:M:S'Z').
	StartDateUtc *string `json:"startDateUtc,omitempty" tf:"start_date_utc,omitempty"`
}

type ScheduleObservation struct {

	// specifies the recurrence UTC end datetime (Y-m-d'T'H:M:S'Z').
	EndDateUtc *string `json:"endDateUtc,omitempty" tf:"end_date_utc,omitempty"`

	// specifies the list of dayOfMonth to recurrence. Possible values are between 1 - 31. Required if recurrence_type is Monthly.
	// +listType=set
	RecurrenceMonthly []*float64 `json:"recurrenceMonthly,omitempty" tf:"recurrence_monthly,omitempty"`

	// specifies the list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and Saturday.
	// +listType=set
	RecurrenceWeekly []*string `json:"recurrenceWeekly,omitempty" tf:"recurrence_weekly,omitempty"`

	// specifies the recurrence UTC start datetime (Y-m-d'T'H:M:S'Z').
	StartDateUtc *string `json:"startDateUtc,omitempty" tf:"start_date_utc,omitempty"`
}

type ScheduleParameters struct {

	// specifies the recurrence UTC end datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Optional
	EndDateUtc *string `json:"endDateUtc" tf:"end_date_utc,omitempty"`

	// specifies the list of dayOfMonth to recurrence. Possible values are between 1 - 31. Required if recurrence_type is Monthly.
	// +kubebuilder:validation:Optional
	// +listType=set
	RecurrenceMonthly []*float64 `json:"recurrenceMonthly,omitempty" tf:"recurrence_monthly,omitempty"`

	// specifies the list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and Saturday.
	// +kubebuilder:validation:Optional
	// +listType=set
	RecurrenceWeekly []*string `json:"recurrenceWeekly,omitempty" tf:"recurrence_weekly,omitempty"`

	// specifies the recurrence UTC start datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Optional
	StartDateUtc *string `json:"startDateUtc" tf:"start_date_utc,omitempty"`
}

type SuppressionInitParameters struct {

	// Specifies the type of suppression. Possible values are Always, Daily, Monthly, Once, and Weekly.
	RecurrenceType *string `json:"recurrenceType,omitempty" tf:"recurrence_type,omitempty"`

	// A schedule block as defined below. Required if recurrence_type is Daily, Monthly, Once or Weekly.
	Schedule *ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type SuppressionObservation struct {

	// Specifies the type of suppression. Possible values are Always, Daily, Monthly, Once, and Weekly.
	RecurrenceType *string `json:"recurrenceType,omitempty" tf:"recurrence_type,omitempty"`

	// A schedule block as defined below. Required if recurrence_type is Daily, Monthly, Once or Weekly.
	Schedule *ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type SuppressionParameters struct {

	// Specifies the type of suppression. Possible values are Always, Daily, Monthly, Once, and Weekly.
	// +kubebuilder:validation:Optional
	RecurrenceType *string `json:"recurrenceType" tf:"recurrence_type,omitempty"`

	// A schedule block as defined below. Required if recurrence_type is Daily, Monthly, Once or Weekly.
	// +kubebuilder:validation:Optional
	Schedule *ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

// MonitorActionRuleSuppressionSpec defines the desired state of MonitorActionRuleSuppression
type MonitorActionRuleSuppressionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorActionRuleSuppressionParameters `json:"forProvider"`
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
	InitProvider MonitorActionRuleSuppressionInitParameters `json:"initProvider,omitempty"`
}

// MonitorActionRuleSuppressionStatus defines the observed state of MonitorActionRuleSuppression.
type MonitorActionRuleSuppressionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorActionRuleSuppressionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MonitorActionRuleSuppression is the Schema for the MonitorActionRuleSuppressions API. Manages an Monitor Action Rule which type is suppression.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorActionRuleSuppression struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.suppression) || (has(self.initProvider) && has(self.initProvider.suppression))",message="spec.forProvider.suppression is a required parameter"
	Spec   MonitorActionRuleSuppressionSpec   `json:"spec"`
	Status MonitorActionRuleSuppressionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionRuleSuppressionList contains a list of MonitorActionRuleSuppressions
type MonitorActionRuleSuppressionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorActionRuleSuppression `json:"items"`
}

// Repository type metadata.
var (
	MonitorActionRuleSuppression_Kind             = "MonitorActionRuleSuppression"
	MonitorActionRuleSuppression_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorActionRuleSuppression_Kind}.String()
	MonitorActionRuleSuppression_KindAPIVersion   = MonitorActionRuleSuppression_Kind + "." + CRDGroupVersion.String()
	MonitorActionRuleSuppression_GroupVersionKind = CRDGroupVersion.WithKind(MonitorActionRuleSuppression_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorActionRuleSuppression{}, &MonitorActionRuleSuppressionList{})
}
