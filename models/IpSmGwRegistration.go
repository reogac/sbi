/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IpSmGwRegistration struct {
	IpSmGwDiameterAddress *NetworkNodeDiameterAddress `json:"ipSmGwDiameterAddress,omitempty"`
	IpsmgwIpv6            string                      `json:"ipsmgwIpv6,omitempty"`
	IpsmgwFqdn            string                      `json:"ipsmgwFqdn,omitempty"`
	UnriIndicator         *bool                       `json:"unriIndicator,omitempty"`
	IpSmGwSbiSupInd       *bool                       `json:"ipSmGwSbiSupInd,omitempty"`
	IpSmGwMapAddress      string                      `json:"ipSmGwMapAddress,omitempty"`
	IpsmgwIpv4            string                      `json:"ipsmgwIpv4,omitempty"`
	NfInstanceId          string                      `json:"nfInstanceId,omitempty"`
	ResetIds              []string                    `json:"resetIds,omitempty"`
}
