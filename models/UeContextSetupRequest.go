/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	Guami             Guami                     `json:"guami"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
}
