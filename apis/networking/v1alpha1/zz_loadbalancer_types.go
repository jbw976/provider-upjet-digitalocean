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

type ForwardingRuleInitParameters struct {

	// Deprecated The ID of the TLS certificate to be used for SSL termination.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The unique name of the TLS certificate to be used for SSL termination.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// An integer representing the port on which the Load Balancer instance will listen.
	EntryPort *float64 `json:"entryPort,omitempty" tf:"entry_port,omitempty"`

	// The protocol used for traffic to the Load Balancer. The possible values are: http, https, http2, http3, tcp, or udp.
	EntryProtocol *string `json:"entryProtocol,omitempty" tf:"entry_protocol,omitempty"`

	// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is false.
	TLSPassthrough *bool `json:"tlsPassthrough,omitempty" tf:"tls_passthrough,omitempty"`

	// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
	TargetPort *float64 `json:"targetPort,omitempty" tf:"target_port,omitempty"`

	// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: http, https, http2, tcp, or udp.
	TargetProtocol *string `json:"targetProtocol,omitempty" tf:"target_protocol,omitempty"`
}

type ForwardingRuleObservation struct {

	// Deprecated The ID of the TLS certificate to be used for SSL termination.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The unique name of the TLS certificate to be used for SSL termination.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// An integer representing the port on which the Load Balancer instance will listen.
	EntryPort *float64 `json:"entryPort,omitempty" tf:"entry_port,omitempty"`

	// The protocol used for traffic to the Load Balancer. The possible values are: http, https, http2, http3, tcp, or udp.
	EntryProtocol *string `json:"entryProtocol,omitempty" tf:"entry_protocol,omitempty"`

	// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is false.
	TLSPassthrough *bool `json:"tlsPassthrough,omitempty" tf:"tls_passthrough,omitempty"`

	// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
	TargetPort *float64 `json:"targetPort,omitempty" tf:"target_port,omitempty"`

	// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: http, https, http2, tcp, or udp.
	TargetProtocol *string `json:"targetProtocol,omitempty" tf:"target_protocol,omitempty"`
}

type ForwardingRuleParameters struct {

	// Deprecated The ID of the TLS certificate to be used for SSL termination.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// The unique name of the TLS certificate to be used for SSL termination.
	// +kubebuilder:validation:Optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// An integer representing the port on which the Load Balancer instance will listen.
	// +kubebuilder:validation:Optional
	EntryPort *float64 `json:"entryPort" tf:"entry_port,omitempty"`

	// The protocol used for traffic to the Load Balancer. The possible values are: http, https, http2, http3, tcp, or udp.
	// +kubebuilder:validation:Optional
	EntryProtocol *string `json:"entryProtocol" tf:"entry_protocol,omitempty"`

	// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is false.
	// +kubebuilder:validation:Optional
	TLSPassthrough *bool `json:"tlsPassthrough,omitempty" tf:"tls_passthrough,omitempty"`

	// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
	// +kubebuilder:validation:Optional
	TargetPort *float64 `json:"targetPort" tf:"target_port,omitempty"`

	// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: http, https, http2, tcp, or udp.
	// +kubebuilder:validation:Optional
	TargetProtocol *string `json:"targetProtocol" tf:"target_protocol,omitempty"`
}

type HealthcheckInitParameters struct {

	// The number of seconds between two consecutive health checks. If not specified, the default value is 10.
	CheckIntervalSeconds *float64 `json:"checkIntervalSeconds,omitempty" tf:"check_interval_seconds,omitempty"`

	// The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is 5.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The path on the backend Droplets to which the Load Balancer instance will send a request.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// An integer representing the port on the backend Droplets on which the health check will attempt a connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol used for health checks sent to the backend Droplets. The possible values are http, https or tcp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is 5.
	ResponseTimeoutSeconds *float64 `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds,omitempty"`

	// The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is 3.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthcheckObservation struct {

	// The number of seconds between two consecutive health checks. If not specified, the default value is 10.
	CheckIntervalSeconds *float64 `json:"checkIntervalSeconds,omitempty" tf:"check_interval_seconds,omitempty"`

	// The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is 5.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The path on the backend Droplets to which the Load Balancer instance will send a request.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// An integer representing the port on the backend Droplets on which the health check will attempt a connection.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol used for health checks sent to the backend Droplets. The possible values are http, https or tcp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is 5.
	ResponseTimeoutSeconds *float64 `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds,omitempty"`

	// The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is 3.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthcheckParameters struct {

	// The number of seconds between two consecutive health checks. If not specified, the default value is 10.
	// +kubebuilder:validation:Optional
	CheckIntervalSeconds *float64 `json:"checkIntervalSeconds,omitempty" tf:"check_interval_seconds,omitempty"`

	// The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is 5.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The path on the backend Droplets to which the Load Balancer instance will send a request.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// An integer representing the port on the backend Droplets on which the health check will attempt a connection.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The protocol used for health checks sent to the backend Droplets. The possible values are http, https or tcp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is 5.
	// +kubebuilder:validation:Optional
	ResponseTimeoutSeconds *float64 `json:"responseTimeoutSeconds,omitempty" tf:"response_timeout_seconds,omitempty"`

	// The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is 3.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LoadbalancerFirewallInitParameters struct {

	// A list of strings describing allow rules. Must be colon delimited strings of the form {type}:{source}
	// the rules for ALLOWING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
	Allow []*string `json:"allow,omitempty" tf:"allow,omitempty"`

	// A list of strings describing deny rules. Must be colon delimited strings of the form {type}:{source}
	// the rules for DENYING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
	Deny []*string `json:"deny,omitempty" tf:"deny,omitempty"`
}

