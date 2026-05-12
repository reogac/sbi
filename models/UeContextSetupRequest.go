/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:33 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	Guami             Guami                     `json:"guami"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	AmfUeInfo         *AmfUeContextInfo         `json:"amfUeInfo,omitempty"`
}
