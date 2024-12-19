/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:48:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	Guami             Guami                     `json:"guami"`
}
