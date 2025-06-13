/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	Dnn                      string          `json:"dnn"`
	SNssai                   Snssai          `json:"sNssai"`
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
}
