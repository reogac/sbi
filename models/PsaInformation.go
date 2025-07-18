/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PsaInformation struct {
	PsaInd       PsaIndication `json:"psaInd,omitempty"`
	DnaiList     []string      `json:"dnaiList,omitempty"`
	UeIpv6Prefix string        `json:"ueIpv6Prefix,omitempty"`
	PsaUpfId     string        `json:"psaUpfId,omitempty"`
}
