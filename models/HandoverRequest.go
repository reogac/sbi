/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:24 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	SecCtx                SecurityContext           `json:"secCtx"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Guami                 Guami                     `json:"guami"`
	AllowedNssai          *AllowedNssai             `json:"allowedNssai,omitempty"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecCap              UeSecurityCapability      `json:"ueSecCap"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
}
