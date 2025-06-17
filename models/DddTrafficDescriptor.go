/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DddTrafficDescriptor struct {
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
	PortNumber *int   `json:"portNumber,omitempty"`
	MacAddr    string `json:"macAddr,omitempty"`
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
}
