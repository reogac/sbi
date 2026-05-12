/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PsaInformation struct {
	PsaInd       PsaIndication `json:"psaInd,omitempty"`
	DnaiList     []string      `json:"dnaiList,omitempty"`
	UeIpv6Prefix string        `json:"ueIpv6Prefix,omitempty"`
	PsaUpfId     string        `json:"psaUpfId,omitempty"`
}