type LoadbalancerFirewallObservation struct {

	// A list of strings describing allow rules. Must be colon delimited strings of the form {type}:{source}
	// the rules for ALLOWING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
	Allow []*string `json:"allow,omitempty" tf:"allow,omitempty"`

	// A list of strings describing deny rules. Must be colon delimited strings of the form {type}:{source}
	// the rules for DENYING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
	Deny []*string `json:"deny,omitempty" tf:"deny,omitempty"`
}

type LoadbalancerFirewallParameters struct {

	// A list of strings describing allow rules. Must be colon delimited strings of the form {type}:{source}
	// the rules for ALLOWING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
	// +kubebuilder:validation:Optional
	Allow []*string `json:"allow,omitempty" tf:"allow,omitempty"`

	// A list of strings describing deny rules. Must be colon delimited strings of the form {type}:{source}
	// the rules for DENYING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16')
	// +kubebuilder:validation:Optional
	Deny []*string `json:"deny,omitempty" tf:"deny,omitempty"`
}

type LoadbalancerInitParameters struct {

	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either round_robin
	// or least_connections. The default value is round_robin.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer. Default value is false.
	DisableLetsEncryptDNSRecords *bool `json:"disableLetsEncryptDnsRecords,omitempty" tf:"disable_lets_encrypt_dns_records,omitempty"`

	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	DropletTag *string `json:"dropletTag,omitempty" tf:"droplet_tag,omitempty"`

	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is false.
	EnableBackendKeepalive *bool `json:"enableBackendKeepalive,omitempty" tf:"enable_backend_keepalive,omitempty"`

	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is false.
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`

	// A block containing rules for allowing/denying traffic to the Load Balancer. The firewall block is documented below. Only 1 firewall is allowed.
	Firewall []LoadbalancerFirewallInitParameters `json:"firewall,omitempty" tf:"firewall,omitempty"`

	// A list of forwarding_rule to be assigned to the
	// Load Balancer. The forwarding_rule block is documented below.
	ForwardingRule []ForwardingRuleInitParameters `json:"forwardingRule,omitempty" tf:"forwarding_rule,omitempty"`

	// Specifies the idle timeout for HTTPS connections on the load balancer in seconds.
	HTTPIdleTimeoutSeconds *float64 `json:"httpIdleTimeoutSeconds,omitempty" tf:"http_idle_timeout_seconds,omitempty"`

	// A healthcheck block to be assigned to the
	// Load Balancer. The healthcheck block is documented below. Only 1 healthcheck is allowed.
	Healthcheck []HealthcheckInitParameters `json:"healthcheck,omitempty" tf:"healthcheck,omitempty"`

	// The Load Balancer name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is false.
	RedirectHTTPToHTTPS *bool `json:"redirectHttpToHttps,omitempty" tf:"redirect_http_to_https,omitempty"`

	// The region to start in
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the Load Balancer. It must be either lb-small, lb-medium, or lb-large. Defaults to lb-small. Only one of size or size_unit may be provided.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// The size of the Load Balancer. It must be in the range (1, 100). Defaults to 1. Only one of size or size_unit may be provided.
	SizeUnit *float64 `json:"sizeUnit,omitempty" tf:"size_unit,omitempty"`

	// A sticky_sessions block to be assigned to the
	// Load Balancer. The sticky_sessions block is documented below. Only 1 sticky_sessions block is allowed.
	StickySessions []StickySessionsInitParameters `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`

	// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are cookies or none. If not specified, the default value is none.
	// the type of the load balancer (GLOBAL or REGIONAL)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LoadbalancerObservation struct {

	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either round_robin
	// or least_connections. The default value is round_robin.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer. Default value is false.
	DisableLetsEncryptDNSRecords *bool `json:"disableLetsEncryptDnsRecords,omitempty" tf:"disable_lets_encrypt_dns_records,omitempty"`

	// A list of the IDs of each droplet to be attached to the Load Balancer.
	DropletIds []*float64 `json:"dropletIds,omitempty" tf:"droplet_ids,omitempty"`

	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	DropletTag *string `json:"dropletTag,omitempty" tf:"droplet_tag,omitempty"`

	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is false.
	EnableBackendKeepalive *bool `json:"enableBackendKeepalive,omitempty" tf:"enable_backend_keepalive,omitempty"`

	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is false.
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`

	// A block containing rules for allowing/denying traffic to the Load Balancer. The firewall block is documented below. Only 1 firewall is allowed.
	Firewall []LoadbalancerFirewallObservation `json:"firewall,omitempty" tf:"firewall,omitempty"`

	// A list of forwarding_rule to be assigned to the
	// Load Balancer. The forwarding_rule block is documented below.
	ForwardingRule []ForwardingRuleObservation `json:"forwardingRule,omitempty" tf:"forwarding_rule,omitempty"`

	// Specifies the idle timeout for HTTPS connections on the load balancer in seconds.
	HTTPIdleTimeoutSeconds *float64 `json:"httpIdleTimeoutSeconds,omitempty" tf:"http_idle_timeout_seconds,omitempty"`

	// A healthcheck block to be assigned to the
	// Load Balancer. The healthcheck block is documented below. Only 1 healthcheck is allowed.
	Healthcheck []HealthcheckObservation `json:"healthcheck,omitempty" tf:"healthcheck,omitempty"`

	// The ID of the Load Balancer
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ip of the Load Balancer
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The Load Balancer name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is false.
	RedirectHTTPToHTTPS *bool `json:"redirectHttpToHttps,omitempty" tf:"redirect_http_to_https,omitempty"`

	// The region to start in
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the Load Balancer. It must be either lb-small, lb-medium, or lb-large. Defaults to lb-small. Only one of size or size_unit may be provided.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// The size of the Load Balancer. It must be in the range (1, 100). Defaults to 1. Only one of size or size_unit may be provided.
	SizeUnit *float64 `json:"sizeUnit,omitempty" tf:"size_unit,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A sticky_sessions block to be assigned to the
	// Load Balancer. The sticky_sessions block is documented below. Only 1 sticky_sessions block is allowed.
	StickySessions []StickySessionsObservation `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`

	// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are cookies or none. If not specified, the default value is none.
	// the type of the load balancer (GLOBAL or REGIONAL)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The uniform resource name for the Load Balancer
	// the uniform resource name for the load balancer
	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`

	// The ID of the VPC where the load balancer will be located.
	VPCUUID *string `json:"vpcUuid,omitempty" tf:"vpc_uuid,omitempty"`
}

