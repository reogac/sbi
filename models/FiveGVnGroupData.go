/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	Dnn                      string          `json:"dnn"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	SNssai                   Snssai          `json:"sNssai"`
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
}
