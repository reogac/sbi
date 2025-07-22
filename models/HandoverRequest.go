/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	AmfUeId               int64                     `json:"amfUeId"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	Guami                 Guami                     `json:"guami"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Cause                 N2Cause                   `json:"cause"`
}
