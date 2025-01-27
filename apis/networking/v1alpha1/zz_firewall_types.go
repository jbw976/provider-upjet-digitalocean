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

type FirewallInitParameters struct {

	// The inbound access rule block for the Firewall.
	// The inbound_rule block is documented below.
	InboundRule []InboundRuleInitParameters `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`

	// The Firewall name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The outbound access rule block for the Firewall.
	// The outbound_rule block is documented below.
	OutboundRule []OutboundRuleInitParameters `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`

	// The names of the Tags assigned to the Firewall.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FirewallObservation struct {

	// A time value given in ISO8601 combined date and time format
	// that represents when the Firewall was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds []*float64 `json:"dropletIds,omitempty" tf:"droplet_ids,omitempty"`

	// A unique ID that can be used to identify and reference a Firewall.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The inbound access rule block for the Firewall.
	// The inbound_rule block is documented below.
	InboundRule []InboundRuleObservation `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`

	// The Firewall name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The outbound access rule block for the Firewall.
	// The outbound_rule block is documented below.
	OutboundRule []OutboundRuleObservation `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`

	// An list of object containing the fields, "droplet_id",
	// "removing", and "status".  It is provided to detail exactly which Droplets
	// are having their security policies updated.  When empty, all changes
	// have been successfully applied.
	PendingChanges []PendingChangesObservation `json:"pendingChanges,omitempty" tf:"pending_changes,omitempty"`

	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The names of the Tags assigned to the Firewall.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FirewallParameters struct {

	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-digitalocean/apis/droplet/v1alpha1.Droplet
	// +kubebuilder:validation:Optional
	DropletIds []*float64 `json:"dropletIds,omitempty" tf:"droplet_ids,omitempty"`

	// References to Droplet in droplet to populate dropletIds.
	// +kubebuilder:validation:Optional
	DropletIdsRefs []v1.Reference `json:"dropletIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Droplet in droplet to populate dropletIds.
	// +kubebuilder:validation:Optional
	DropletIdsSelector *v1.Selector `json:"dropletIdsSelector,omitempty" tf:"-"`

	// The inbound access rule block for the Firewall.
	// The inbound_rule block is documented below.
	// +kubebuilder:validation:Optional
	InboundRule []InboundRuleParameters `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`

	// The Firewall name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The outbound access rule block for the Firewall.
	// The outbound_rule block is documented below.
	// +kubebuilder:validation:Optional
	OutboundRule []OutboundRuleParameters `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`

	// The names of the Tags assigned to the Firewall.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InboundRuleInitParameters struct {

	// The ports on which traffic will be allowed
	// specified as a string containing a single port, a range (e.g. "8000-9000"),
	// or "1-65535" to open all ports for a protocol. Required for when protocol is
	// tcp or udp.
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The type of traffic to be allowed.
	// This may be one of "tcp", "udp", or "icmp".
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// An array of strings containing the IPv4
	// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
	// inbound traffic will be accepted.
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// An array containing the IDs of
	// the Droplets from which the inbound traffic will be accepted.
	SourceDropletIds []*float64 `json:"sourceDropletIds,omitempty" tf:"source_droplet_ids,omitempty"`

	// An array containing the IDs of
	// the Kubernetes clusters from which the inbound traffic will be accepted.
	SourceKubernetesIds []*string `json:"sourceKubernetesIds,omitempty" tf:"source_kubernetes_ids,omitempty"`

	// An array containing the IDs
	// of the Load Balancers from which the inbound traffic will be accepted.
	SourceLoadBalancerUids []*string `json:"sourceLoadBalancerUids,omitempty" tf:"source_load_balancer_uids,omitempty"`

	// An array containing the names of Tags
	// corresponding to groups of Droplets from which the inbound traffic
	// will be accepted.
	SourceTags []*string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`
}

type InboundRuleObservation struct {

	// The ports on which traffic will be allowed
	// specified as a string containing a single port, a range (e.g. "8000-9000"),
	// or "1-65535" to open all ports for a protocol. Required for when protocol is
	// tcp or udp.
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The type of traffic to be allowed.
	// This may be one of "tcp", "udp", or "icmp".
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// An array of strings containing the IPv4
	// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
	// inbound traffic will be accepted.
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// An array containing the IDs of
	// the Droplets from which the inbound traffic will be accepted.
	SourceDropletIds []*float64 `json:"sourceDropletIds,omitempty" tf:"source_droplet_ids,omitempty"`

	// An array containing the IDs of
	// the Kubernetes clusters from which the inbound traffic will be accepted.
	SourceKubernetesIds []*string `json:"sourceKubernetesIds,omitempty" tf:"source_kubernetes_ids,omitempty"`

	// An array containing the IDs
	// of the Load Balancers from which the inbound traffic will be accepted.
	SourceLoadBalancerUids []*string `json:"sourceLoadBalancerUids,omitempty" tf:"source_load_balancer_uids,omitempty"`

	// An array containing the names of Tags
	// corresponding to groups of Droplets from which the inbound traffic
	// will be accepted.
	SourceTags []*string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`
}

type InboundRuleParameters struct {

	// The ports on which traffic will be allowed
	// specified as a string containing a single port, a range (e.g. "8000-9000"),
	// or "1-65535" to open all ports for a protocol. Required for when protocol is
	// tcp or udp.
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The type of traffic to be allowed.
	// This may be one of "tcp", "udp", or "icmp".
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// An array of strings containing the IPv4
	// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
	// inbound traffic will be accepted.
	// +kubebuilder:validation:Optional
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// An array containing the IDs of
	// the Droplets from which the inbound traffic will be accepted.
	// +kubebuilder:validation:Optional
	SourceDropletIds []*float64 `json:"sourceDropletIds,omitempty" tf:"source_droplet_ids,omitempty"`

	// An array containing the IDs of
	// the Kubernetes clusters from which the inbound traffic will be accepted.
	// +kubebuilder:validation:Optional
	SourceKubernetesIds []*string `json:"sourceKubernetesIds,omitempty" tf:"source_kubernetes_ids,omitempty"`

	// An array containing the IDs
	// of the Load Balancers from which the inbound traffic will be accepted.
	// +kubebuilder:validation:Optional
	SourceLoadBalancerUids []*string `json:"sourceLoadBalancerUids,omitempty" tf:"source_load_balancer_uids,omitempty"`

	// An array containing the names of Tags
	// corresponding to groups of Droplets from which the inbound traffic
	// will be accepted.
	// +kubebuilder:validation:Optional
	SourceTags []*string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`
}

type OutboundRuleInitParameters struct {

	// An array of strings containing the IPv4
	// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
	// outbound traffic will be allowed.
	DestinationAddresses []*string `json:"destinationAddresses,omitempty" tf:"destination_addresses,omitempty"`

	// An array containing the IDs of
	// the Droplets to which the outbound traffic will be allowed.
	DestinationDropletIds []*float64 `json:"destinationDropletIds,omitempty" tf:"destination_droplet_ids,omitempty"`

	// An array containing the IDs of
	// the Kubernetes clusters to which the outbound traffic will be allowed.
	DestinationKubernetesIds []*string `json:"destinationKubernetesIds,omitempty" tf:"destination_kubernetes_ids,omitempty"`

	// An array containing the IDs
	// of the Load Balancers to which the outbound traffic will be allowed.
	DestinationLoadBalancerUids []*string `json:"destinationLoadBalancerUids,omitempty" tf:"destination_load_balancer_uids,omitempty"`

	// An array containing the names of Tags
	// corresponding to groups of Droplets to which the outbound traffic will
	// be allowed.
	DestinationTags []*string `json:"destinationTags,omitempty" tf:"destination_tags,omitempty"`

	// The ports on which traffic will be allowed
	// specified as a string containing a single port, a range (e.g. "8000-9000"),
	// or "1-65535" to open all ports for a protocol. Required for when protocol is
	// tcp or udp.
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The type of traffic to be allowed.
	// This may be one of "tcp", "udp", or "icmp".
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OutboundRuleObservation struct {

	// An array of strings containing the IPv4
	// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
	// outbound traffic will be allowed.
	DestinationAddresses []*string `json:"destinationAddresses,omitempty" tf:"destination_addresses,omitempty"`

	// An array containing the IDs of
	// the Droplets to which the outbound traffic will be allowed.
	DestinationDropletIds []*float64 `json:"destinationDropletIds,omitempty" tf:"destination_droplet_ids,omitempty"`

	// An array containing the IDs of
	// the Kubernetes clusters to which the outbound traffic will be allowed.
	DestinationKubernetesIds []*string `json:"destinationKubernetesIds,omitempty" tf:"destination_kubernetes_ids,omitempty"`

	// An array containing the IDs
	// of the Load Balancers to which the outbound traffic will be allowed.
	DestinationLoadBalancerUids []*string `json:"destinationLoadBalancerUids,omitempty" tf:"destination_load_balancer_uids,omitempty"`

	// An array containing the names of Tags
	// corresponding to groups of Droplets to which the outbound traffic will
	// be allowed.
	DestinationTags []*string `json:"destinationTags,omitempty" tf:"destination_tags,omitempty"`

	// The ports on which traffic will be allowed
	// specified as a string containing a single port, a range (e.g. "8000-9000"),
	// or "1-65535" to open all ports for a protocol. Required for when protocol is
	// tcp or udp.
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The type of traffic to be allowed.
	// This may be one of "tcp", "udp", or "icmp".
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OutboundRuleParameters struct {

	// An array of strings containing the IPv4
	// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
	// outbound traffic will be allowed.
	// +kubebuilder:validation:Optional
	DestinationAddresses []*string `json:"destinationAddresses,omitempty" tf:"destination_addresses,omitempty"`

	// An array containing the IDs of
	// the Droplets to which the outbound traffic will be allowed.
	// +kubebuilder:validation:Optional
	DestinationDropletIds []*float64 `json:"destinationDropletIds,omitempty" tf:"destination_droplet_ids,omitempty"`

	// An array containing the IDs of
	// the Kubernetes clusters to which the outbound traffic will be allowed.
	// +kubebuilder:validation:Optional
	DestinationKubernetesIds []*string `json:"destinationKubernetesIds,omitempty" tf:"destination_kubernetes_ids,omitempty"`

	// An array containing the IDs
	// of the Load Balancers to which the outbound traffic will be allowed.
	// +kubebuilder:validation:Optional
	DestinationLoadBalancerUids []*string `json:"destinationLoadBalancerUids,omitempty" tf:"destination_load_balancer_uids,omitempty"`

	// An array containing the names of Tags
	// corresponding to groups of Droplets to which the outbound traffic will
	// be allowed.
	// +kubebuilder:validation:Optional
	DestinationTags []*string `json:"destinationTags,omitempty" tf:"destination_tags,omitempty"`

	// The ports on which traffic will be allowed
	// specified as a string containing a single port, a range (e.g. "8000-9000"),
	// or "1-65535" to open all ports for a protocol. Required for when protocol is
	// tcp or udp.
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// The type of traffic to be allowed.
	// This may be one of "tcp", "udp", or "icmp".
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type PendingChangesInitParameters struct {
}

type PendingChangesObservation struct {

	// A unique ID that can be used to identify and reference a Firewall.
	DropletID *float64 `json:"dropletId,omitempty" tf:"droplet_id,omitempty"`

	Removing *bool `json:"removing,omitempty" tf:"removing,omitempty"`

	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type PendingChangesParameters struct {
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
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
	InitProvider FirewallInitParameters `json:"initProvider,omitempty"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FirewallSpec   `json:"spec"`
	Status FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
