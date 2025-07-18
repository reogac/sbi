/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupData struct {
	DnAaaAddress             *IpAddress      `json:"dnAaaAddress,omitempty"`
	DnAaaFqdn                string          `json:"dnAaaFqdn,omitempty"`
	Dnn                      string          `json:"dnn"`
	DnAaaIpAddressAllocation *bool           `json:"dnAaaIpAddressAllocation,omitempty"`
	AppDescriptors           []AppDescriptor `json:"appDescriptors,omitempty"`
	SecondaryAuth            *bool           `json:"secondaryAuth,omitempty"`
	AdditionalDnAaaAddresses []IpAddress     `json:"additionalDnAaaAddresses,omitempty"`
	SNssai                   Snssai          `json:"sNssai"`
	PduSessionTypes          []string        `json:"pduSessionTypes,omitempty"`
}
