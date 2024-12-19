/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:26:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	Cause                 N2Cause                   `json:"cause"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	Guami                 Guami                     `json:"guami"`
	SecurityContext       *SecurityContext          `json:"securityContext,omitempty"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
}
