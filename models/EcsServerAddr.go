/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EcsServerAddr struct {
	EcsUriList       []string `json:"ecsUriList,omitempty"`
	EcsProviderId    string   `json:"ecsProviderId,omitempty"`
	EcsFqdnList      []string `json:"ecsFqdnList,omitempty"`
	EcsIpAddressList []IpAddr `json:"ecsIpAddressList,omitempty"`
}
