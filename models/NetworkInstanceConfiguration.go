/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:59:14 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NetworkInstanceConfiguration struct {
	Pcscf         *IpAddr                     `json:"pcscf,omitempty"`
	Mtu           *int32                      `json:"mtu,omitempty"`
	Name          string                      `json:"name"`
	Dnais         []string                    `json:"dnais,omitempty"`
	AddressSpaces []AddressSpaceConfiguration `json:"addressSpaces,omitempty"`
	Dns           *IpAddr                     `json:"dns,omitempty"`
}
