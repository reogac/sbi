/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	Guami             Guami                     `json:"guami"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
}
