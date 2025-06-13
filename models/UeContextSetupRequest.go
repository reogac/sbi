/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	Guami             Guami                     `json:"guami"`
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
}
