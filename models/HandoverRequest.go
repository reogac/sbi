/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:28:45 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	Guami                 Guami                     `json:"guami"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
}
