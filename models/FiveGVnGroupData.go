/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	Dnn                      string          `json:"dnn"`
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	SNssai                   Snssai          `json:"sNssai"`
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
}
