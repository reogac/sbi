/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EcsServerAddr struct {
	EcsFqdnList      []string `json:"ecsFqdnList,omitempty"`
	EcsIpAddressList []IpAddr `json:"ecsIpAddressList,omitempty"`
	EcsUriList       []string `json:"ecsUriList,omitempty"`
	EcsProviderId    string   `json:"ecsProviderId,omitempty"`
}
