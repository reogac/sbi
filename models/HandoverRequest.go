/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:06 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	HandoverType          int16                     `json:"handoverType"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Guami                 Guami                     `json:"guami"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
}
