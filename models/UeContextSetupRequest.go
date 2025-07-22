/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
	Guami             Guami                     `json:"guami"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
}
