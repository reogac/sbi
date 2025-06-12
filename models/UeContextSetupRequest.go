/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupRequest struct {
	UeRadCap          string                    `json:"ueRadCap,omitempty"`
	SecKey            []byte                    `json:"secKey,omitempty"`
	AllowedNssai      []AllowedSnssai           `json:"allowedNssai,omitempty"`
	UeSecCap          UeSecurityCapability      `json:"ueSecCap"`
	OldAmf            string                    `json:"oldAmf,omitempty"`
	Guami             Guami                     `json:"guami"`
	UeAmbr            UeAmbr                    `json:"ueAmbr"`
	Rfsp              *int64                    `json:"rfsp,omitempty"`
	OldAmfNgapId      *int64                    `json:"oldAmfNgapId,omitempty"`
	N2SmInfoDownlinks []N2SmInfoDownlinkContent `json:"N2SmInfoDownlinks,omitempty"`
	NasPdu            []byte                    `json:"nasPdu,omitempty"`
	RrcStatusReport   *int16                    `json:"rrcStatusReport,omitempty"`
}
