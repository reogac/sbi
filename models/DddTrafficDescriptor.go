/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DddTrafficDescriptor struct {
	MacAddr    string `json:"macAddr,omitempty"`
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
	PortNumber *int   `json:"portNumber,omitempty"`
}
