/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServerAddressingInfo struct {
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty"`
	FqdnList      []string `json:"fqdnList,omitempty"`
}
