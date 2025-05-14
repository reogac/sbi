/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type WAgfInfo struct {
	Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses,omitempty"`
	Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses,omitempty"`
	EndpointFqdn          string   `json:"endpointFqdn,omitempty"`
}
