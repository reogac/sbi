/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:35 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	Guami                 Guami                     `json:"guami"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	AmfUeId               int64                     `json:"amfUeId"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Cause                 N2Cause                   `json:"cause"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
}
