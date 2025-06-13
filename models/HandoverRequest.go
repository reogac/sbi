/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Guami                 Guami                     `json:"guami"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	AmfUeId               int64                     `json:"amfUeId"`
}
