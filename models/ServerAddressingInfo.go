/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServerAddressingInfo struct {
	FqdnList      []string `json:"fqdnList,omitempty"`
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty"`
}
