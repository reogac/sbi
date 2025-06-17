/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	SNssai                   Snssai          `json:"sNssai"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	Dnn                      string          `json:"dnn"`
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
}
