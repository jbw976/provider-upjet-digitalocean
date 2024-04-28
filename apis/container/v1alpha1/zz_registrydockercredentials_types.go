// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegistryDockerCredentialsInitParameters struct {

	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	ExpirySeconds *float64 `json:"expirySeconds,omitempty" tf:"expiry_seconds,omitempty"`

	// The name of the container registry.
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Allow for write access to the container registry. Defaults to false.
	Write *bool `json:"write,omitempty" tf:"write,omitempty"`
}

type RegistryDockerCredentialsObservation struct {

	// : The date and time the registry access token will expire.
	CredentialExpirationTime *string `json:"credentialExpirationTime,omitempty" tf:"credential_expiration_time,omitempty"`

	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	ExpirySeconds *float64 `json:"expirySeconds,omitempty" tf:"expiry_seconds,omitempty"`

	// : The ID of the tag. This is the same as the name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the container registry.
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Allow for write access to the container registry. Defaults to false.
	Write *bool `json:"write,omitempty" tf:"write,omitempty"`
}

type RegistryDockerCredentialsParameters struct {

	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	// +kubebuilder:validation:Optional
	ExpirySeconds *float64 `json:"expirySeconds,omitempty" tf:"expiry_seconds,omitempty"`

	// The name of the container registry.
	// +kubebuilder:validation:Optional
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Allow for write access to the container registry. Defaults to false.
	// +kubebuilder:validation:Optional
	Write *bool `json:"write,omitempty" tf:"write,omitempty"`
}

// RegistryDockerCredentialsSpec defines the desired state of RegistryDockerCredentials
type RegistryDockerCredentialsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryDockerCredentialsParameters `json:"forProvider"`
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
	InitProvider RegistryDockerCredentialsInitParameters `json:"initProvider,omitempty"`
}

// RegistryDockerCredentialsStatus defines the observed state of RegistryDockerCredentials.
type RegistryDockerCredentialsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryDockerCredentialsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryDockerCredentials is the Schema for the RegistryDockerCredentialss API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type RegistryDockerCredentials struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.registryName) || (has(self.initProvider) && has(self.initProvider.registryName))",message="spec.forProvider.registryName is a required parameter"
	Spec   RegistryDockerCredentialsSpec   `json:"spec"`
	Status RegistryDockerCredentialsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryDockerCredentialsList contains a list of RegistryDockerCredentialss
type RegistryDockerCredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryDockerCredentials `json:"items"`
}

// Repository type metadata.
var (
	RegistryDockerCredentials_Kind             = "RegistryDockerCredentials"
	RegistryDockerCredentials_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistryDockerCredentials_Kind}.String()
	RegistryDockerCredentials_KindAPIVersion   = RegistryDockerCredentials_Kind + "." + CRDGroupVersion.String()
	RegistryDockerCredentials_GroupVersionKind = CRDGroupVersion.WithKind(RegistryDockerCredentials_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistryDockerCredentials{}, &RegistryDockerCredentialsList{})
}