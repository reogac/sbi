/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	Dnn                      string          `json:"dnn"`
	SNssai                   Snssai          `json:"sNssai"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
}
