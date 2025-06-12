/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsRouterInfo struct {
	RouterIpv4      string                      `json:"routerIpv4,omitempty"`
	RouterIpv6      string                      `json:"routerIpv6,omitempty"`
	RouterFqdn      string                      `json:"routerFqdn,omitempty"`
	NfInstanceId    string                      `json:"nfInstanceId,omitempty"`
	DiameterAddress *NetworkNodeDiameterAddress `json:"diameterAddress,omitempty"`
	MapAddress      string                      `json:"mapAddress,omitempty"`
}
