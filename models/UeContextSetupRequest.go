/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	Guami             Guami                     `json:"guami"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
}
