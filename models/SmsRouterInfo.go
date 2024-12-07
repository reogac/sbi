/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:27 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsRouterInfo struct {
	DiameterAddress *NetworkNodeDiameterAddress `json:"diameterAddress,omitempty"`
	MapAddress      string                      `json:"mapAddress,omitempty"`
	RouterIpv4      string                      `json:"routerIpv4,omitempty"`
	RouterIpv6      string                      `json:"routerIpv6,omitempty"`
	RouterFqdn      string                      `json:"routerFqdn,omitempty"`
	NfInstanceId    string                      `json:"nfInstanceId,omitempty"`
}
