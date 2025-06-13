/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	Guami             Guami                     `json:"guami"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
}
