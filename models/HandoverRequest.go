/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 11:08:48 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	Guami                 Guami                     `json:"guami"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	SecCtx                SecurityContext           `json:"secCtx"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	Cause                 N2Cause                   `json:"cause"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
}
