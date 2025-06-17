/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:58 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RouteInformation struct {
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
	PortNumber int    `json:"portNumber"`
}
