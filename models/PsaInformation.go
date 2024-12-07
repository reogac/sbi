/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PsaInformation struct {
	DnaiList     []string      `json:"dnaiList,omitempty"`
	UeIpv6Prefix string        `json:"ueIpv6Prefix,omitempty"`
	PsaUpfId     string        `json:"psaUpfId,omitempty"`
	PsaInd       PsaIndication `json:"psaInd,omitempty"`
}
