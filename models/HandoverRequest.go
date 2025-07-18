/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	AmfUeId               int64                     `json:"amfUeId"`
	HandoverType          int16                     `json:"handoverType"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Guami                 Guami                     `json:"guami"`
	Cause                 N2Cause                   `json:"cause"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
}