type LoadbalancerParameters struct {

	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either round_robin
	// or least_connections. The default value is round_robin.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer. Default value is false.
	// +kubebuilder:validation:Optional
	DisableLetsEncryptDNSRecords *bool `json:"disableLetsEncryptDnsRecords,omitempty" tf:"disable_lets_encrypt_dns_records,omitempty"`

	// A list of the IDs of each droplet to be attached to the Load Balancer.
	// +crossplane:generate:reference:type=github.com/straw-hat-team/provider-digitalocean/apis/compute/v1alpha1.Droplet
	// +kubebuilder:validation:Optional
	DropletIds []*float64 `json:"dropletIds,omitempty" tf:"droplet_ids,omitempty"`

	// References to Droplet in compute to populate dropletIds.
	// +kubebuilder:validation:Optional
	DropletIdsRefs []v1.Reference `json:"dropletIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Droplet in compute to populate dropletIds.
	// +kubebuilder:validation:Optional
	DropletIdsSelector *v1.Selector `json:"dropletIdsSelector,omitempty" tf:"-"`

	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	// +kubebuilder:validation:Optional
	DropletTag *string `json:"dropletTag,omitempty" tf:"droplet_tag,omitempty"`

	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is false.
	// +kubebuilder:validation:Optional
	EnableBackendKeepalive *bool `json:"enableBackendKeepalive,omitempty" tf:"enable_backend_keepalive,omitempty"`

	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is false.
	// +kubebuilder:validation:Optional
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`

	// A block containing rules for allowing/denying traffic to the Load Balancer. The firewall block is documented below. Only 1 firewall is allowed.
	// +kubebuilder:validation:Optional
	Firewall []LoadbalancerFirewallParameters `json:"firewall,omitempty" tf:"firewall,omitempty"`

	// A list of forwarding_rule to be assigned to the
	// Load Balancer. The forwarding_rule block is documented below.
	// +kubebuilder:validation:Optional
	ForwardingRule []ForwardingRuleParameters `json:"forwardingRule,omitempty" tf:"forwarding_rule,omitempty"`

	// Specifies the idle timeout for HTTPS connections on the load balancer in seconds.
	// +kubebuilder:validation:Optional
	HTTPIdleTimeoutSeconds *float64 `json:"httpIdleTimeoutSeconds,omitempty" tf:"http_idle_timeout_seconds,omitempty"`

	// A healthcheck block to be assigned to the
	// Load Balancer. The healthcheck block is documented below. Only 1 healthcheck is allowed.
	// +kubebuilder:validation:Optional
	Healthcheck []HealthcheckParameters `json:"healthcheck,omitempty" tf:"healthcheck,omitempty"`

	// The Load Balancer name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is false.
	// +kubebuilder:validation:Optional
	RedirectHTTPToHTTPS *bool `json:"redirectHttpToHttps,omitempty" tf:"redirect_http_to_https,omitempty"`

	// The region to start in
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the Load Balancer. It must be either lb-small, lb-medium, or lb-large. Defaults to lb-small. Only one of size or size_unit may be provided.
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// The size of the Load Balancer. It must be in the range (1, 100). Defaults to 1. Only one of size or size_unit may be provided.
	// +kubebuilder:validation:Optional
	SizeUnit *float64 `json:"sizeUnit,omitempty" tf:"size_unit,omitempty"`

	// A sticky_sessions block to be assigned to the
	// Load Balancer. The sticky_sessions block is documented below. Only 1 sticky_sessions block is allowed.
	// +kubebuilder:validation:Optional
	StickySessions []StickySessionsParameters `json:"stickySessions,omitempty" tf:"sticky_sessions,omitempty"`

	// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are cookies or none. If not specified, the default value is none.
	// the type of the load balancer (GLOBAL or REGIONAL)
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the VPC where the load balancer will be located.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VPCUUID *string `json:"vpcUuid,omitempty" tf:"vpc_uuid,omitempty"`

	// Reference to a VPC to populate vpcUuid.
	// +kubebuilder:validation:Optional
	VPCUUIDRef *v1.Reference `json:"vpcUuidRef,omitempty" tf:"-"`

	// Selector for a VPC to populate vpcUuid.
	// +kubebuilder:validation:Optional
	VPCUUIDSelector *v1.Selector `json:"vpcUuidSelector,omitempty" tf:"-"`
}

