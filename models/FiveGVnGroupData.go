/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	Dnn                      string          `json:"dnn"`
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	SNssai                   Snssai          `json:"sNssai"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
}
