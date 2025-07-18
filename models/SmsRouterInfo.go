/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:43 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsRouterInfo struct {
	RouterFqdn      string                      `json:"routerFqdn,omitempty"`
	NfInstanceId    string                      `json:"nfInstanceId,omitempty"`
	DiameterAddress *NetworkNodeDiameterAddress `json:"diameterAddress,omitempty"`
	MapAddress      string                      `json:"mapAddress,omitempty"`
	RouterIpv4      string                      `json:"routerIpv4,omitempty"`
	RouterIpv6      string                      `json:"routerIpv6,omitempty"`
}
