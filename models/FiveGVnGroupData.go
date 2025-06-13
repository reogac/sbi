/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	Dnn                      string          `json:"dnn"`
	SNssai                   Snssai          `json:"sNssai"`
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
}