type StickySessionsInitParameters struct {

	// The name to be used for the cookie sent to the client. This attribute is required when using cookies for the sticky sessions type.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using cookies for the sticky sessions type.
	CookieTTLSeconds *float64 `json:"cookieTtlSeconds,omitempty" tf:"cookie_ttl_seconds,omitempty"`

	// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are cookies or none. If not specified, the default value is none.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StickySessionsObservation struct {

	// The name to be used for the cookie sent to the client. This attribute is required when using cookies for the sticky sessions type.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using cookies for the sticky sessions type.
	CookieTTLSeconds *float64 `json:"cookieTtlSeconds,omitempty" tf:"cookie_ttl_seconds,omitempty"`

	// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are cookies or none. If not specified, the default value is none.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StickySessionsParameters struct {

	// The name to be used for the cookie sent to the client. This attribute is required when using cookies for the sticky sessions type.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using cookies for the sticky sessions type.
	// +kubebuilder:validation:Optional
	CookieTTLSeconds *float64 `json:"cookieTtlSeconds,omitempty" tf:"cookie_ttl_seconds,omitempty"`

	// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are cookies or none. If not specified, the default value is none.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// LoadbalancerSpec defines the desired state of Loadbalancer
type LoadbalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadbalancerParameters `json:"forProvider"`
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
	InitProvider LoadbalancerInitParameters `json:"initProvider,omitempty"`
}

// LoadbalancerStatus defines the observed state of Loadbalancer.
type LoadbalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadbalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Loadbalancer is the Schema for the Loadbalancers API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,do}
type Loadbalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.forwardingRule) || (has(self.initProvider) && has(self.initProvider.forwardingRule))",message="spec.forProvider.forwardingRule is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   LoadbalancerSpec   `json:"spec"`
	Status LoadbalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadbalancerList contains a list of Loadbalancers
type LoadbalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Loadbalancer `json:"items"`
}

// Repository type metadata.
var (
	Loadbalancer_Kind             = "Loadbalancer"
	Loadbalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Loadbalancer_Kind}.String()
	Loadbalancer_KindAPIVersion   = Loadbalancer_Kind + "." + CRDGroupVersion.String()
	Loadbalancer_GroupVersionKind = CRDGroupVersion.WithKind(Loadbalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&Loadbalancer{}, &LoadbalancerList{})
}
